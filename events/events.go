package events

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"strings"

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
	boundContract                     *bind.BoundContract
	from                              common.Address
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
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		log.Fatal(err)
	}
	return auth
}

func Init(client *ethclient.Client, commitRevealVotingAddress string) {
	commitRevealVotingContractAddress = common.HexToAddress(commitRevealVotingAddress)

	filter = ethereum.FilterQuery{
		Addresses: []common.Address{commitRevealVotingContractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil, // Latest block
		Topics:    nil, // Match any topic, for testing
		//Topics: [][]common.Hash{{
		//RevealPeriodStartedLogTopic}},
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
}

func processLog(client *ethclient.Client, ctx context.Context, l types.Log) {
	if l.Removed {
		// TODO: need to handle chain re-organisations gracefully by undoing the affected database change!
		// For now we expect this to be uncommon and would rather die than provide incorrect data
		log.Fatalf("Found removed flag on log %+v", l)
	}

	switch l.Topics[0] {
	case RevealPeriodStartedLogTopic:
		revealPeriodStarted := new(contract.TruSetCommitRevealVotingRevealPeriodStarted)
		// shouldn't have to use this low level boundContract
		boundContract.UnpackLog(revealPeriodStarted, "RevealPeriodStarted", l)

		log.Printf("[Reveal Period Started]\t%s", hexutil.Encode(revealPeriodStarted.PollID[:]))

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

		log.Printf("[Vote Committed]\t%s : %s", hexutil.Encode(voteCommitted.PollID[:]), voteCommitted.Voter.Hex())
	case VoteRevealedLogTopic:
		voteRevealed := new(contract.TruSetCommitRevealVotingVoteRevealed)
		boundContract.UnpackLog(voteRevealed, "VoteRevealed", l)

		log.Printf("[Vote Revealed]\t%s : %s", hexutil.Encode(voteRevealed.PollID[:]), voteRevealed.Voter.Hex())
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

func logTrxResult(ctx context.Context, b DeployBackend, tx *types.Transaction, description string) {
	receipt, err := bind.WaitMined(ctx, b, tx)
	if err != nil || receipt.Status == 0 {
		log.Printf("[%v FAILED] %+v %v %+v", s, err, receipt)
	}
	log.Printf("[%v Successful]", s)
}

func RevealCommitments(client *ethclient.Client, revealPeriodStarted *contract.TruSetCommitRevealVotingRevealPeriodStarted) {
	commitments := fetchCommitments(revealPeriodStarted.PollID)
	retryRevealsIndividually := false

	// First try revealing all votes together.
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
		fmt.Printf("[Reveal-All Submission Failed] %v", err)
		retryRevealsIndividually = true
	} else {
		// TODO: here and elsewhere we want to use a cancellable context
  	//       this call will hang indefinitely until our transaction is mined or the context is cancelled
   logTrxResult(context.Background(), client, trans, "Reveal-All Trx")
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
        log.Fatalf("[Reveal Submission Failed] %+v %v", commitment, err)
      } else {
        log.Printf("[Reveal Submission Succeeded] %+v", commitment)
        // TODO: here and elsewhere we want to use a cancellable context
        //       this call will hang indefinitely until our transaction is mined or the context is cancelled
        logTrxResult(context.Background(), client, trans, "Reveal Trx " + commitment.VoterAddress)
      }
    }
	}
}

func ProcessPastEvents(client *ethclient.Client) {
	ctx := context.Background()

	// get past logs
	logs, err := client.FilterLogs(context.Background(), filter)
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
