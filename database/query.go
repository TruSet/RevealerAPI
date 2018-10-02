package database

import (
  "encoding/hex"
  "log"
  "math/big"
  "net/http"  
  "github.com/gin-gonic/gin"
  "github.com/miguelmota/go-solidity-sha3"
)

type CommitmentBody struct {
  PollID       string `json:"pollID"`
  VoterAddress string `json:"voterAddress"`
  CommitHash   string `json:"commitHash"`
  VoteOption   int64  `json:"voteOption"`
  Salt         int64  `json:"salt"`
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

  log.Printf("Storing commitment: %v %v %d %d (%v)", commitmentBody.PollID, commitmentBody.VoterAddress, commitmentBody.VoteOption, commitmentBody.Salt, commitmentBody.CommitHash)

  commitment := Commitment{
    PollID:       commitmentBody.PollID,
    VoterAddress: commitmentBody.VoterAddress,
    CommitHash:   commitmentBody.CommitHash,
    VoteOption:   uint8(commitmentBody.VoteOption),
    Salt:         uint64(commitmentBody.Salt),
  }
  Db.Create(&commitment)
  //Db.Debug().Create(&commitment)
  
  c.JSON(http.StatusCreated, gin.H{"message": "vote will be revealed when voting closes"})
}
