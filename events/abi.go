package events

import (
	"fmt"
  //"math/big"
	"strings"

  "github.com/TruSet/RevealerAPI/contract"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type RevealPeriodStartedLog struct {
	// N.B. Field names must match those used in the smart contract!
	PollID               string
	InstrumentAddress    common.Address
	DataIdentifier       [32]byte
  PayloadHash          [32]byte
}

func (l RevealPeriodStartedLog) String() string {
	return fmt.Sprintf("{PollID: %x; InstrumentAddress: %x;}", l.PollID, l.InstrumentAddress)
}

var CommitRevealVotingABI abi.ABI

func init() {
	//loadABI(&CommitRevealVotingABI)
  CommitRevealVotingABI, _ = abi.JSON(strings.NewReader(contract.TruSetCommitRevealVotingABI))
}
