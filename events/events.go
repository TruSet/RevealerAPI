package events

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
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
	currentNonce                      uint64
)

//const key = "{\"address\":\"05bbf03bf6f3293c826bc924e7cacc7b43ed5589\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"3dcbd47fe205399a7a59353b407f6132ac07656c4f0bd7268f1393ae8f4269f1\",\"cipherparams\":{\"iv\":\"473f6f9ff474d180fad649f653c89ac0\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"f940031c6e41f1e187877e964d4caa58d1c4789e58e174fddc06f6fb16dfeb4b\"},\"mac\":\"19fd966b0310dc91439bbf9c009b44aae2ac03d71d789fef1543bf25c103401b\"},\"id\":\"9744b359-d944-4a49-92bf-32abc657dcdd\",\"version\":3}"
// TODO: (A) get this from somewhere more secure (Heroku environment variable?)
//       (B) use an account that is never used anywhere else
//       If we use an account that is used elsewhere, the possibility of concurrent transactions exposes us to bugs in geth's
//       logic for getting the next transaction nonce. (See https://github.com/ethereum/go-ethereum/issues/15349 and
//       https://github.com/ethereum/go-ethereum/issues/2880; note that PendingNonceAt calls eth_getTransactionCount.)
// This is the account we send our transactions from
const key = "{\"address\":\"9772dfd46049db8208efd79f01b44dfbd1c33638\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"37d2d7358aeb6b0c3bbcafe6b67ae580ed820ab5faa7075075017fa57a2a7f29\",\"cipherparams\":{\"iv\":\"d48799327768a8e4889aaa96bc62226e\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"dbc1e7a5a84e1fa0bf0081a2559a5123ce2c49d9df59416c5791452598031030\"},\"mac\":\"5f1da9f5c1b1dc6339068fd5c5dbebbd15488627a01d4b6fa5d30a546cc44eb1\"},\"id\":\"fdffe93f-f59a-493a-a652-5222ed8dc206\",\"version\":3}"

func getLogTopic(eventSignature string) common.Hash {
	return common.HexToHash("0x" + hex.EncodeToString(solsha3.SoliditySHA3(solsha3.String(eventSignature))))
}

var VoteCommittedLogTopic = getLogTopic("VoteCommitted(bytes32,address,bytes32)")
var VoteRevealedLogTopic = getLogTopic("VoteRevealed(bytes32,bytes32,uint256,address,address,uint256,uint256)")
var CommitPeriodHaltedLogTopic = getLogTopic("CommitPeriodHalted(bytes32,address,uint256)")
var RevealPeriodStartedLogTopic = getLogTopic("RevealPeriodStarted(bytes32,address,bytes32,bytes32)")
var RevealPeriodHaltedLogTopic = getLogTopic("RevealPeriodHalted(bytes32,address,uint256)")
var PollCreatedLogTopic = getLogTopic("PollCreated(bytes32,address,uint256,uint256)")

func revealTransactionOpts(client *ethclient.Client) *bind.TransactOpts {
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	log.Println("FROM", auth.From.Hex())
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
	suggestedNonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatalf("failed to retrieve account nonce: %v", err)
	}
	currentNonce = suggestedNonce

	commitRevealVotingABI, _ := abi.JSON(strings.NewReader(contract.TruSetCommitRevealVotingABI))
	boundContract = bind.NewBoundContract(commitRevealVotingContractAddress, commitRevealVotingABI, nil, nil, nil)
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

func RevealCommitments(client *ethclient.Client, revealPeriodStarted *contract.TruSetCommitRevealVotingRevealPeriodStarted) {
	//var mutex = &sync.Mutex{}
	commitments := fetchCommitments(revealPeriodStarted.PollID)
	log.Println("-- NEW COMMITMENTS TO REVEAL --", len(commitments))

	//revealQueue := make(chan *database.Commitment)
	//buf := make([]big.Int,0,len(commitments))
	for i := 0; i < len(commitments); i++ {

		currentNonce++
		//mutex.Lock()
		//newNonce := int64(currentNonce)
		//votingSession.TransactOpts.Nonce = big.NewInt(newNonce)
		votingSession.TransactOpts.Nonce = nil
		commitment := commitments[i]
		//revealQueue <- &commitments[i]
		//}
		//close(revealQueue)
		//for commitment := range revealQueue {
		//for _, commitment := range commitments {
		//commitment := commitments[1]
		//log.Println("Nonce in votingSession", votingSession.

		//func(commitment database.Commitment, votingSession *contract.TruSetCommitRevealVotingSession) {
		//time.Sleep(5 * time.Second)

		nonce := votingSession.TransactOpts.Nonce
		log.Println("[Revealing Vote]\t", nonce, hexutil.Encode(revealPeriodStarted.PollID[:]), commitment.VoterAddress, votingSession.TransactOpts)
		//buf = append(buf, *nonce)
		trans, err := votingSession.RevealVote(
			revealPeriodStarted.InstrumentAddress,
			revealPeriodStarted.DataIdentifier,
			revealPeriodStarted.PayloadHash,
			common.HexToAddress(commitment.VoterAddress),
			big.NewInt(int64(commitment.VoteOption)),
			big.NewInt(int64(commitment.Salt)),
		)

		if err != nil {
			fmt.Println("[Reveal Submission Failed]\t", nonce, commitment, err)
		} else {
			//fmt.Println("[Reveal Succeeded]\t", commitment)
			log.Println("[Reveal Submitted]", trans.Nonce())
		}

		// TODO: here and elsewhere we want to use a cancellable context
		//       this call will hang indefinitely until our transaction is mined or the context is cancelled
		receipt, err := bind.WaitMined(context.Background(), client, trans)
		if err != nil {
			log.Fatalf("[Reveal Trx Failed]\t", nonce, commitment, err)
		}
		//receipt, err := client.TransactionReceipt(context.Background(), client, trans.hash)
		log.Println("[Reveal Successful]", receipt)

		//mutex.Unlock()
		//}(commitment, votingSession)
	}
	//log.Println("Nonces", buf)
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
