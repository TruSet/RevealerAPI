package database

import (
	//"fmt"
  //"strconv"
  "net/http"
  "encoding/hex"

  "github.com/miguelmota/go-solidity-sha3"
	"github.com/gin-gonic/gin"
)

type CommitmentBody struct {
  PollID string `json:"pollID"`
  VoterAddress string `json:"voterAddress"`
  CommitHash string `json:"commitHash"`
  VoteOption uint8 `json:"voteOption"`
  Salt uint64 `json:"salt"`
}

func StoreCommitment(c *gin.Context) {
  var commitmentBody CommitmentBody
  c.BindJSON(&commitmentBody)

  commitHash := solsha3.SoliditySHA3(
    solsha3.Uint256(commitmentBody.VoteOption),
    solsha3.Uint256(commitmentBody.Salt),
  )

  if (hex.EncodeToString(commitHash) != commitmentBody.CommitHash) {
    c.AbortWithStatus(http.StatusNotAcceptable)
    return
  }

  commitment := Commitment{
    PollID:    commitmentBody.PollID,
    VoterAddress: commitmentBody.VoterAddress,
    CommitHash: commitmentBody.CommitHash,
    VoteOption: commitmentBody.VoteOption,
    Salt: commitmentBody.Salt,
  }
	Db.Create(&commitment)

  
  c.JSON(http.StatusCreated, gin.H{"message": "vote will be revealed when voting closes"})
}
