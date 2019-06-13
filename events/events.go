package events

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TruSet/RevealerAPI/contract"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/getsentry/raven-go"
	"github.com/miguelmota/go-solidity-sha3"
)

var (
	commitRevealVotingContractAddress common.Address
	filter                            ethereum.FilterQuery
	topics                            [][]common.Hash
	votingSession                     *contract.TruSetCommitRevealVotingSession
	contractCallSession               *contract.TruSetCommitRevealVotingCallerSession
	boundContract                     *bind.BoundContract
	from                              common.Address
	client                            *ethclient.Client
	clientString                      string
	processingPastEvents              bool
)

func getLogTopic(eventSignature string) common.Hash {
	return common.HexToHash("0x" + hex.EncodeToString(solsha3.SoliditySHA3(solsha3.String(eventSignature))))
}

var VoteCommittedLogTopic = getLogTopic("VoteCommitted(bytes32,address,bytes32)")
var VoteRevealedLogTopic = getLogTopic("VoteRevealed(bytes32,bytes32,uint256,address,address,uint256,uint256,uint256)")
var CommitPeriodHaltedLogTopic = getLogTopic("CommitPeriodHalted(bytes32,address,uint256)")
var RevealPeriodStartedLogTopic = getLogTopic("RevealPeriodStarted(bytes32,address,bytes32,bytes32)")
var RevealPeriodHaltedLogTopic = getLogTopic("RevealPeriodHalted(bytes32,address,uint256)")
var PollCreatedLogTopic = getLogTopic("PollCreated(bytes32,address,uint256,uint256)")

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvInt(key string) (int, error) {
	s := os.Getenv(key)
	if s == "" {
		return 0, ErrEnvVarEmpty
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func revealTransactionOpts(client *ethclient.Client) *bind.TransactOpts {
	key := os.Getenv("REVEALER_KEY")
	passphrase := os.Getenv("REVEALER_PASSPHRASE")

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatal("Error reading private key: ", err)
	}

	// In some environments, estimating gas is a problem so we allow it to be hardcoded
	// Defaults to zero, which tells go-ethereum to estimate it
	gasLimit, _ := getenvInt("GAS_LIMIT")
	auth.GasLimit = (uint64)(gasLimit)

	gasPriceInGwei, err := getenvInt("GAS_PRICE_IN_GWEI")
	if err != nil {
		// Use a gas price oracle
		auth.GasPrice = nil
	} else {
		auth.GasPrice = new(big.Int).Mul(big.NewInt((int64)(gasPriceInGwei)), big.NewInt(params.GWei))
	}

	log.Printf("Using gas limit %v and gas price %v", auth.GasLimit, auth.GasPrice)

	return auth
}

func dialClient(_clientString string) {
	var err error

	log.Println("Dialling client:", _clientString)
	client, err = ethclient.Dial(_clientString)
	log.Println("... dialling complete.")

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client %v: %v", _clientString, err)
	}
}

func Init(_clientString string, commitRevealVotingAddress string) {
	commitRevealVotingContractAddress = common.HexToAddress(commitRevealVotingAddress)
	clientString = _clientString
	dialClient(clientString)
	processingPastEvents = false

	topics = [][]common.Hash{{
		VoteCommittedLogTopic,
		VoteRevealedLogTopic,
		CommitPeriodHaltedLogTopic,
		VoteRevealedLogTopic,
		RevealPeriodHaltedLogTopic,
		PollCreatedLogTopic,
		RevealPeriodStartedLogTopic}}

	votingContract, _ := contract.NewTruSetCommitRevealVoting(commitRevealVotingContractAddress, client)
	opts := revealTransactionOpts(client)

	votingSession = &contract.TruSetCommitRevealVotingSession{
		Contract:     votingContract,
		TransactOpts: *opts,
	}

	from = opts.From
	commitRevealVotingABI, _ := abi.JSON(strings.NewReader(contract.TruSetCommitRevealVotingABI))
	boundContract = bind.NewBoundContract(commitRevealVotingContractAddress, commitRevealVotingABI, nil, nil, nil)

	votingContractCaller, _ := contract.NewTruSetCommitRevealVotingCaller(commitRevealVotingContractAddress, client)

	contractCallSession = &contract.TruSetCommitRevealVotingCallerSession{
		Contract: votingContractCaller,
		CallOpts: *new(bind.CallOpts),
	}
}

func getCRVLogFilter() ethereum.FilterQuery {
	return ethereum.FilterQuery{
		Addresses: []common.Address{commitRevealVotingContractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // Latest block
		//Topics:    nil, // Match any topic, for testing
		Topics: topics,
	}
}

func ProcessPastEvents(fromBlock uint64) {
	var startingWindowSize uint64 = 1000000 // TODO: get from env var?
	windowSize := startingWindowSize
	var startBlock uint64 = fromBlock
	endBlock := startBlock + windowSize

	log.Printf("Processing past events")
	processingPastEvents = true

	for startBlock <= CurrentBlockNumber() {
		log.Printf("Processing logs for address %v, blocks %v (exclusive) to %v (inclusive)", commitRevealVotingContractAddress.String(), startBlock, endBlock)
		err := processPastEventWindow(startBlock, endBlock)

		if err != nil {
			log.Printf("error: %v", err.Error())
			if strings.Contains(err.Error(), "returned more than") {
				// Retry with smaller window size, always non-empty
				windowSize = windowSize / 4
				if windowSize < 1 {
					windowSize = 1
				}

				endBlock = startBlock + windowSize
				log.Printf("Too many logs found for address %v; adjusting window size to %v and retrying", commitRevealVotingContractAddress.String(), windowSize)
			} else {
				// Unexpected error
				raven.CaptureError(err, nil)
				log.Fatalf("Failed to get past logs for address %v. Error: %v", commitRevealVotingContractAddress.String(), err)
			}
		} else {
			// Success; move and enlarge the window
			windowSize = windowSize * 2
			startBlock = endBlock
			endBlock = startBlock + windowSize
		}
	}
	processingPastEvents = false
}

func processPastEventWindow(fromBlock uint64, toBlock uint64) error {
	ctx := context.Background()

	filter := getCRVLogFilter()
	filter.FromBlock = big.NewInt(int64(fromBlock))
	filter.ToBlock = big.NewInt(int64(toBlock))

	if toBlock > CurrentBlockNumber() {
		filter.ToBlock = nil
	}

	// get past logs
	logs, err := client.FilterLogs(context.Background(), filter)
	if err != nil {
		return err
	}

	for _, l := range logs {
		processLog(client, ctx, l)
	}

	log.Println("End of existing logs")
	return nil
}

// Uncomment if mocking websockets connection errors
// func sleepThenSendTrueTo(duration time.Duration, channelToSendTrueTo chan<- bool) {
// 	time.Sleep(duration)
// 	channelToSendTrueTo <- true
// }

func ProcessFutureEvents() {
	ctx := context.Background()
	filter := getCRVLogFilter()

	// TODO:
	// Infura closes idle websockets connections. We could keep the websockets
	// connection alive using Ping and Pong packets, but the websocket library
	// we are using does not yet support these
	// (https://github.com/golang/go/issues/5958). As a workaround,
	// we attempt to re-establish the connection and the subscription whenever
	// the connection times out.
	//
	// We naively guard against an infinite tight loop by bailing out completely if
	// the connection fails or is closed twice within a second.
	//
	// Nothing about this is water-tight. We are prone to bailing out
	// fatally, unnecessarily, when periodic retries would be preferable.
	// And even where the re-subscription succeeds, there's a danger that
	// we miss logs in the interval since the last subscription closed.
	// But in practice this quick and dirty solution seems to be
	// "Good Enough", probably until we do proper Ping/Pong or move
	// away from Infura towards our own local client.

	// Uncomment to mock websockets connection errors
	// testChan := make(chan bool)
	// go sleepThenSendTrueTo(time.Duration(20)*time.Second, testChan)
	for {
		log.Println("Log subscription starting")
		subscriptionTime := time.Now()

		// subscribe for new logs
		// TODO: how to handle chain rollbacks and the resulting invalid logs? Probably monitor even status.
		ch := make(chan types.Log, 100)
		sub, err := client.SubscribeFilterLogs(ctx, filter, ch)
		if err != nil {
			log.Fatal(err)
			return
		}

		defer sub.Unsubscribe()
		errChan := sub.Err()

	pollingLoop:
		for {
			select {
			// Uncomment to mock websockets connection errors
			// case <-testChan:
			// 	log.Println("Subscription closed by test channel at time", time.Now())
			// 	break pollingLoop
			case err := <-errChan:
				log.Println("Logs subscription error", err)
				msg := fmt.Sprintf("Subscription terminated with error")
				raven.CaptureError(err, map[string]string{"message": msg, "startedAt": subscriptionTime.String()})
				break pollingLoop
			case l := <-ch:
				processLog(client, ctx, l)
			}
		}
		log.Println("Log subscription terminated")

		// If we've failed twice in quick succession, we abort. Otherwise, dial the client again
		// and re-subscribe for events.
		secondsSinceLastFailure := time.Since(subscriptionTime).Seconds()
		if secondsSinceLastFailure < 1 {
			log.Printf("It has been only %v seconds since our previous failure. Sleeping.", secondsSinceLastFailure)
			time.Sleep(time.Duration(10) * time.Second)
		}

		dialClient(clientString)
	}
	log.Println("Log subscription terminated permanently.")
}
