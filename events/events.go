package events

import (
	"context"
	"fmt"
	"log"
  "strings"
  "crypto/ecdsa"
	"math/big"
  "encoding/hex"

  "github.com/TruSet/RevealerAPI/database"
  "github.com/TruSet/RevealerAPI/contract"

  ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/crypto"

  "github.com/miguelmota/go-solidity-sha3"
)

const key = "this is a key"

var (
	commitRevealVotingContractAddress                     common.Address
	filter                                                ethereum.FilterQuery
  PollCreatedLogTopic               common.Hash
	CommitPeriodHaltedLogTopic        common.Hash
	RevealPeriodStartedLogTopic       common.Hash
	RevealPeriodHaltedLogTopic        common.Hash
	VoteCommittedLogTopic             common.Hash
	VoteRevealedLogTopic              common.Hash
  instance                          *contract.TruSetCommitRevealVoting
  boundContract                     *bind.BoundContract
  client                            *ethclient.Client
)

func getLogTopic(eventSignature string) (common.Hash) {
  return common.HexToHash("0x" + hex.EncodeToString(solsha3.SoliditySHA3(solsha3.String(eventSignature))))
}

func transactionOpts2(client *ethclient.Client) (*bind.TransactOpts) {
  auth, _ := bind.NewTransactor(strings.NewReader(key), "passphrase associated with your JSON key file")
  return auth
}

func transactionOpts(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts) {
  privateKeyString := "f3782740e767072087ebf6a2a77730b8cc9bddc1bd078c63da511e77763dce65"
  privateKey, err := crypto.HexToECDSA(privateKeyString)
  if err != nil {
    log.Fatal(err)
  }

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
    log.Fatal("error casting public key to ECDSA")
  }

  fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
  log.Printf("- Setting transaction options with address %s", hexutil.Encode(fromAddress[:]))
  nonce, err := client.PendingNonceAt(ctx, fromAddress)
  if err != nil {
    log.Println("Fails here")
    log.Fatal(err)
  }
  gasPrice, err := client.SuggestGasPrice(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  auth := bind.NewKeyedTransactor(privateKey)
  auth.Nonce = big.NewInt(int64(nonce))
  auth.Value = big.NewInt(int64(0))     // in wei
  auth.GasLimit = uint64(300000) // in units
  auth.GasPrice = gasPrice

  return auth
}

func Init(_client *ethclient.Client, commitRevealVotingAddress string) {
  client := _client
	commitRevealVotingContractAddress = common.HexToAddress(commitRevealVotingAddress)

  VoteCommittedLogTopic = getLogTopic("VoteCommitted(bytes32,address,bytes32)")
  VoteRevealedLogTopic = getLogTopic("VoteRevealed(bytes32,bytes32,uint256,address,address,uint256,uint256)")
  CommitPeriodHaltedLogTopic = getLogTopic("CommitPeriodHalted(bytes32,address,uint256)")
  RevealPeriodStartedLogTopic = getLogTopic("RevealPeriodStarted(bytes32,address,bytes32,bytes32)")
  RevealPeriodHaltedLogTopic = getLogTopic("RevealPeriodHalted(bytes32,address,uint256)")
  PollCreatedLogTopic = getLogTopic("PollCreated(bytes32,address,uint256,uint256)")

	filter = ethereum.FilterQuery{
		Addresses: []common.Address{commitRevealVotingContractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // Latest block
    Topics:    nil, // Match any topic, for testing
    //Topics: [][]common.Hash{{
      //RevealPeriodStartedLogTopic}},
	}
  instance, _ = contract.NewTruSetCommitRevealVoting(commitRevealVotingContractAddress, client)

  boundContract = bind.NewBoundContract(commitRevealVotingContractAddress, CommitRevealVotingABI, nil, nil , nil)
}

func processLog(client *ethclient.Client, ctx context.Context, l types.Log) {
	// log.Printf("Found log from address %x with topics %x and data %x\n", l.Address, l.Topics, l.Data)

	if l.Removed {
		// TODO: need to handle chain re-organisations gracefully by undoing the affected database change!
		// For now we expect this to be uncommon and would rather die than provide incorrect data
		log.Fatalf("Found removed flag on log %+v", l)
	}

	switch l.Topics[0] {
	case RevealPeriodStartedLogTopic:

    revealStarted := new(contract.TruSetCommitRevealVotingRevealPeriodStarted)
    boundContract.UnpackLog(revealStarted, "RevealPeriodStarted", l)

    commitments := fetchCommitments(revealStarted.PollID)
    RevealCommitments(ctx, commitments, revealStarted.InstrumentAddress, revealStarted.DataIdentifier, revealStarted.PayloadHash)
	case CommitPeriodHaltedLogTopic:
    //log.Println("COMMIT PERIOD HALTED")
	case PollCreatedLogTopic:
    log.Printf("[Poll Created]\t%x", l.Topics[1])
	case RevealPeriodHaltedLogTopic:
      //log.Println("REVEAL PERIOD HALTED")
	case VoteCommittedLogTopic:
      //log.Println("VOTE COMMITTED")
  case VoteRevealedLogTopic:
      //log.Println("VOTE REVEALED")
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

func RevealCommitments(ctx context.Context, commitments []database.Commitment, instrumentAddress common.Address, dataIdentifier [32]byte, payloadHash [32]byte) {
  //log.Println(instrumentAddress)
  //log.Println(hexutil.Encode(dataIdentifier[:]))
  //log.Println(hexutil.Encode(payloadHash[:]))

  for i := 0; i < len(commitments); i++ {
    commitment := commitments[i]
    log.Printf("[Vote Revealed TODO]\t%s voted %d on %s", commitment.VoterAddress, commitment.VoteOption, commitment.PollID)
    opts := transactionOpts2(client)
    log.Println(opts)

    // TODO this throws a segmentation fault
    //instance.RevealVote(
      //opts,
      //instrumentAddress,
      //dataIdentifier,
      //payloadHash,
      //common.HexToAddress(commitment.VoterAddress),
      //big.NewInt(int64(commitment.VoteOption)),
      //big.NewInt(int64(commitment.Salt)))
  }
}

func unpackCommitRevealVoting(dest interface{}, logName string, l types.Log) {
	err := CommitRevealVotingABI.Unpack(dest, logName, l.Data)
	if err != nil {
		fmt.Println("Failed to unpack:", err)
	}
	//log.Printf("%v: %+v\n", logName, dest)
}

func ProcessPastEvents(client *ethclient.Client) {
	// TODO: we probably want a top-level cancellable context to pass to geth
	ctx := context.Background()

	// get past logs
	logs, err := client.FilterLogs(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, l := range logs {
		processLog(client, ctx, l)
	}

	log.Println("End of existing logs")
}

func ProcessFutureEvents(client *ethclient.Client) {
	// TODO: we probably want a top-level cancellable context to pass to geth
	ctx := context.Background()
	log.Println("Log subscription starting")

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

	for {
		select {
		case err := <-errChan:
			log.Println("Logs subscription error", err)
			break
		case l := <-ch:
			processLog(client, ctx, l)
		}
	}

	log.Println("Log subscription terminated")
}
