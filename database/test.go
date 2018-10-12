package database

import (
	//"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/jinzhu/gorm/dialects/postgres"
)

func TestSetupData() {
	// Some data for testing. TODO: Delete once log extraction is working.
	a0 := common.HexToAddress("0x0011223344556677889900112233445566778899")
	a1 := common.HexToAddress("0x0011111111111111111111111111111111111111")
	a2 := common.HexToAddress("0x0022222222222222222222222222222222222222")
	a3 := common.HexToAddress("0x0033333333333333333333333333333333333333")

	c1 := Commitment{PollID: "Poll1", VoterAddress: a0.Hex(), CommitHash: "1234", VoteOption: 2, Salt: 666}
	c2 := Commitment{PollID: "Poll1", VoterAddress: a1.Hex(), CommitHash: "4321", VoteOption: 2, Salt: 666}
	c3 := Commitment{PollID: "Poll1", VoterAddress: a2.Hex(), CommitHash: "6789", VoteOption: 1, Salt: 666}
	c4 := Commitment{PollID: "Poll1", VoterAddress: a3.Hex(), CommitHash: "9876", VoteOption: 1, Salt: 666}

	Db.Create(&c1)
	Db.Create(&c2)
	Db.Create(&c3)
	Db.Create(&c4)
}
