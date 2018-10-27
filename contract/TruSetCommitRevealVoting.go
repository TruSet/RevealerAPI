// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TruSetCommitRevealVotingABI is the input ABI used to generate the binding from.
const TruSetCommitRevealVotingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"commitDeadline\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"stakedForPollID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_AGAINST\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"numStakersForPoll\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"commitPeriodStartedTimestamp\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"getVoteCounts\",\"outputs\":[{\"name\":\"numForVotes\",\"type\":\"uint256\"},{\"name\":\"numAgainstVotes\",\"type\":\"uint256\"},{\"name\":\"numCommittedButNotRevealedVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"voters\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REVEAL_DURATION_IN_SECONDS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rbac\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVote\",\"outputs\":[{\"name\":\"hasVoted\",\"type\":\"bool\"},{\"name\":\"hasRevealed\",\"type\":\"bool\"},{\"name\":\"vote\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_COMMIT_DURATION_IN_SECONDS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"revealDeadline\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_FOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pollMap\",\"outputs\":[{\"name\":\"commitPeriodStartedAt\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"commitsHaltedAt\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"name\":\"revealsHaltedAt\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"name\":\"votesCommittedButNotRevealed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"revealPeriodStartedTimestamp\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"stakersForPoll\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"instrumentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataIdentifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RevealPeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"StakeReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StakeBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"VoteCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesCommittedButNotRevealed\",\"type\":\"uint256\"}],\"name\":\"VoteRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealDuration\",\"type\":\"uint256\"}],\"name\":\"PollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"haltedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommitPeriodHalted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pollID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"haltedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RevealPeriodHalted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"getPollIdentifier\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instruments\",\"type\":\"address[]\"},{\"name\":\"_dataIdentifiers\",\"type\":\"bytes32[]\"},{\"name\":\"_payloadHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_secretHashes\",\"type\":\"bytes32[]\"}],\"name\":\"commitVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealMyVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"name\":\"_voters\",\"type\":\"address[]\"},{\"name\":\"_voteOptions\",\"type\":\"uint256[]\"},{\"name\":\"_salts\",\"type\":\"uint256[]\"}],\"name\":\"revealVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposer\",\"type\":\"address\"},{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"haltCommitPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"haltRevealPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_instrumentAddress\",\"type\":\"address\"},{\"name\":\"_dataIdentifier\",\"type\":\"bytes32\"},{\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"stakedForProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"returnPollStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"returnAllPollStakes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"bytes32\"}],\"name\":\"burnPollStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TruSetCommitRevealVoting is an auto generated Go binding around an Ethereum contract.
type TruSetCommitRevealVoting struct {
	TruSetCommitRevealVotingCaller     // Read-only binding to the contract
	TruSetCommitRevealVotingTransactor // Write-only binding to the contract
	TruSetCommitRevealVotingFilterer   // Log filterer for contract events
}

// TruSetCommitRevealVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TruSetCommitRevealVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruSetCommitRevealVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TruSetCommitRevealVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruSetCommitRevealVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TruSetCommitRevealVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruSetCommitRevealVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TruSetCommitRevealVotingSession struct {
	Contract     *TruSetCommitRevealVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TruSetCommitRevealVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TruSetCommitRevealVotingCallerSession struct {
	Contract *TruSetCommitRevealVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TruSetCommitRevealVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TruSetCommitRevealVotingTransactorSession struct {
	Contract     *TruSetCommitRevealVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TruSetCommitRevealVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TruSetCommitRevealVotingRaw struct {
	Contract *TruSetCommitRevealVoting // Generic contract binding to access the raw methods on
}

// TruSetCommitRevealVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TruSetCommitRevealVotingCallerRaw struct {
	Contract *TruSetCommitRevealVotingCaller // Generic read-only contract binding to access the raw methods on
}

// TruSetCommitRevealVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TruSetCommitRevealVotingTransactorRaw struct {
	Contract *TruSetCommitRevealVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTruSetCommitRevealVoting creates a new instance of TruSetCommitRevealVoting, bound to a specific deployed contract.
func NewTruSetCommitRevealVoting(address common.Address, backend bind.ContractBackend) (*TruSetCommitRevealVoting, error) {
	contract, err := bindTruSetCommitRevealVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVoting{TruSetCommitRevealVotingCaller: TruSetCommitRevealVotingCaller{contract: contract}, TruSetCommitRevealVotingTransactor: TruSetCommitRevealVotingTransactor{contract: contract}, TruSetCommitRevealVotingFilterer: TruSetCommitRevealVotingFilterer{contract: contract}}, nil
}

// NewTruSetCommitRevealVotingCaller creates a new read-only instance of TruSetCommitRevealVoting, bound to a specific deployed contract.
func NewTruSetCommitRevealVotingCaller(address common.Address, caller bind.ContractCaller) (*TruSetCommitRevealVotingCaller, error) {
	contract, err := bindTruSetCommitRevealVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingCaller{contract: contract}, nil
}

// NewTruSetCommitRevealVotingTransactor creates a new write-only instance of TruSetCommitRevealVoting, bound to a specific deployed contract.
func NewTruSetCommitRevealVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*TruSetCommitRevealVotingTransactor, error) {
	contract, err := bindTruSetCommitRevealVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingTransactor{contract: contract}, nil
}

// NewTruSetCommitRevealVotingFilterer creates a new log filterer instance of TruSetCommitRevealVoting, bound to a specific deployed contract.
func NewTruSetCommitRevealVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*TruSetCommitRevealVotingFilterer, error) {
	contract, err := bindTruSetCommitRevealVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingFilterer{contract: contract}, nil
}

// bindTruSetCommitRevealVoting binds a generic wrapper to an already deployed contract.
func bindTruSetCommitRevealVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TruSetCommitRevealVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TruSetCommitRevealVoting.Contract.TruSetCommitRevealVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.TruSetCommitRevealVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.TruSetCommitRevealVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TruSetCommitRevealVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.contract.Transact(opts, method, params...)
}

// MAXCOMMITDURATIONINSECONDS is a free data retrieval call binding the contract method 0xc34d2732.
//
// Solidity: function MAX_COMMIT_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) MAXCOMMITDURATIONINSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "MAX_COMMIT_DURATION_IN_SECONDS")
	return *ret0, err
}

// MAXCOMMITDURATIONINSECONDS is a free data retrieval call binding the contract method 0xc34d2732.
//
// Solidity: function MAX_COMMIT_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) MAXCOMMITDURATIONINSECONDS() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.MAXCOMMITDURATIONINSECONDS(&_TruSetCommitRevealVoting.CallOpts)
}

// MAXCOMMITDURATIONINSECONDS is a free data retrieval call binding the contract method 0xc34d2732.
//
// Solidity: function MAX_COMMIT_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) MAXCOMMITDURATIONINSECONDS() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.MAXCOMMITDURATIONINSECONDS(&_TruSetCommitRevealVoting.CallOpts)
}

// MAXREVEALDURATIONINSECONDS is a free data retrieval call binding the contract method 0x9234c1fd.
//
// Solidity: function MAX_REVEAL_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) MAXREVEALDURATIONINSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "MAX_REVEAL_DURATION_IN_SECONDS")
	return *ret0, err
}

// MAXREVEALDURATIONINSECONDS is a free data retrieval call binding the contract method 0x9234c1fd.
//
// Solidity: function MAX_REVEAL_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) MAXREVEALDURATIONINSECONDS() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.MAXREVEALDURATIONINSECONDS(&_TruSetCommitRevealVoting.CallOpts)
}

// MAXREVEALDURATIONINSECONDS is a free data retrieval call binding the contract method 0x9234c1fd.
//
// Solidity: function MAX_REVEAL_DURATION_IN_SECONDS() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) MAXREVEALDURATIONINSECONDS() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.MAXREVEALDURATIONINSECONDS(&_TruSetCommitRevealVoting.CallOpts)
}

// VOTEAGAINST is a free data retrieval call binding the contract method 0x5dfd3278.
//
// Solidity: function VOTE_AGAINST() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) VOTEAGAINST(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "VOTE_AGAINST")
	return *ret0, err
}

// VOTEAGAINST is a free data retrieval call binding the contract method 0x5dfd3278.
//
// Solidity: function VOTE_AGAINST() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) VOTEAGAINST() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.VOTEAGAINST(&_TruSetCommitRevealVoting.CallOpts)
}

// VOTEAGAINST is a free data retrieval call binding the contract method 0x5dfd3278.
//
// Solidity: function VOTE_AGAINST() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) VOTEAGAINST() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.VOTEAGAINST(&_TruSetCommitRevealVoting.CallOpts)
}

// VOTEFOR is a free data retrieval call binding the contract method 0xd1a75e0d.
//
// Solidity: function VOTE_FOR() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) VOTEFOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "VOTE_FOR")
	return *ret0, err
}

// VOTEFOR is a free data retrieval call binding the contract method 0xd1a75e0d.
//
// Solidity: function VOTE_FOR() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) VOTEFOR() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.VOTEFOR(&_TruSetCommitRevealVoting.CallOpts)
}

// VOTEFOR is a free data retrieval call binding the contract method 0xd1a75e0d.
//
// Solidity: function VOTE_FOR() constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) VOTEFOR() (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.VOTEFOR(&_TruSetCommitRevealVoting.CallOpts)
}

// CommitDeadline is a free data retrieval call binding the contract method 0x01c2b6f5.
//
// Solidity: function commitDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) CommitDeadline(opts *bind.CallOpts, _pollID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "commitDeadline", _pollID)
	return *ret0, err
}

// CommitDeadline is a free data retrieval call binding the contract method 0x01c2b6f5.
//
// Solidity: function commitDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) CommitDeadline(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.CommitDeadline(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// CommitDeadline is a free data retrieval call binding the contract method 0x01c2b6f5.
//
// Solidity: function commitDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) CommitDeadline(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.CommitDeadline(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0x16879344.
//
// Solidity: function commitPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) CommitPeriodActive(opts *bind.CallOpts, _pollID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "commitPeriodActive", _pollID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0x16879344.
//
// Solidity: function commitPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) CommitPeriodActive(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.CommitPeriodActive(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0x16879344.
//
// Solidity: function commitPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) CommitPeriodActive(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.CommitPeriodActive(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// CommitPeriodStartedTimestamp is a free data retrieval call binding the contract method 0x62aff4d5.
//
// Solidity: function commitPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) CommitPeriodStartedTimestamp(opts *bind.CallOpts, _pollID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "commitPeriodStartedTimestamp", _pollID)
	return *ret0, err
}

// CommitPeriodStartedTimestamp is a free data retrieval call binding the contract method 0x62aff4d5.
//
// Solidity: function commitPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) CommitPeriodStartedTimestamp(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.CommitPeriodStartedTimestamp(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// CommitPeriodStartedTimestamp is a free data retrieval call binding the contract method 0x62aff4d5.
//
// Solidity: function commitPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) CommitPeriodStartedTimestamp(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.CommitPeriodStartedTimestamp(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x9fce2d80.
//
// Solidity: function didCommit(_pollID bytes32, _voter address) constant returns(committed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) DidCommit(opts *bind.CallOpts, _pollID [32]byte, _voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "didCommit", _pollID, _voter)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x9fce2d80.
//
// Solidity: function didCommit(_pollID bytes32, _voter address) constant returns(committed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) DidCommit(_pollID [32]byte, _voter common.Address) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.DidCommit(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// DidCommit is a free data retrieval call binding the contract method 0x9fce2d80.
//
// Solidity: function didCommit(_pollID bytes32, _voter address) constant returns(committed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) DidCommit(_pollID [32]byte, _voter common.Address) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.DidCommit(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// DidReveal is a free data retrieval call binding the contract method 0x1593796e.
//
// Solidity: function didReveal(_pollID bytes32, _voter address) constant returns(revealed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) DidReveal(opts *bind.CallOpts, _pollID [32]byte, _voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "didReveal", _pollID, _voter)
	return *ret0, err
}

// DidReveal is a free data retrieval call binding the contract method 0x1593796e.
//
// Solidity: function didReveal(_pollID bytes32, _voter address) constant returns(revealed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) DidReveal(_pollID [32]byte, _voter common.Address) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.DidReveal(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// DidReveal is a free data retrieval call binding the contract method 0x1593796e.
//
// Solidity: function didReveal(_pollID bytes32, _voter address) constant returns(revealed bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) DidReveal(_pollID [32]byte, _voter common.Address) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.DidReveal(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xc3d4d75a.
//
// Solidity: function getCommitHash(_pollID bytes32, _voter address) constant returns(commitHash bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) GetCommitHash(opts *bind.CallOpts, _pollID [32]byte, _voter common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "getCommitHash", _pollID, _voter)
	return *ret0, err
}

// GetCommitHash is a free data retrieval call binding the contract method 0xc3d4d75a.
//
// Solidity: function getCommitHash(_pollID bytes32, _voter address) constant returns(commitHash bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) GetCommitHash(_pollID [32]byte, _voter common.Address) ([32]byte, error) {
	return _TruSetCommitRevealVoting.Contract.GetCommitHash(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xc3d4d75a.
//
// Solidity: function getCommitHash(_pollID bytes32, _voter address) constant returns(commitHash bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) GetCommitHash(_pollID [32]byte, _voter common.Address) ([32]byte, error) {
	return _TruSetCommitRevealVoting.Contract.GetCommitHash(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// GetPollIdentifier is a free data retrieval call binding the contract method 0x4309b4b2.
//
// Solidity: function getPollIdentifier(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) GetPollIdentifier(opts *bind.CallOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "getPollIdentifier", _instrument, _dataIdentifier, _payloadHash)
	return *ret0, err
}

// GetPollIdentifier is a free data retrieval call binding the contract method 0x4309b4b2.
//
// Solidity: function getPollIdentifier(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) GetPollIdentifier(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) ([32]byte, error) {
	return _TruSetCommitRevealVoting.Contract.GetPollIdentifier(&_TruSetCommitRevealVoting.CallOpts, _instrument, _dataIdentifier, _payloadHash)
}

// GetPollIdentifier is a free data retrieval call binding the contract method 0x4309b4b2.
//
// Solidity: function getPollIdentifier(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) GetPollIdentifier(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) ([32]byte, error) {
	return _TruSetCommitRevealVoting.Contract.GetPollIdentifier(&_TruSetCommitRevealVoting.CallOpts, _instrument, _dataIdentifier, _payloadHash)
}

// GetVote is a free data retrieval call binding the contract method 0xb3e7c2bd.
//
// Solidity: function getVote(_pollID bytes32, _voter address) constant returns(hasVoted bool, hasRevealed bool, vote uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) GetVote(opts *bind.CallOpts, _pollID [32]byte, _voter common.Address) (struct {
	HasVoted    bool
	HasRevealed bool
	Vote        *big.Int
}, error) {
	ret := new(struct {
		HasVoted    bool
		HasRevealed bool
		Vote        *big.Int
	})
	out := ret
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "getVote", _pollID, _voter)
	return *ret, err
}

// GetVote is a free data retrieval call binding the contract method 0xb3e7c2bd.
//
// Solidity: function getVote(_pollID bytes32, _voter address) constant returns(hasVoted bool, hasRevealed bool, vote uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) GetVote(_pollID [32]byte, _voter common.Address) (struct {
	HasVoted    bool
	HasRevealed bool
	Vote        *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.GetVote(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// GetVote is a free data retrieval call binding the contract method 0xb3e7c2bd.
//
// Solidity: function getVote(_pollID bytes32, _voter address) constant returns(hasVoted bool, hasRevealed bool, vote uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) GetVote(_pollID [32]byte, _voter common.Address) (struct {
	HasVoted    bool
	HasRevealed bool
	Vote        *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.GetVote(&_TruSetCommitRevealVoting.CallOpts, _pollID, _voter)
}

// GetVoteCounts is a free data retrieval call binding the contract method 0x782fb5d4.
//
// Solidity: function getVoteCounts(_pollID bytes32) constant returns(numForVotes uint256, numAgainstVotes uint256, numCommittedButNotRevealedVotes uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) GetVoteCounts(opts *bind.CallOpts, _pollID [32]byte) (struct {
	NumForVotes                     *big.Int
	NumAgainstVotes                 *big.Int
	NumCommittedButNotRevealedVotes *big.Int
}, error) {
	ret := new(struct {
		NumForVotes                     *big.Int
		NumAgainstVotes                 *big.Int
		NumCommittedButNotRevealedVotes *big.Int
	})
	out := ret
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "getVoteCounts", _pollID)
	return *ret, err
}

// GetVoteCounts is a free data retrieval call binding the contract method 0x782fb5d4.
//
// Solidity: function getVoteCounts(_pollID bytes32) constant returns(numForVotes uint256, numAgainstVotes uint256, numCommittedButNotRevealedVotes uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) GetVoteCounts(_pollID [32]byte) (struct {
	NumForVotes                     *big.Int
	NumAgainstVotes                 *big.Int
	NumCommittedButNotRevealedVotes *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.GetVoteCounts(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// GetVoteCounts is a free data retrieval call binding the contract method 0x782fb5d4.
//
// Solidity: function getVoteCounts(_pollID bytes32) constant returns(numForVotes uint256, numAgainstVotes uint256, numCommittedButNotRevealedVotes uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) GetVoteCounts(_pollID [32]byte) (struct {
	NumForVotes                     *big.Int
	NumAgainstVotes                 *big.Int
	NumCommittedButNotRevealedVotes *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.GetVoteCounts(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// GetVoters is a free data retrieval call binding the contract method 0x8b4a781c.
//
// Solidity: function getVoters(_pollID bytes32) constant returns(voters address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) GetVoters(opts *bind.CallOpts, _pollID [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "getVoters", _pollID)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x8b4a781c.
//
// Solidity: function getVoters(_pollID bytes32) constant returns(voters address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) GetVoters(_pollID [32]byte) ([]common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.GetVoters(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// GetVoters is a free data retrieval call binding the contract method 0x8b4a781c.
//
// Solidity: function getVoters(_pollID bytes32) constant returns(voters address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) GetVoters(_pollID [32]byte) ([]common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.GetVoters(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) IsExpired(opts *bind.CallOpts, _terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "isExpired", _terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.IsExpired(&_TruSetCommitRevealVoting.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.IsExpired(&_TruSetCommitRevealVoting.CallOpts, _terminationDate)
}

// NumStakersForPoll is a free data retrieval call binding the contract method 0x5eface14.
//
// Solidity: function numStakersForPoll(_pollID bytes32) constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) NumStakersForPoll(opts *bind.CallOpts, _pollID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "numStakersForPoll", _pollID)
	return *ret0, err
}

// NumStakersForPoll is a free data retrieval call binding the contract method 0x5eface14.
//
// Solidity: function numStakersForPoll(_pollID bytes32) constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) NumStakersForPoll(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.NumStakersForPoll(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// NumStakersForPoll is a free data retrieval call binding the contract method 0x5eface14.
//
// Solidity: function numStakersForPoll(_pollID bytes32) constant returns(uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) NumStakersForPoll(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.NumStakersForPoll(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xfd2cee78.
//
// Solidity: function pollEnded(_pollID bytes32) constant returns(ended bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) PollEnded(opts *bind.CallOpts, _pollID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "pollEnded", _pollID)
	return *ret0, err
}

// PollEnded is a free data retrieval call binding the contract method 0xfd2cee78.
//
// Solidity: function pollEnded(_pollID bytes32) constant returns(ended bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) PollEnded(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.PollEnded(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xfd2cee78.
//
// Solidity: function pollEnded(_pollID bytes32) constant returns(ended bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) PollEnded(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.PollEnded(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x0dcf298b.
//
// Solidity: function pollExists(_pollID bytes32) constant returns(exists bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) PollExists(opts *bind.CallOpts, _pollID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "pollExists", _pollID)
	return *ret0, err
}

// PollExists is a free data retrieval call binding the contract method 0x0dcf298b.
//
// Solidity: function pollExists(_pollID bytes32) constant returns(exists bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) PollExists(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.PollExists(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x0dcf298b.
//
// Solidity: function pollExists(_pollID bytes32) constant returns(exists bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) PollExists(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.PollExists(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// PollMap is a free data retrieval call binding the contract method 0xd2b77264.
//
// Solidity: function pollMap( bytes32) constant returns(commitPeriodStartedAt uint256, commitDuration uint256, commitsHaltedAt uint256, revealDuration uint256, revealsHaltedAt uint256, votesFor uint256, votesAgainst uint256, votesCommittedButNotRevealed uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) PollMap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CommitPeriodStartedAt        *big.Int
	CommitDuration               *big.Int
	CommitsHaltedAt              *big.Int
	RevealDuration               *big.Int
	RevealsHaltedAt              *big.Int
	VotesFor                     *big.Int
	VotesAgainst                 *big.Int
	VotesCommittedButNotRevealed *big.Int
}, error) {
	ret := new(struct {
		CommitPeriodStartedAt        *big.Int
		CommitDuration               *big.Int
		CommitsHaltedAt              *big.Int
		RevealDuration               *big.Int
		RevealsHaltedAt              *big.Int
		VotesFor                     *big.Int
		VotesAgainst                 *big.Int
		VotesCommittedButNotRevealed *big.Int
	})
	out := ret
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "pollMap", arg0)
	return *ret, err
}

// PollMap is a free data retrieval call binding the contract method 0xd2b77264.
//
// Solidity: function pollMap( bytes32) constant returns(commitPeriodStartedAt uint256, commitDuration uint256, commitsHaltedAt uint256, revealDuration uint256, revealsHaltedAt uint256, votesFor uint256, votesAgainst uint256, votesCommittedButNotRevealed uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) PollMap(arg0 [32]byte) (struct {
	CommitPeriodStartedAt        *big.Int
	CommitDuration               *big.Int
	CommitsHaltedAt              *big.Int
	RevealDuration               *big.Int
	RevealsHaltedAt              *big.Int
	VotesFor                     *big.Int
	VotesAgainst                 *big.Int
	VotesCommittedButNotRevealed *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.PollMap(&_TruSetCommitRevealVoting.CallOpts, arg0)
}

// PollMap is a free data retrieval call binding the contract method 0xd2b77264.
//
// Solidity: function pollMap( bytes32) constant returns(commitPeriodStartedAt uint256, commitDuration uint256, commitsHaltedAt uint256, revealDuration uint256, revealsHaltedAt uint256, votesFor uint256, votesAgainst uint256, votesCommittedButNotRevealed uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) PollMap(arg0 [32]byte) (struct {
	CommitPeriodStartedAt        *big.Int
	CommitDuration               *big.Int
	CommitsHaltedAt              *big.Int
	RevealDuration               *big.Int
	RevealsHaltedAt              *big.Int
	VotesFor                     *big.Int
	VotesAgainst                 *big.Int
	VotesCommittedButNotRevealed *big.Int
}, error) {
	return _TruSetCommitRevealVoting.Contract.PollMap(&_TruSetCommitRevealVoting.CallOpts, arg0)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) Rbac(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "rbac")
	return *ret0, err
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) Rbac() (common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.Rbac(&_TruSetCommitRevealVoting.CallOpts)
}

// Rbac is a free data retrieval call binding the contract method 0xa8ecc7f1.
//
// Solidity: function rbac() constant returns(address)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) Rbac() (common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.Rbac(&_TruSetCommitRevealVoting.CallOpts)
}

// RevealDeadline is a free data retrieval call binding the contract method 0xcb27bdc1.
//
// Solidity: function revealDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) RevealDeadline(opts *bind.CallOpts, _pollID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "revealDeadline", _pollID)
	return *ret0, err
}

// RevealDeadline is a free data retrieval call binding the contract method 0xcb27bdc1.
//
// Solidity: function revealDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealDeadline(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.RevealDeadline(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// RevealDeadline is a free data retrieval call binding the contract method 0xcb27bdc1.
//
// Solidity: function revealDeadline(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) RevealDeadline(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.RevealDeadline(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x28fbf28a.
//
// Solidity: function revealPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) RevealPeriodActive(opts *bind.CallOpts, _pollID [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "revealPeriodActive", _pollID)
	return *ret0, err
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x28fbf28a.
//
// Solidity: function revealPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealPeriodActive(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.RevealPeriodActive(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x28fbf28a.
//
// Solidity: function revealPeriodActive(_pollID bytes32) constant returns(active bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) RevealPeriodActive(_pollID [32]byte) (bool, error) {
	return _TruSetCommitRevealVoting.Contract.RevealPeriodActive(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// RevealPeriodStartedTimestamp is a free data retrieval call binding the contract method 0xe2a5c39f.
//
// Solidity: function revealPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) RevealPeriodStartedTimestamp(opts *bind.CallOpts, _pollID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "revealPeriodStartedTimestamp", _pollID)
	return *ret0, err
}

// RevealPeriodStartedTimestamp is a free data retrieval call binding the contract method 0xe2a5c39f.
//
// Solidity: function revealPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealPeriodStartedTimestamp(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.RevealPeriodStartedTimestamp(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// RevealPeriodStartedTimestamp is a free data retrieval call binding the contract method 0xe2a5c39f.
//
// Solidity: function revealPeriodStartedTimestamp(_pollID bytes32) constant returns(timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) RevealPeriodStartedTimestamp(_pollID [32]byte) (*big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.RevealPeriodStartedTimestamp(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// StakedForPollID is a free data retrieval call binding the contract method 0x3c2dba23.
//
// Solidity: function stakedForPollID(_user address, _pollID bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) StakedForPollID(opts *bind.CallOpts, _user common.Address, _pollID [32]byte) (uint8, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "stakedForPollID", _user, _pollID)
	return *ret0, *ret1, err
}

// StakedForPollID is a free data retrieval call binding the contract method 0x3c2dba23.
//
// Solidity: function stakedForPollID(_user address, _pollID bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) StakedForPollID(_user common.Address, _pollID [32]byte) (uint8, *big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.StakedForPollID(&_TruSetCommitRevealVoting.CallOpts, _user, _pollID)
}

// StakedForPollID is a free data retrieval call binding the contract method 0x3c2dba23.
//
// Solidity: function stakedForPollID(_user address, _pollID bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) StakedForPollID(_user common.Address, _pollID [32]byte) (uint8, *big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.StakedForPollID(&_TruSetCommitRevealVoting.CallOpts, _user, _pollID)
}

// StakedForProposal is a free data retrieval call binding the contract method 0xbe6c03ff.
//
// Solidity: function stakedForProposal(_user address, _instrumentAddress address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) StakedForProposal(opts *bind.CallOpts, _user common.Address, _instrumentAddress common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (uint8, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "stakedForProposal", _user, _instrumentAddress, _dataIdentifier, _payloadHash)
	return *ret0, *ret1, err
}

// StakedForProposal is a free data retrieval call binding the contract method 0xbe6c03ff.
//
// Solidity: function stakedForProposal(_user address, _instrumentAddress address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) StakedForProposal(_user common.Address, _instrumentAddress common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (uint8, *big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.StakedForProposal(&_TruSetCommitRevealVoting.CallOpts, _user, _instrumentAddress, _dataIdentifier, _payloadHash)
}

// StakedForProposal is a free data retrieval call binding the contract method 0xbe6c03ff.
//
// Solidity: function stakedForProposal(_user address, _instrumentAddress address, _dataIdentifier bytes32, _payloadHash bytes32) constant returns(uint8, uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) StakedForProposal(_user common.Address, _instrumentAddress common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (uint8, *big.Int, error) {
	return _TruSetCommitRevealVoting.Contract.StakedForProposal(&_TruSetCommitRevealVoting.CallOpts, _user, _instrumentAddress, _dataIdentifier, _payloadHash)
}

// StakersForPoll is a free data retrieval call binding the contract method 0xef5850c7.
//
// Solidity: function stakersForPoll(_pollID bytes32) constant returns(address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCaller) StakersForPoll(opts *bind.CallOpts, _pollID [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TruSetCommitRevealVoting.contract.Call(opts, out, "stakersForPoll", _pollID)
	return *ret0, err
}

// StakersForPoll is a free data retrieval call binding the contract method 0xef5850c7.
//
// Solidity: function stakersForPoll(_pollID bytes32) constant returns(address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) StakersForPoll(_pollID [32]byte) ([]common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.StakersForPoll(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// StakersForPoll is a free data retrieval call binding the contract method 0xef5850c7.
//
// Solidity: function stakersForPoll(_pollID bytes32) constant returns(address[])
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingCallerSession) StakersForPoll(_pollID [32]byte) ([]common.Address, error) {
	return _TruSetCommitRevealVoting.Contract.StakersForPoll(&_TruSetCommitRevealVoting.CallOpts, _pollID)
}

// BurnPollStake is a paid mutator transaction binding the contract method 0x52f804a8.
//
// Solidity: function burnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) BurnPollStake(opts *bind.TransactOpts, _user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "burnPollStake", _user, _pollID)
}

// BurnPollStake is a paid mutator transaction binding the contract method 0x52f804a8.
//
// Solidity: function burnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) BurnPollStake(_user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.BurnPollStake(&_TruSetCommitRevealVoting.TransactOpts, _user, _pollID)
}

// BurnPollStake is a paid mutator transaction binding the contract method 0x52f804a8.
//
// Solidity: function burnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) BurnPollStake(_user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.BurnPollStake(&_TruSetCommitRevealVoting.TransactOpts, _user, _pollID)
}

// CommitVote is a paid mutator transaction binding the contract method 0x2ab24e3b.
//
// Solidity: function commitVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _secretHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) CommitVote(opts *bind.TransactOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _secretHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "commitVote", _instrument, _dataIdentifier, _payloadHash, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x2ab24e3b.
//
// Solidity: function commitVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _secretHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) CommitVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _secretHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.CommitVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x2ab24e3b.
//
// Solidity: function commitVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _secretHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) CommitVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _secretHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.CommitVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _secretHash)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x2726f061.
//
// Solidity: function commitVotes(_instruments address[], _dataIdentifiers bytes32[], _payloadHashes bytes32[], _secretHashes bytes32[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) CommitVotes(opts *bind.TransactOpts, _instruments []common.Address, _dataIdentifiers [][32]byte, _payloadHashes [][32]byte, _secretHashes [][32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "commitVotes", _instruments, _dataIdentifiers, _payloadHashes, _secretHashes)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x2726f061.
//
// Solidity: function commitVotes(_instruments address[], _dataIdentifiers bytes32[], _payloadHashes bytes32[], _secretHashes bytes32[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) CommitVotes(_instruments []common.Address, _dataIdentifiers [][32]byte, _payloadHashes [][32]byte, _secretHashes [][32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.CommitVotes(&_TruSetCommitRevealVoting.TransactOpts, _instruments, _dataIdentifiers, _payloadHashes, _secretHashes)
}

// CommitVotes is a paid mutator transaction binding the contract method 0x2726f061.
//
// Solidity: function commitVotes(_instruments address[], _dataIdentifiers bytes32[], _payloadHashes bytes32[], _secretHashes bytes32[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) CommitVotes(_instruments []common.Address, _dataIdentifiers [][32]byte, _payloadHashes [][32]byte, _secretHashes [][32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.CommitVotes(&_TruSetCommitRevealVoting.TransactOpts, _instruments, _dataIdentifiers, _payloadHashes, _secretHashes)
}

// HaltCommitPeriod is a paid mutator transaction binding the contract method 0xf3c957f2.
//
// Solidity: function haltCommitPeriod(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) HaltCommitPeriod(opts *bind.TransactOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "haltCommitPeriod", _instrument, _dataIdentifier, _payloadHash)
}

// HaltCommitPeriod is a paid mutator transaction binding the contract method 0xf3c957f2.
//
// Solidity: function haltCommitPeriod(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) HaltCommitPeriod(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.HaltCommitPeriod(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash)
}

// HaltCommitPeriod is a paid mutator transaction binding the contract method 0xf3c957f2.
//
// Solidity: function haltCommitPeriod(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) HaltCommitPeriod(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.HaltCommitPeriod(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash)
}

// HaltRevealPeriod is a paid mutator transaction binding the contract method 0xb8847e9d.
//
// Solidity: function haltRevealPeriod(_pollID bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) HaltRevealPeriod(opts *bind.TransactOpts, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "haltRevealPeriod", _pollID)
}

// HaltRevealPeriod is a paid mutator transaction binding the contract method 0xb8847e9d.
//
// Solidity: function haltRevealPeriod(_pollID bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) HaltRevealPeriod(_pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.HaltRevealPeriod(&_TruSetCommitRevealVoting.TransactOpts, _pollID)
}

// HaltRevealPeriod is a paid mutator transaction binding the contract method 0xb8847e9d.
//
// Solidity: function haltRevealPeriod(_pollID bytes32) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) HaltRevealPeriod(_pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.HaltRevealPeriod(&_TruSetCommitRevealVoting.TransactOpts, _pollID)
}

// ReturnAllPollStakes is a paid mutator transaction binding the contract method 0x4f88cc3c.
//
// Solidity: function returnAllPollStakes(_pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) ReturnAllPollStakes(opts *bind.TransactOpts, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "returnAllPollStakes", _pollID)
}

// ReturnAllPollStakes is a paid mutator transaction binding the contract method 0x4f88cc3c.
//
// Solidity: function returnAllPollStakes(_pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) ReturnAllPollStakes(_pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.ReturnAllPollStakes(&_TruSetCommitRevealVoting.TransactOpts, _pollID)
}

// ReturnAllPollStakes is a paid mutator transaction binding the contract method 0x4f88cc3c.
//
// Solidity: function returnAllPollStakes(_pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) ReturnAllPollStakes(_pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.ReturnAllPollStakes(&_TruSetCommitRevealVoting.TransactOpts, _pollID)
}

// ReturnPollStake is a paid mutator transaction binding the contract method 0x973882e9.
//
// Solidity: function returnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) ReturnPollStake(opts *bind.TransactOpts, _user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "returnPollStake", _user, _pollID)
}

// ReturnPollStake is a paid mutator transaction binding the contract method 0x973882e9.
//
// Solidity: function returnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) ReturnPollStake(_user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.ReturnPollStake(&_TruSetCommitRevealVoting.TransactOpts, _user, _pollID)
}

// ReturnPollStake is a paid mutator transaction binding the contract method 0x973882e9.
//
// Solidity: function returnPollStake(_user address, _pollID bytes32) returns(bool)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) ReturnPollStake(_user common.Address, _pollID [32]byte) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.ReturnPollStake(&_TruSetCommitRevealVoting.TransactOpts, _user, _pollID)
}

// RevealMyVote is a paid mutator transaction binding the contract method 0x04c64e25.
//
// Solidity: function revealMyVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) RevealMyVote(opts *bind.TransactOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "revealMyVote", _instrument, _dataIdentifier, _payloadHash, _voteOption, _salt)
}

// RevealMyVote is a paid mutator transaction binding the contract method 0x04c64e25.
//
// Solidity: function revealMyVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealMyVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealMyVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voteOption, _salt)
}

// RevealMyVote is a paid mutator transaction binding the contract method 0x04c64e25.
//
// Solidity: function revealMyVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) RevealMyVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealMyVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x6a27c41d.
//
// Solidity: function revealVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voter address, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) RevealVote(opts *bind.TransactOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voter common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "revealVote", _instrument, _dataIdentifier, _payloadHash, _voter, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x6a27c41d.
//
// Solidity: function revealVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voter address, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voter common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voter, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x6a27c41d.
//
// Solidity: function revealVote(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voter address, _voteOption uint256, _salt uint256) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) RevealVote(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voter common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealVote(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voter, _voteOption, _salt)
}

// RevealVotes is a paid mutator transaction binding the contract method 0xf9191b18.
//
// Solidity: function revealVotes(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voters address[], _voteOptions uint256[], _salts uint256[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) RevealVotes(opts *bind.TransactOpts, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voters []common.Address, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "revealVotes", _instrument, _dataIdentifier, _payloadHash, _voters, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0xf9191b18.
//
// Solidity: function revealVotes(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voters address[], _voteOptions uint256[], _salts uint256[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) RevealVotes(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voters []common.Address, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealVotes(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voters, _voteOptions, _salts)
}

// RevealVotes is a paid mutator transaction binding the contract method 0xf9191b18.
//
// Solidity: function revealVotes(_instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _voters address[], _voteOptions uint256[], _salts uint256[]) returns()
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) RevealVotes(_instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _voters []common.Address, _voteOptions []*big.Int, _salts []*big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.RevealVotes(&_TruSetCommitRevealVoting.TransactOpts, _instrument, _dataIdentifier, _payloadHash, _voters, _voteOptions, _salts)
}

// StartPoll is a paid mutator transaction binding the contract method 0xbb2ced25.
//
// Solidity: function startPoll(_proposer address, _instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _commitDuration uint256, _revealDuration uint256) returns(_pollID bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactor) StartPoll(opts *bind.TransactOpts, _proposer common.Address, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.contract.Transact(opts, "startPoll", _proposer, _instrument, _dataIdentifier, _payloadHash, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0xbb2ced25.
//
// Solidity: function startPoll(_proposer address, _instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _commitDuration uint256, _revealDuration uint256) returns(_pollID bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingSession) StartPoll(_proposer common.Address, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.StartPoll(&_TruSetCommitRevealVoting.TransactOpts, _proposer, _instrument, _dataIdentifier, _payloadHash, _commitDuration, _revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0xbb2ced25.
//
// Solidity: function startPoll(_proposer address, _instrument address, _dataIdentifier bytes32, _payloadHash bytes32, _commitDuration uint256, _revealDuration uint256) returns(_pollID bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingTransactorSession) StartPoll(_proposer common.Address, _instrument common.Address, _dataIdentifier [32]byte, _payloadHash [32]byte, _commitDuration *big.Int, _revealDuration *big.Int) (*types.Transaction, error) {
	return _TruSetCommitRevealVoting.Contract.StartPoll(&_TruSetCommitRevealVoting.TransactOpts, _proposer, _instrument, _dataIdentifier, _payloadHash, _commitDuration, _revealDuration)
}

// TruSetCommitRevealVotingCommitPeriodHaltedIterator is returned from FilterCommitPeriodHalted and is used to iterate over the raw logs and unpacked data for CommitPeriodHalted events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingCommitPeriodHaltedIterator struct {
	Event *TruSetCommitRevealVotingCommitPeriodHalted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingCommitPeriodHaltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingCommitPeriodHalted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingCommitPeriodHalted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingCommitPeriodHaltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingCommitPeriodHaltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingCommitPeriodHalted represents a CommitPeriodHalted event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingCommitPeriodHalted struct {
	PollID    [32]byte
	HaltedBy  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommitPeriodHalted is a free log retrieval operation binding the contract event 0xc0f2b57ffd01c411e41e4bb32ef5241f1d5eb3fce828f5ee820a323ac2e45193.
//
// Solidity: e CommitPeriodHalted(pollID indexed bytes32, haltedBy address, timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterCommitPeriodHalted(opts *bind.FilterOpts, pollID [][32]byte) (*TruSetCommitRevealVotingCommitPeriodHaltedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "CommitPeriodHalted", pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingCommitPeriodHaltedIterator{contract: _TruSetCommitRevealVoting.contract, event: "CommitPeriodHalted", logs: logs, sub: sub}, nil
}

// WatchCommitPeriodHalted is a free log subscription operation binding the contract event 0xc0f2b57ffd01c411e41e4bb32ef5241f1d5eb3fce828f5ee820a323ac2e45193.
//
// Solidity: e CommitPeriodHalted(pollID indexed bytes32, haltedBy address, timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchCommitPeriodHalted(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingCommitPeriodHalted, pollID [][32]byte) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "CommitPeriodHalted", pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingCommitPeriodHalted)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "CommitPeriodHalted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingPollCreatedIterator is returned from FilterPollCreated and is used to iterate over the raw logs and unpacked data for PollCreated events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingPollCreatedIterator struct {
	Event *TruSetCommitRevealVotingPollCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingPollCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingPollCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingPollCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingPollCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingPollCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingPollCreated represents a PollCreated event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingPollCreated struct {
	PollID         [32]byte
	Creator        common.Address
	CommitDuration *big.Int
	RevealDuration *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPollCreated is a free log retrieval operation binding the contract event 0xb63cd7eae18e725657860cbbf12d757b6e516f142db7fd8cabd61fb5d93c24e9.
//
// Solidity: e PollCreated(pollID indexed bytes32, creator address, commitDuration uint256, revealDuration uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterPollCreated(opts *bind.FilterOpts, pollID [][32]byte) (*TruSetCommitRevealVotingPollCreatedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "PollCreated", pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingPollCreatedIterator{contract: _TruSetCommitRevealVoting.contract, event: "PollCreated", logs: logs, sub: sub}, nil
}

// WatchPollCreated is a free log subscription operation binding the contract event 0xb63cd7eae18e725657860cbbf12d757b6e516f142db7fd8cabd61fb5d93c24e9.
//
// Solidity: e PollCreated(pollID indexed bytes32, creator address, commitDuration uint256, revealDuration uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchPollCreated(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingPollCreated, pollID [][32]byte) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "PollCreated", pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingPollCreated)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "PollCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingRevealPeriodHaltedIterator is returned from FilterRevealPeriodHalted and is used to iterate over the raw logs and unpacked data for RevealPeriodHalted events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingRevealPeriodHaltedIterator struct {
	Event *TruSetCommitRevealVotingRevealPeriodHalted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingRevealPeriodHaltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingRevealPeriodHalted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingRevealPeriodHalted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingRevealPeriodHaltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingRevealPeriodHaltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingRevealPeriodHalted represents a RevealPeriodHalted event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingRevealPeriodHalted struct {
	PollID    [32]byte
	HaltedBy  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealPeriodHalted is a free log retrieval operation binding the contract event 0x2e3281c9d2188bb5bad0561e502ceb5224998ab2b0a49af0a123cc044c0a1b6a.
//
// Solidity: e RevealPeriodHalted(pollID indexed bytes32, haltedBy address, timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterRevealPeriodHalted(opts *bind.FilterOpts, pollID [][32]byte) (*TruSetCommitRevealVotingRevealPeriodHaltedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "RevealPeriodHalted", pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingRevealPeriodHaltedIterator{contract: _TruSetCommitRevealVoting.contract, event: "RevealPeriodHalted", logs: logs, sub: sub}, nil
}

// WatchRevealPeriodHalted is a free log subscription operation binding the contract event 0x2e3281c9d2188bb5bad0561e502ceb5224998ab2b0a49af0a123cc044c0a1b6a.
//
// Solidity: e RevealPeriodHalted(pollID indexed bytes32, haltedBy address, timestamp uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchRevealPeriodHalted(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingRevealPeriodHalted, pollID [][32]byte) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "RevealPeriodHalted", pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingRevealPeriodHalted)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "RevealPeriodHalted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingRevealPeriodStartedIterator is returned from FilterRevealPeriodStarted and is used to iterate over the raw logs and unpacked data for RevealPeriodStarted events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingRevealPeriodStartedIterator struct {
	Event *TruSetCommitRevealVotingRevealPeriodStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingRevealPeriodStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingRevealPeriodStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingRevealPeriodStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingRevealPeriodStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingRevealPeriodStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingRevealPeriodStarted represents a RevealPeriodStarted event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingRevealPeriodStarted struct {
	PollID            [32]byte
	InstrumentAddress common.Address
	DataIdentifier    [32]byte
	PayloadHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRevealPeriodStarted is a free log retrieval operation binding the contract event 0x832d41f3f913a16d24e73edf9c864aa51505ab5fe003cc125fe8877947a1aa9f.
//
// Solidity: e RevealPeriodStarted(pollID indexed bytes32, instrumentAddress indexed address, dataIdentifier bytes32, payloadHash bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterRevealPeriodStarted(opts *bind.FilterOpts, pollID [][32]byte, instrumentAddress []common.Address) (*TruSetCommitRevealVotingRevealPeriodStartedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var instrumentAddressRule []interface{}
	for _, instrumentAddressItem := range instrumentAddress {
		instrumentAddressRule = append(instrumentAddressRule, instrumentAddressItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "RevealPeriodStarted", pollIDRule, instrumentAddressRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingRevealPeriodStartedIterator{contract: _TruSetCommitRevealVoting.contract, event: "RevealPeriodStarted", logs: logs, sub: sub}, nil
}

// WatchRevealPeriodStarted is a free log subscription operation binding the contract event 0x832d41f3f913a16d24e73edf9c864aa51505ab5fe003cc125fe8877947a1aa9f.
//
// Solidity: e RevealPeriodStarted(pollID indexed bytes32, instrumentAddress indexed address, dataIdentifier bytes32, payloadHash bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchRevealPeriodStarted(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingRevealPeriodStarted, pollID [][32]byte, instrumentAddress []common.Address) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var instrumentAddressRule []interface{}
	for _, instrumentAddressItem := range instrumentAddress {
		instrumentAddressRule = append(instrumentAddressRule, instrumentAddressItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "RevealPeriodStarted", pollIDRule, instrumentAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingRevealPeriodStarted)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "RevealPeriodStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingStakeBurntIterator is returned from FilterStakeBurnt and is used to iterate over the raw logs and unpacked data for StakeBurnt events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStakeBurntIterator struct {
	Event *TruSetCommitRevealVotingStakeBurnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingStakeBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingStakeBurnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingStakeBurnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingStakeBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingStakeBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingStakeBurnt represents a StakeBurnt event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStakeBurnt struct {
	User   common.Address
	Amount *big.Int
	PollID [32]byte
	Action uint8
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeBurnt is a free log retrieval operation binding the contract event 0x98b9e629a6a2ee5cbb9ad28c2318b43f55958d4a24a403b03267b512224bea59.
//
// Solidity: e StakeBurnt(user indexed address, amount uint256, pollID indexed bytes32, action uint8, reason string)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterStakeBurnt(opts *bind.FilterOpts, user []common.Address, pollID [][32]byte) (*TruSetCommitRevealVotingStakeBurntIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "StakeBurnt", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingStakeBurntIterator{contract: _TruSetCommitRevealVoting.contract, event: "StakeBurnt", logs: logs, sub: sub}, nil
}

// WatchStakeBurnt is a free log subscription operation binding the contract event 0x98b9e629a6a2ee5cbb9ad28c2318b43f55958d4a24a403b03267b512224bea59.
//
// Solidity: e StakeBurnt(user indexed address, amount uint256, pollID indexed bytes32, action uint8, reason string)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchStakeBurnt(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingStakeBurnt, user []common.Address, pollID [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "StakeBurnt", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingStakeBurnt)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "StakeBurnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingStakeReturnedIterator is returned from FilterStakeReturned and is used to iterate over the raw logs and unpacked data for StakeReturned events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStakeReturnedIterator struct {
	Event *TruSetCommitRevealVotingStakeReturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingStakeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingStakeReturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingStakeReturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingStakeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingStakeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingStakeReturned represents a StakeReturned event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStakeReturned struct {
	User   common.Address
	Amount *big.Int
	PollID [32]byte
	Action uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeReturned is a free log retrieval operation binding the contract event 0x3db9f61f4d7052d2a81173036ad7e9e276cd26a7079a316526549ca581f4b36d.
//
// Solidity: e StakeReturned(user indexed address, amount uint256, pollID indexed bytes32, action uint8)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterStakeReturned(opts *bind.FilterOpts, user []common.Address, pollID [][32]byte) (*TruSetCommitRevealVotingStakeReturnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "StakeReturned", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingStakeReturnedIterator{contract: _TruSetCommitRevealVoting.contract, event: "StakeReturned", logs: logs, sub: sub}, nil
}

// WatchStakeReturned is a free log subscription operation binding the contract event 0x3db9f61f4d7052d2a81173036ad7e9e276cd26a7079a316526549ca581f4b36d.
//
// Solidity: e StakeReturned(user indexed address, amount uint256, pollID indexed bytes32, action uint8)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchStakeReturned(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingStakeReturned, user []common.Address, pollID [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "StakeReturned", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingStakeReturned)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "StakeReturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStakedIterator struct {
	Event *TruSetCommitRevealVotingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingStaked represents a Staked event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingStaked struct {
	User   common.Address
	Amount *big.Int
	PollID [32]byte
	Action uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xffcaaa0c4ca25b5827074c948fc6acd7c25db62ecf524f4469ff88dfbd8346a6.
//
// Solidity: e Staked(user indexed address, amount uint256, pollID indexed bytes32, action uint8)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, pollID [][32]byte) (*TruSetCommitRevealVotingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "Staked", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingStakedIterator{contract: _TruSetCommitRevealVoting.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xffcaaa0c4ca25b5827074c948fc6acd7c25db62ecf524f4469ff88dfbd8346a6.
//
// Solidity: e Staked(user indexed address, amount uint256, pollID indexed bytes32, action uint8)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingStaked, user []common.Address, pollID [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "Staked", userRule, pollIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingStaked)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingVoteCommittedIterator is returned from FilterVoteCommitted and is used to iterate over the raw logs and unpacked data for VoteCommitted events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingVoteCommittedIterator struct {
	Event *TruSetCommitRevealVotingVoteCommitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingVoteCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingVoteCommitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingVoteCommitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingVoteCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingVoteCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingVoteCommitted represents a VoteCommitted event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingVoteCommitted struct {
	PollID     [32]byte
	Voter      common.Address
	SecretHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCommitted is a free log retrieval operation binding the contract event 0xd9fd84967d0d7a1e30c4712a111e11619db9beec04a8c48f362b2c1724b6c58a.
//
// Solidity: e VoteCommitted(pollID indexed bytes32, voter indexed address, secretHash indexed bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterVoteCommitted(opts *bind.FilterOpts, pollID [][32]byte, voter []common.Address, secretHash [][32]byte) (*TruSetCommitRevealVotingVoteCommittedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "VoteCommitted", pollIDRule, voterRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingVoteCommittedIterator{contract: _TruSetCommitRevealVoting.contract, event: "VoteCommitted", logs: logs, sub: sub}, nil
}

// WatchVoteCommitted is a free log subscription operation binding the contract event 0xd9fd84967d0d7a1e30c4712a111e11619db9beec04a8c48f362b2c1724b6c58a.
//
// Solidity: e VoteCommitted(pollID indexed bytes32, voter indexed address, secretHash indexed bytes32)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchVoteCommitted(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingVoteCommitted, pollID [][32]byte, voter []common.Address, secretHash [][32]byte) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "VoteCommitted", pollIDRule, voterRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingVoteCommitted)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "VoteCommitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TruSetCommitRevealVotingVoteRevealedIterator is returned from FilterVoteRevealed and is used to iterate over the raw logs and unpacked data for VoteRevealed events raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingVoteRevealedIterator struct {
	Event *TruSetCommitRevealVotingVoteRevealed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TruSetCommitRevealVotingVoteRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruSetCommitRevealVotingVoteRevealed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TruSetCommitRevealVotingVoteRevealed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TruSetCommitRevealVotingVoteRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruSetCommitRevealVotingVoteRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruSetCommitRevealVotingVoteRevealed represents a VoteRevealed event raised by the TruSetCommitRevealVoting contract.
type TruSetCommitRevealVotingVoteRevealed struct {
	PollID                       [32]byte
	SecretHash                   [32]byte
	Choice                       *big.Int
	Voter                        common.Address
	Revealer                     common.Address
	VotesFor                     *big.Int
	VotesAgainst                 *big.Int
	VotesCommittedButNotRevealed *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealed is a free log retrieval operation binding the contract event 0x7656c59ff19441c541f32aebe589dbafba2d1a195276586916a69f3affc6c559.
//
// Solidity: e VoteRevealed(pollID indexed bytes32, secretHash indexed bytes32, choice indexed uint256, voter address, revealer address, votesFor uint256, votesAgainst uint256, votesCommittedButNotRevealed uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) FilterVoteRevealed(opts *bind.FilterOpts, pollID [][32]byte, secretHash [][32]byte, choice []*big.Int) (*TruSetCommitRevealVotingVoteRevealedIterator, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}
	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.FilterLogs(opts, "VoteRevealed", pollIDRule, secretHashRule, choiceRule)
	if err != nil {
		return nil, err
	}
	return &TruSetCommitRevealVotingVoteRevealedIterator{contract: _TruSetCommitRevealVoting.contract, event: "VoteRevealed", logs: logs, sub: sub}, nil
}

// WatchVoteRevealed is a free log subscription operation binding the contract event 0x7656c59ff19441c541f32aebe589dbafba2d1a195276586916a69f3affc6c559.
//
// Solidity: e VoteRevealed(pollID indexed bytes32, secretHash indexed bytes32, choice indexed uint256, voter address, revealer address, votesFor uint256, votesAgainst uint256, votesCommittedButNotRevealed uint256)
func (_TruSetCommitRevealVoting *TruSetCommitRevealVotingFilterer) WatchVoteRevealed(opts *bind.WatchOpts, sink chan<- *TruSetCommitRevealVotingVoteRevealed, pollID [][32]byte, secretHash [][32]byte, choice []*big.Int) (event.Subscription, error) {

	var pollIDRule []interface{}
	for _, pollIDItem := range pollID {
		pollIDRule = append(pollIDRule, pollIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}
	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}

	logs, sub, err := _TruSetCommitRevealVoting.contract.WatchLogs(opts, "VoteRevealed", pollIDRule, secretHashRule, choiceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruSetCommitRevealVotingVoteRevealed)
				if err := _TruSetCommitRevealVoting.contract.UnpackLog(event, "VoteRevealed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
