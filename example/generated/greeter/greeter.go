// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package greeter

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GreeterABI is the input ABI used to generate the binding from.
const GreeterABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_greeting\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"lalala\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mortal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GreeterFuncSigs maps the 4-byte function signature to its string representation.
var GreeterFuncSigs = map[string]string{
	"cfae3217": "greet()",
	"41c0e1b5": "kill()",
	"f1eae25c": "mortal()",
}

// GreeterBin is the compiled bytecode used for deploying new contracts.
var GreeterBin = "0x608060405234801561001057600080fd5b506040516103843803806103848339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b5060405260209081015184519093506100fb9250600191850190610103565b50505061019e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014457805160ff1916838001178555610171565b82800160010185558215610171579182015b82811115610171578251825591602001919060010190610156565b5061017d929150610181565b5090565b61019b91905b8082111561017d5760008155600101610187565b90565b6101d7806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806341c0e1b514610046578063cfae321714610050578063f1eae25c146100cd575b600080fd5b61004e6100d5565b005b6100586100f8565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009257818101518382015260200161007a565b50505050905090810190601f1680156100bf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61004e61018d565b6000546001600160a01b03163314156100f6576000546001600160a01b0316ff5b565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156101835780601f1061015857610100808354040283529160200191610183565b820191906000526020600020905b81548152906001019060200180831161016657829003601f168201915b5050505050905090565b600080546001600160a01b0319163317905556fea26469706673582212204327b18e6d6d17c8d204778f2f32b393628fe49044f6eb3e612102718ec7b2d964736f6c63430006020033"

// DeployGreeter deploys a new Ethereum contract, binding an instance of Greeter to it.
func DeployGreeter(auth *bind.TransactOpts, backend bind.ContractBackend, _greeting string, lalala uint8) (common.Address, *types.Transaction, *Greeter, error) {
	parsed, err := abi.JSON(strings.NewReader(GreeterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GreeterBin), backend, _greeting, lalala)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Greeter{GreeterCaller: GreeterCaller{contract: contract}, GreeterTransactor: GreeterTransactor{contract: contract}, GreeterFilterer: GreeterFilterer{contract: contract}}, nil
}

// Greeter is an auto generated Go binding around an Ethereum contract.
type Greeter struct {
	GreeterCaller     // Read-only binding to the contract
	GreeterTransactor // Write-only binding to the contract
	GreeterFilterer   // Log filterer for contract events
}

// GreeterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GreeterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GreeterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GreeterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreeterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GreeterSession struct {
	Contract     *Greeter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GreeterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GreeterCallerSession struct {
	Contract *GreeterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GreeterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GreeterTransactorSession struct {
	Contract     *GreeterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GreeterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GreeterRaw struct {
	Contract *Greeter // Generic contract binding to access the raw methods on
}

// GreeterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GreeterCallerRaw struct {
	Contract *GreeterCaller // Generic read-only contract binding to access the raw methods on
}

// GreeterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GreeterTransactorRaw struct {
	Contract *GreeterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGreeter creates a new instance of Greeter, bound to a specific deployed contract.
func NewGreeter(address common.Address, backend bind.ContractBackend) (*Greeter, error) {
	contract, err := bindGreeter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Greeter{GreeterCaller: GreeterCaller{contract: contract}, GreeterTransactor: GreeterTransactor{contract: contract}, GreeterFilterer: GreeterFilterer{contract: contract}}, nil
}

// NewGreeterCaller creates a new read-only instance of Greeter, bound to a specific deployed contract.
func NewGreeterCaller(address common.Address, caller bind.ContractCaller) (*GreeterCaller, error) {
	contract, err := bindGreeter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GreeterCaller{contract: contract}, nil
}

// NewGreeterTransactor creates a new write-only instance of Greeter, bound to a specific deployed contract.
func NewGreeterTransactor(address common.Address, transactor bind.ContractTransactor) (*GreeterTransactor, error) {
	contract, err := bindGreeter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GreeterTransactor{contract: contract}, nil
}

// NewGreeterFilterer creates a new log filterer instance of Greeter, bound to a specific deployed contract.
func NewGreeterFilterer(address common.Address, filterer bind.ContractFilterer) (*GreeterFilterer, error) {
	contract, err := bindGreeter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GreeterFilterer{contract: contract}, nil
}

// bindGreeter binds a generic wrapper to an already deployed contract.
func bindGreeter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GreeterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greeter *GreeterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Greeter.Contract.GreeterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greeter *GreeterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.Contract.GreeterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greeter *GreeterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greeter.Contract.GreeterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greeter *GreeterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Greeter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greeter *GreeterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greeter *GreeterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greeter.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_Greeter *GreeterCaller) Greet(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Greeter.contract.Call(opts, out, "greet")
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_Greeter *GreeterSession) Greet() (string, error) {
	return _Greeter.Contract.Greet(&_Greeter.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_Greeter *GreeterCallerSession) Greet() (string, error) {
	return _Greeter.Contract.Greet(&_Greeter.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greeter *GreeterTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greeter *GreeterSession) Kill() (*types.Transaction, error) {
	return _Greeter.Contract.Kill(&_Greeter.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greeter *GreeterTransactorSession) Kill() (*types.Transaction, error) {
	return _Greeter.Contract.Kill(&_Greeter.TransactOpts)
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_Greeter *GreeterTransactor) Mortal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greeter.contract.Transact(opts, "mortal")
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_Greeter *GreeterSession) Mortal() (*types.Transaction, error) {
	return _Greeter.Contract.Mortal(&_Greeter.TransactOpts)
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_Greeter *GreeterTransactorSession) Mortal() (*types.Transaction, error) {
	return _Greeter.Contract.Mortal(&_Greeter.TransactOpts)
}
