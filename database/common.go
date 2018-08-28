package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type Commitment struct {
	PollID        string
	CreatedAt time.Time
	VoterAddress      string
	CommitHash   string
  VoteOption  uint8
	Salt   uint64
}

func InitDb(new_db *gorm.DB) {
	Db = new_db

	// Migration creates and adds to the database as required, but it won't change columns.
	// If our columns change we must manually drop the affected tables.
	Db.AutoMigrate(&Commitment{})
}
