package events

import (
	"context"
	"fmt"
	"log"
	"math/big"


	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/TruSet/RevealerAPI/database"
)

var (
	commitRevealVotingContractAddress                     string
	filter                                                ethereum.FilterQuery
	CommitPeriodHaltedLogTopic        common.Hash
)

func Init(commitRevealVotingAddress string) {
	commitRevealVotingContractAddress = commitRevealVotingAddress

  //TODO correct this
  CommitPeriodHaltedLogTopic = common.HexToHash("0x4226fb316091e086ca1435e14c0c26a2d232c473f5d751d15eea24e996592dc1")

	// TODO: for now, our filter makes no attempt to skip blocks already processed.
	//       this may be functionally OK because the database contraints prevent duplicate rows
	//       but it leads to warnings and is inefficient so it would be nice to be more selective.
	//       When fixing, need to be mindful of chain re-orgs.
	filter = ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(commitRevealVotingContractAddress)},
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // Latest block
		//Topics:    nil, // Match any topic, for testing
		Topics: [][]common.Hash{{
      CommitPeriodHaltedLogTopic}},
	}
}

func processLog(client *ethclient.Client, ctx context.Context, l types.Log) {
	// log.Printf("Found log from address %x with topics %x and data %x\n", l.Address, l.Topics, l.Data)

	if l.Removed {
		// TODO: need to handle chain re-organisations gracefully by undoing the affected database change!
		// For now we expect this to be uncommon and would rather die than provide incorrect data
		log.Fatalf("Found removed flag on log %+v", l)
	}

  log.Println(l.Topics[0])
	switch l.Topics[0] {
	case CommitPeriodHaltedLogTopic:
    var revealStarted CommitPeriodHaltedLog
    //revealStarted.PollID = l.Topics[1]

    unpackCommitRevealVoting(&revealStarted, "CommitPeriodHaltedLog", l)

    // TODO
    // fetch commitments
    //commitments := fetchCommitments(revealStarted.PollID)
    // call out to abi for each one to reveal
	default:
		log.Fatalf("UNEXPECTED: log from address %x with topics %x and data %x\n", l.Address, l.Topics, l.Data)
	}
}

func fetchCommitments(pollID string) []database.Commitment {
  var commitments []database.Commitment
  database.Db.Where("poll_id = ? and voter_address = ?", pollID).Find(&commitments)
  return commitments
}

func revealCommitments(commiments []database.Commitment) {
  //var commitRevealVoting = eth.contract(CommitRevealVotingABI).at(commitRevealVotingContractAddress);
  //for i := 0; i < len(commitments); i++ {
    //commitment := commitments[i]
    //commitRevealVoting.revealCommitment(commitment.PiollID, commitment.VoterAddress, commitment.VoteOption, commitment.Salt).sendTransaction({from:eth.accounts[0]})
  //}
}

func unpackCommitRevealVoting(dest interface{}, logName string, l types.Log) {
	err := CommitRevealVotingABI.Unpack(dest, logName, l.Data)
	if err != nil {
		fmt.Println("Failed to unpack:", err)
	}
	log.Printf("%v: %+v\n", logName, dest)
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
