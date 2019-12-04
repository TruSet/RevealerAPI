package database

import (
	"encoding/hex"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/miguelmota/go-solidity-sha3"
)

type CommitmentBody struct {
	PollID       string `json:"pollID" example:"0x12345678901234567890123456789012"`
	VoterAddress string `json:"voterAddress" example:"0x11223344556677889900"`
	CommitHash   string `json:"commitHash" example:"0x12345678909876543210123456789012"`
	VoteOption   int64  `json:"voteOption" example:"1"`
	Salt         int64  `json:"salt" example:"5866984321541876564"`
}

type Response struct {
	Message string `json:"message" example:"vote will be revealed when voting closes"`
}

// Swagger documentation
// @Summary Store a commitment privately, to ensure it can be revealed at a later date
// @Description Save a vote and the matching hash commitment to that vote
// @ID store-commitment
// @Produce json
// @Param payload body database.CommitmentBody true "The (about to be) committed vote details data you would like to store"
// @Router /commitments/ [post]
// @Success 200 {object} database.Response "Success"
// @Failure 406 {object} database.Response "Bad payload"
func StoreCommitment(c *gin.Context) {
	var commitmentBody CommitmentBody
	if err := c.BindJSON(&commitmentBody); err != nil {
		c.JSON(http.StatusInternalServerError, Response{"could not decode body: " + err.Error()})
		return
	}

	commitHash := solsha3.SoliditySHA3(
		solsha3.Uint256(big.NewInt(commitmentBody.VoteOption)),
		solsha3.Uint256(big.NewInt(commitmentBody.Salt)),
	)

	calculatedCommitHash := "0x" + hex.EncodeToString(commitHash)
	if calculatedCommitHash != commitmentBody.CommitHash {
		c.JSON(http.StatusNotAcceptable, Response{"could not save commitment: hash of {" + string(commitmentBody.VoteOption) + "," + string(commitmentBody.Salt) + "} is incorrect"})
		return
	}

	log.Printf("Storing commitment: %v %v %d %d (%v)", commitmentBody.PollID, commitmentBody.VoterAddress, commitmentBody.VoteOption, commitmentBody.Salt, commitmentBody.CommitHash)

	commitment := Commitment{
		PollID:       commitmentBody.PollID,
		VoterAddress: commitmentBody.VoterAddress,
		CommitHash:   commitmentBody.CommitHash,
		VoteOption:   uint8(commitmentBody.VoteOption),
		Salt:         uint64(commitmentBody.Salt),
	}

	if dbc := Db.Create(&commitment); dbc.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{"could not save commitment: " + dbc.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{"vote will be revealed when voting closes"})
}

func MarkAsMostRecentlySeen(pollID [32]byte, voterAddress string, commitHash [32]byte) {
	// Mark all commitments for this user as "not the most recent one", then mark this commit hash as the most recent one
	//log.Printf("Marking commitHash %v as most recently seen for %v on poll %v", hexutil.Encode(commitHash[:]), strings.ToLower(voterAddress), hexutil.Encode(pollID[:]))
	var c Commitment
	Db.Model(&c).Where("poll_id = ? and lower(voter_address) = ?", hexutil.Encode(pollID[:]), strings.ToLower(voterAddress)).Update("last_on_chain", false)
	Db.Model(&c).Where("poll_id = ? and lower(voter_address) = ? and commit_hash = ?", hexutil.Encode(pollID[:]), strings.ToLower(voterAddress), hexutil.Encode(commitHash[:])).Update("last_on_chain", true)
}

func SoftDeleteRevealed(pollID [32]byte, voterAddress string) {
	// We make use of gorm's "soft delete"
	// Records are not really deleted, but flagged as such and ignored in queries
	if voterAddress == "" {
		Db.Where("poll_id = ?", hexutil.Encode(pollID[:])).Delete(Commitment{})
	} else {
		Db.Where("poll_id = ? and voter_address = ?", hexutil.Encode(pollID[:]), voterAddress).Delete(Commitment{})
	}
}
