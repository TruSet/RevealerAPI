package events

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/TruSet/RevealerAPI/contract"
	"github.com/TruSet/RevealerAPI/database"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/miguelmota/go-solidity-sha3"
)

var (
	commitRevealVotingContractAddress common.Address
	filter                            ethereum.FilterQuery
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

func revealTransactionOpts(client *ethclient.Client) *bind.TransactOpts {
	key := os.Getenv("REVEALER_KEY")
	passphrase := os.Getenv("REVEALER_PASSPHRASE")
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatal(err)
	}
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

	filter = ethereum.FilterQuery{
		Addresses: []common.Address{commitRevealVotingContractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // Latest block
		//Topics:    nil, // Match any topic, for testing
    Topics: [][]common.Hash{{
      VoteCommittedLogTopic,
      VoteRevealedLogTopic,
      CommitPeriodHaltedLogTopic,
      VoteRevealedLogTopic,
      RevealPeriodHaltedLogTopic,
      PollCreatedLogTopic,
      RevealPeriodStartedLogTopic}},
	}

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

func processLog(client *ethclient.Client, ctx context.Context, l types.Log) {
	if l.Removed {
		// There has been a chain re-org but we don't really care
		// We will re-process this log if/when we see it included in the chain again
		return
	}

	switch l.Topics[0] {
	case RevealPeriodStartedLogTopic:
		revealPeriodStarted := new(contract.TruSetCommitRevealVotingRevealPeriodStarted)
		// shouldn't have to use this low level boundContract
		boundContract.UnpackLog(revealPeriodStarted, "RevealPeriodStarted", l)

		log.Printf("[Reveal Period Started]\t%s: %s / %s / %s",
			hexutil.Encode(revealPeriodStarted.PollID[:]),
			revealPeriodStarted.InstrumentAddress.String(),
			hexutil.Encode(revealPeriodStarted.DataIdentifier[:]),
			hexutil.Encode(revealPeriodStarted.PayloadHash[:]))

		RevealCommitments(client, revealPeriodStarted)
	case CommitPeriodHaltedLogTopic:
		commitPeriodHalted := new(contract.TruSetCommitRevealVotingCommitPeriodHalted)
		boundContract.UnpackLog(commitPeriodHalted, "CommitPeriodHalted", l)

		log.Printf("[Commit Period Halted]\t%s", hexutil.Encode(commitPeriodHalted.PollID[:]))
	case RevealPeriodHaltedLogTopic:
		revealPeriodHalted := new(contract.TruSetCommitRevealVotingRevealPeriodHalted)
		boundContract.UnpackLog(revealPeriodHalted, "RevealPeriodHalted", l)

		log.Printf("[Reveal Period Halted]\t%s", hexutil.Encode(revealPeriodHalted.PollID[:]))
	case PollCreatedLogTopic:
		pollCreated := new(contract.TruSetCommitRevealVotingPollCreated)
		boundContract.UnpackLog(pollCreated, "PollCreated", l)

		log.Printf("[Poll Created]\t%s", hexutil.Encode(pollCreated.PollID[:]))
	case RevealPeriodHaltedLogTopic:
		revealPeriodHalted := new(contract.TruSetCommitRevealVotingRevealPeriodHalted)
		boundContract.UnpackLog(revealPeriodHalted, "RevealPeriodHalted", l)

		log.Printf("[Reveal Period Halted]\t%s", hexutil.Encode(revealPeriodHalted.PollID[:]))
	case VoteCommittedLogTopic:
		voteCommitted := new(contract.TruSetCommitRevealVotingVoteCommitted)
		boundContract.UnpackLog(voteCommitted, "VoteCommitted", l)

		if knownCommitment(voteCommitted.PollID, voteCommitted.SecretHash) {
			log.Printf("[Vote Committed] (recognised): %s : %s : %s", hexutil.Encode(voteCommitted.PollID[:]), voteCommitted.Voter.Hex(), hexutil.Encode(voteCommitted.SecretHash[:]))
		} else {
			log.Printf("[Vote Committed] UNRECOGNISED: %s : %s : %s", hexutil.Encode(voteCommitted.PollID[:]), voteCommitted.Voter.Hex(), hexutil.Encode(voteCommitted.SecretHash[:]))
		}

	case VoteRevealedLogTopic:
		voteRevealed := new(contract.TruSetCommitRevealVotingVoteRevealed)
		boundContract.UnpackLog(voteRevealed, "VoteRevealed", l)

		log.Printf("[Vote Revealed]\t%s : %s : %d", hexutil.Encode(voteRevealed.PollID[:]), voteRevealed.Voter.Hex(), voteRevealed.Choice)
	default:
		log.Printf("[Unexpected]\tlog with topics %x\n", l.Topics)
		log.Println(l.Topics[0])
	}
}

func fetchCommitments(pollID [32]byte) []database.Commitment {
	var commitments []database.Commitment
	database.Db.Where("poll_id = ?", hexutil.Encode(pollID[:])).Find(&commitments)
	return commitments
}

func knownCommitment(pollID [32]byte, commitHash [32]byte) bool {
	var commitments []database.Commitment
	database.Db.Unscoped().Where("poll_id = ? and commit_hash = ?", hexutil.Encode(pollID[:]), hexutil.Encode(commitHash[:])).Find(&commitments)
	return len(commitments) > 0
}

func processRevealResult(ctx context.Context, client *ethclient.Client, tx *types.Transaction, pollID [32]byte, voterAddress string) {
	var description string
	receipt, err := bind.WaitMined(ctx, client, tx)

	if voterAddress == "" {
		description = "Reveal-All Trx"
	} else {
		description = "Reveal Trx for " + voterAddress
	}

	if err != nil || receipt.Status == 0 {
		log.Printf("[%v FAILED] %v %+v", description, err, receipt)
	} else {
		// Mark the revealed proposals as revealed
		log.Printf("[%v Successful]", description)
		database.SoftDeleteRevealed(pollID, voterAddress)
	}
}

func RevealCommitments(client *ethclient.Client, revealPeriodStarted *contract.TruSetCommitRevealVotingRevealPeriodStarted) {
	commitments := fetchCommitments(revealPeriodStarted.PollID)
	ignoreThisPoll := len(commitments) == 0
	log.Println("Ignore due to length?", ignoreThisPoll)

	// We don't bother checking the poll status for real-time events
	if !ignoreThisPoll && processingPastEvents {
		pollEnded, err := contractCallSession.PollEnded(revealPeriodStarted.PollID)
		ignoreThisPoll = (err == nil) && pollEnded
		log.Println("Ignore due to poll status?", ignoreThisPoll)
	}

	if !ignoreThisPoll {
		// There is work to do. First try revealing all votes in one transaction.
		retryRevealsIndividually := false

		var voters []common.Address
		var voteOptions []*big.Int
		var salts []*big.Int
		for i := 0; i < len(commitments); i++ {
			commitment := commitments[i]
			voters = append(voters, common.HexToAddress(commitment.VoterAddress))
			voteOptions = append(voteOptions, big.NewInt(int64(commitment.VoteOption)))
			salts = append(salts, big.NewInt(int64(commitment.Salt)))
		}

		trans, err := votingSession.RevealVotes(
			revealPeriodStarted.InstrumentAddress,
			revealPeriodStarted.DataIdentifier,
			revealPeriodStarted.PayloadHash,
			voters,
			voteOptions,
			salts,
		)

		if err != nil {
			log.Printf("[Reveal-All Submission FAILED] %v", err)
			retryRevealsIndividually = true
		} else {
			// TODO: here and elsewhere we want to use a cancellable context
			//       this call will hang indefinitely until our transaction is mined or the context is cancelled
			processRevealResult(context.Background(), client, trans, revealPeriodStarted.PollID, "")
		}

		// // If we could not reveal all votes together, fall back to revealing them individually
		if retryRevealsIndividually {
			for i := 0; i < len(commitments); i++ {
				commitment := commitments[i]
				trans, err := votingSession.RevealVote(
					revealPeriodStarted.InstrumentAddress,
					revealPeriodStarted.DataIdentifier,
					revealPeriodStarted.PayloadHash,
					common.HexToAddress(commitment.VoterAddress),
					big.NewInt(int64(commitment.VoteOption)),
					big.NewInt(int64(commitment.Salt)),
				)
				if err != nil {
					log.Printf("[Reveal Submission FAILED] %+v %v", commitment, err)
				} else {
					// TODO: here and elsewhere we want to use a cancellable context
					//       this call will hang indefinitely until our transaction is mined or the context is cancelled
					processRevealResult(context.Background(), client, trans, revealPeriodStarted.PollID, commitment.VoterAddress)
				}
			}
		}
	}
}

func ProcessPastEvents() {

	log.Printf("Processing past events")
	processingPastEvents = true
	ctx := context.Background()

	// get past logs
	logs, err := client.FilterLogs(context.Background(), filter)
	if err != nil {
		log.Fatalf("Failed to get past logs. Error: %v", err)
		return
	}

	for _, l := range logs {
		processLog(client, ctx, l)
	}

	processingPastEvents = false
	log.Println("End of existing logs")
}

// Uncomment if mocking websockets connection errors
// func sleepThenSendTrueTo(duration time.Duration, channelToSendTrueTo chan<- bool) {
// 	time.Sleep(duration)
// 	channelToSendTrueTo <- true
// }

func ProcessFutureEvents() {
	ctx := context.Background()

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
