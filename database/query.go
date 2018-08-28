package database

import (
	//"fmt"
  //"strconv"
  "net/http"
  //"log"

  //"github.com/TruSet/RevealerAPI/database/common"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
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

  // TODO verify valid log
  //c.AbortWithStatus(404)

  commitment := Commitment{
    PollID:    commitmentBody.PollID,
    VoterAddress: commitmentBody.VoterAddress,
    CommitHash: commitmentBody.CommitHash,
    VoteOption: commitmentBody.VoteOption,
    Salt: commitmentBody.Salt,
  }
	Db.Debug().Create(&commitment)

  
	//log.Println(fmt.Sprintf("returning with status created %v mode...", *environment))
  c.JSON(http.StatusCreated, gin.H{"message": "vote will be revealed when voting closes"})
}
