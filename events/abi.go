package events

import (
	"fmt"
	"io/ioutil"
  "math/big"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type CommitPeriodHaltedLog struct {
	// N.B. Field names must match those used in the smart contract!
	PollID               string
	HaltedBy common.Address
	Timestamp                *big.Int
}

func (l CommitPeriodHaltedLog) String() string {
	return fmt.Sprintf("{PollID: %s; HaltedBy: %x; Timestamp: %v}", l.PollID, l.HaltedBy, l.Timestamp)
}

var CommitRevealVotingABI abi.ABI

func init() {
	createABIFromPath(&CommitRevealVotingABI, "./events/TruSetCommitRevealVoting.abi")
}

func createABIFromPath(destABI *abi.ABI, _path string) {
	path, _ := filepath.Abs(_path)
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Errorf("Failed to read file at path ", path, ": ", err)
	}

	*destABI, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Errorf("Invalid abi at path ", path, ": ", err)
	}
}
