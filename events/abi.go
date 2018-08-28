package events

import (
	"fmt"
	"io/ioutil"
	//"math/big"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type RevealPeriodStartLog struct {
	// N.B. Field names must match those used in the smart contract!
	Creator              common.Address
	NewInstrumentAddress common.Address
	Ident                string
}

func (l RevealPeriodStartLog) String() string {
	return fmt.Sprintf("{Creator: %x; Address: %x; Name: %v}", l.Creator, l.NewInstrumentAddress, l.Ident)
}

var CommitRevealVotingABI abi.ABI

func init() {
	createABIFromPath(&CommitRevealVotingABI, "./events/CommitRevealVoting.abi")
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
