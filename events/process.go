package events

import (
	"context"
	"log"
	"math/big"

	"github.com/TruSet/RevealerAPI/contract"
	"github.com/TruSet/RevealerAPI/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/getsentry/raven-go"
)

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
	case PollCreatedLogTopic:
		pollCreated := new(contract.TruSetCommitRevealVotingPollCreated)
		boundContract.UnpackLog(pollCreated, "PollCreated", l)

		//log.Printf("[Poll Created]\t%s", hexutil.Encode(pollCreated.PollID[:]))
	case RevealPeriodHaltedLogTopic:
		revealPeriodHalted := new(contract.TruSetCommitRevealVotingRevealPeriodHalted)
		boundContract.UnpackLog(revealPeriodHalted, "RevealPeriodHalted", l)

		//log.Printf("[Reveal Period Halted]\t%s", hexutil.Encode(revealPeriodHalted.PollID[:]))
	case VoteCommittedLogTopic:
		voteCommitted := new(contract.TruSetCommitRevealVotingVoteCommitted)
		boundContract.UnpackLog(voteCommitted, "VoteCommitted", l)

		if knownCommitment(voteCommitted.PollID, voteCommitted.SecretHash) {
			log.Printf("[Vote Committed] (recognised): %s : %s : %s", hexutil.Encode(voteCommitted.PollID[:]), voteCommitted.Voter.Hex(), hexutil.Encode(voteCommitted.SecretHash[:]))
			database.MarkAsMostRecentlySeen(voteCommitted.PollID, voteCommitted.Voter.Hex(), voteCommitted.SecretHash)
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
			raven.CaptureError(err, nil)
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
					raven.CaptureError(err, nil)
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

func fetchCommitments(pollID [32]byte) []database.Commitment {
	var commitments []database.Commitment
	database.Db.Where("poll_id = ? and last_on_chain = ?", hexutil.Encode(pollID[:]), true).Find(&commitments)
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
