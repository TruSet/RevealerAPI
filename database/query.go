package database

import (
  "net/http"
  "encoding/hex"
  "math/big"
  "log"

  "github.com/miguelmota/go-solidity-sha3"
	"github.com/gin-gonic/gin"
)

type CommitmentBody struct {
  PollID string `json:"pollID"`
  VoterAddress string `json:"voterAddress"`
  CommitHash string `json:"commitHash"`
  VoteOption int64 `json:"voteOption"`
  Salt int64 `json:"salt"`
}

func StoreCommitment(c *gin.Context) {
  var commitmentBody CommitmentBody
  c.BindJSON(&commitmentBody)

  commitHash := solsha3.SoliditySHA3(
    solsha3.Uint256(big.NewInt(commitmentBody.VoteOption)),
    solsha3.Uint256(big.NewInt(commitmentBody.Salt)),
  )

  calculatedCommitHash := "0x" + hex.EncodeToString(commitHash)
  if (calculatedCommitHash != commitmentBody.CommitHash) {
    c.AbortWithStatus(http.StatusNotAcceptable)
    return
  }

  commitment := Commitment{
    PollID:    commitmentBody.PollID,
    VoterAddress: commitmentBody.VoterAddress,
    CommitHash: commitmentBody.CommitHash,
    VoteOption: uint8(commitmentBody.VoteOption),
    Salt: uint64(commitmentBody.Salt),
  }
  log.Println(Db)
	Db.Debug().Create(&commitment)

  
  c.JSON(http.StatusCreated, gin.H{"message": "vote will be revealed when voting closes"})
}
