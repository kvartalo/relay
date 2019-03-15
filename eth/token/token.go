// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// OffchainsigABI is the input ABI used to generate the binding from.
const OffchainsigABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OffchainsigBin is the compiled bytecode used for deploying new contracts.
const OffchainsigBin = `0x608060405234801561001057600080fd5b5060df8061001f6000396000f3fe6080604052348015600f57600080fd5b50600436106044577c01000000000000000000000000000000000000000000000000000000006000350463ed2a2d6481146049575b600080fd5b607960048036036020811015605d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16608b565b60408051918252519081900360200190f35b73ffffffffffffffffffffffffffffffffffffffff166000908152602081905260409020549056fea165627a7a723058202ec18b67ca88d6d027c9ba30a29705c49529ceba0ec2149806b0c30bc133db880029`

// DeployOffchainsig deploys a new Ethereum contract, binding an instance of Offchainsig to it.
func DeployOffchainsig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Offchainsig, error) {
	parsed, err := abi.JSON(strings.NewReader(OffchainsigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OffchainsigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Offchainsig{OffchainsigCaller: OffchainsigCaller{contract: contract}, OffchainsigTransactor: OffchainsigTransactor{contract: contract}, OffchainsigFilterer: OffchainsigFilterer{contract: contract}}, nil
}

// Offchainsig is an auto generated Go binding around an Ethereum contract.
type Offchainsig struct {
	OffchainsigCaller     // Read-only binding to the contract
	OffchainsigTransactor // Write-only binding to the contract
	OffchainsigFilterer   // Log filterer for contract events
}

// OffchainsigCaller is an auto generated read-only Go binding around an Ethereum contract.
type OffchainsigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainsigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OffchainsigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainsigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OffchainsigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainsigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OffchainsigSession struct {
	Contract     *Offchainsig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OffchainsigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OffchainsigCallerSession struct {
	Contract *OffchainsigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OffchainsigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OffchainsigTransactorSession struct {
	Contract     *OffchainsigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OffchainsigRaw is an auto generated low-level Go binding around an Ethereum contract.
type OffchainsigRaw struct {
	Contract *Offchainsig // Generic contract binding to access the raw methods on
}

// OffchainsigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OffchainsigCallerRaw struct {
	Contract *OffchainsigCaller // Generic read-only contract binding to access the raw methods on
}

// OffchainsigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OffchainsigTransactorRaw struct {
	Contract *OffchainsigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOffchainsig creates a new instance of Offchainsig, bound to a specific deployed contract.
func NewOffchainsig(address common.Address, backend bind.ContractBackend) (*Offchainsig, error) {
	contract, err := bindOffchainsig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Offchainsig{OffchainsigCaller: OffchainsigCaller{contract: contract}, OffchainsigTransactor: OffchainsigTransactor{contract: contract}, OffchainsigFilterer: OffchainsigFilterer{contract: contract}}, nil
}

// NewOffchainsigCaller creates a new read-only instance of Offchainsig, bound to a specific deployed contract.
func NewOffchainsigCaller(address common.Address, caller bind.ContractCaller) (*OffchainsigCaller, error) {
	contract, err := bindOffchainsig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainsigCaller{contract: contract}, nil
}

// NewOffchainsigTransactor creates a new write-only instance of Offchainsig, bound to a specific deployed contract.
func NewOffchainsigTransactor(address common.Address, transactor bind.ContractTransactor) (*OffchainsigTransactor, error) {
	contract, err := bindOffchainsig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainsigTransactor{contract: contract}, nil
}

// NewOffchainsigFilterer creates a new log filterer instance of Offchainsig, bound to a specific deployed contract.
func NewOffchainsigFilterer(address common.Address, filterer bind.ContractFilterer) (*OffchainsigFilterer, error) {
	contract, err := bindOffchainsig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffchainsigFilterer{contract: contract}, nil
}

// bindOffchainsig binds a generic wrapper to an already deployed contract.
func bindOffchainsig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffchainsigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Offchainsig *OffchainsigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Offchainsig.Contract.OffchainsigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Offchainsig *OffchainsigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Offchainsig.Contract.OffchainsigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Offchainsig *OffchainsigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Offchainsig.Contract.OffchainsigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Offchainsig *OffchainsigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Offchainsig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Offchainsig *OffchainsigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Offchainsig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Offchainsig *OffchainsigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Offchainsig.Contract.contract.Transact(opts, method, params...)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Offchainsig *OffchainsigCaller) NonceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Offchainsig.contract.Call(opts, out, "nonceOf", _owner)
	return *ret0, err
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Offchainsig *OffchainsigSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _Offchainsig.Contract.NonceOf(&_Offchainsig.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Offchainsig *OffchainsigCallerSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _Offchainsig.Contract.NonceOf(&_Offchainsig.CallOpts, _owner)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058200b941bfffd9358f34bd01240810b8ce720f774ded6328e073457cd6173c164930029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxDestination\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultBalanceLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"limitOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_taxDestination\",\"type\":\"address\"},{\"name\":\"_taxPercent\",\"type\":\"uint256\"}],\"name\":\"setTaxDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nameOf\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isTax\",\"type\":\"bool\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"TaxDestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x60806040819052600060028190556000196003558054600160a060020a0319163317808255600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3610f8a806100626000396000f3fe608060405234801561001057600080fd5b5060043610610133576000357c010000000000000000000000000000000000000000000000000000000090048063715018a6116100bf578063d0c3fe0b1161008e578063d0c3fe0b146102c1578063e69babe8146102ed578063ed2a2d641461037c578063f2fde38b146103a2578063f5c57382146103c857610133565b8063715018a61461028d5780637541f41c146102955780638da5cb5b1461029d5780638f32d59b146102a557610133565b806340c10f191161010657806340c10f19146101f057806342966c681461021c57806347cf5f0414610239578063546a2ca41461024157806370a082311461026757610133565b806318160ddd146101385780631f6b6fa7146101525780632c547b3d146101a057806336db43b5146101c4575b600080fd5b610140610463565b60408051918252519081900360200190f35b61019e600480360360c081101561016857600080fd5b508035600160a060020a039081169160208101359091169060408101359060608101359060808101359060a0013560ff16610469565b005b6101a86104d0565b60408051600160a060020a039092168252519081900360200190f35b61019e600480360360408110156101da57600080fd5b50600160a060020a0381351690602001356104df565b61019e6004803603604081101561020657600080fd5b50600160a060020a03813516906020013561054e565b61019e6004803603602081101561023257600080fd5b5035610654565b61014061072a565b6101406004803603602081101561025757600080fd5b5035600160a060020a0316610730565b6101406004803603602081101561027d57600080fd5b5035600160a060020a031661074e565b61019e610769565b6101406107d3565b6101a86107d9565b6102ad6107e9565b604080519115158252519081900360200190f35b61019e600480360360408110156102d757600080fd5b50600160a060020a0381351690602001356107fa565b61019e600480360360a081101561030357600080fd5b600160a060020a03823516919081019060408101602082013564010000000081111561032e57600080fd5b82018360208201111561034057600080fd5b8035906020019184600183028401116401000000008311171561036257600080fd5b91935091508035906020810135906040013560ff166108e6565b6101406004803603602081101561039257600080fd5b5035600160a060020a0316610973565b61019e600480360360208110156103b857600080fd5b5035600160a060020a031661098e565b6103ee600480360360208110156103de57600080fd5b5035600160a060020a03166109ad565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610428578181015183820152602001610410565b50505050905090810190601f1680156104555780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60055481565b604080516c01000000000000000000000000600160a060020a03808a1682026020840152881602603482015260488082018790528251808303909101815260689091019091526104bd908790858585610a5e565b6104c8868686610c50565b505050505050565b600454600160a060020a031681565b6104e76107e9565b15156104f257600080fd5b600160a060020a0382166000818152600660209081526040918290206001018490558151928352820183905280517fef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad39281900390910190a15050565b6105566107e9565b151561056157600080fd5b600160a060020a038216151561057657600080fd5b600160a060020a03821660009081526006602052604090205461059f908263ffffffff610d3c16565b600160a060020a0383166000908152600660205260409020556005546105cb908263ffffffff610d3c16565b600555604080516000808252600160a060020a0385166020830152818301849052606082015290517ff431703d7399230d223d254d7e8479fcd3149952c18616d3c9eb716c6ef5d6729181900360800190a16040805182815290517f07883703ed0e86588a40d76551c92f8a4b329e3bf19765e0e6749473c1a846659181900360200190a15050565b61065c6107e9565b151561066757600080fd5b33600090815260066020526040902054610687908263ffffffff610d5516565b336000908152600660205260409020556005546106aa908263ffffffff610d5516565b60055560408051338152600060208201819052818301849052606082015290517ff431703d7399230d223d254d7e8479fcd3149952c18616d3c9eb716c6ef5d6729181900360800190a16040805182815290517fb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb9181900360200190a150565b60035481565b600160a060020a031660009081526006602052604090206001015490565b600160a060020a031660009081526006602052604090205490565b6107716107e9565b151561077c57600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60025481565b600054600160a060020a03165b90565b600054600160a060020a0316331490565b6108026107e9565b151561080d57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03848116919091178083556002849055811660009081526006602090815260408083206000196001918201559454909316808352918390209093015482519182529281019290925280517fef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad39281900390910190a160045460408051600160a060020a039092168252517f09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa9181900360200190a15050565b610944868787876040516020018084600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401838380828437808301925050509350505050604051602081830303815290604052858585610a5e565b600160a060020a038616600090815260066020526040902061096a906002018686610ec6565b50505050505050565b600160a060020a031660009081526001602052604090205490565b6109966107e9565b15156109a157600080fd5b6109aa81610d6a565b50565b600160a060020a0381166000908152600660209081526040918290206002908101805484516001821615610100026000190190911692909204601f81018490048402830184019094528382526060939192909190830182828015610a525780601f10610a2757610100808354040283529160200191610a52565b820191906000526020600020905b815481529060010190602001808311610a3557829003601f168201915b50505050509050919050565b600160a060020a03851660009081526001602090815260408083205490517f190000000000000000000000000000000000000000000000000000000000000081840181815260218301869052306c0100000000000000000000000081026022850152603684018590528a5192958795919491938c9392605601918401908083835b60208310610afe5780518252601f199092019160209182019101610adf565b6001836020036101000a03801982511681845116808217855250505050505090500195505050505050604051602081830303815290604052805190602001209050600060018284878760405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610b9e573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a0380821690881614610c2557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f73656e6465722d616464726573732d646f65732d6e6f742d6d61746368000000604482015290519081900360640190fd5b505050600160a060020a03909316600090815260016020819052604090912080549091019055505050565b610c5d8383836000610de7565b600254600160a060020a03841660009081526006602052604090205460649183029190910490811115610ca55750600160a060020a0383166000908152600660205260409020545b6000811115610cc857600454610cc8908590600160a060020a0316836001610de7565b600160a060020a0383166000908152600660205260409020600101541515610d0657600354610cf68461074e565b1115610d0157600080fd5b610d36565b600160a060020a038316600090815260066020526040902060010154610d2b8461074e565b1115610d3657600080fd5b50505050565b600082820183811015610d4e57600080fd5b9392505050565b600082821115610d6457600080fd5b50900390565b600160a060020a0381161515610d7f57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a0383161515610dfc57600080fd5b600160a060020a038416600090815260066020526040902054610e25908363ffffffff610d5516565b600160a060020a038086166000908152600660205260408082209390935590851681522054610e5a908363ffffffff610d3c16565b600160a060020a03808516600081815260066020908152604091829020949094558051928816835292820152808201849052821515606082015290517ff431703d7399230d223d254d7e8479fcd3149952c18616d3c9eb716c6ef5d6729181900360800190a150505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f075782800160ff19823516178555610f34565b82800160010185558215610f34579182015b82811115610f34578235825591602001919060010190610f19565b50610f40929150610f44565b5090565b6107e691905b80821115610f405760008155600101610f4a56fea165627a7a7230582005988379de664f9345d467661de50d8c0b8c7f6d5934227251b5b6f1310992cd0029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// DefaultBalanceLimit is a free data retrieval call binding the contract method 0x47cf5f04.
//
// Solidity: function defaultBalanceLimit() constant returns(uint256)
func (_Token *TokenCaller) DefaultBalanceLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "defaultBalanceLimit")
	return *ret0, err
}

// DefaultBalanceLimit is a free data retrieval call binding the contract method 0x47cf5f04.
//
// Solidity: function defaultBalanceLimit() constant returns(uint256)
func (_Token *TokenSession) DefaultBalanceLimit() (*big.Int, error) {
	return _Token.Contract.DefaultBalanceLimit(&_Token.CallOpts)
}

// DefaultBalanceLimit is a free data retrieval call binding the contract method 0x47cf5f04.
//
// Solidity: function defaultBalanceLimit() constant returns(uint256)
func (_Token *TokenCallerSession) DefaultBalanceLimit() (*big.Int, error) {
	return _Token.Contract.DefaultBalanceLimit(&_Token.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenSession) IsOwner() (bool, error) {
	return _Token.Contract.IsOwner(&_Token.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenCallerSession) IsOwner() (bool, error) {
	return _Token.Contract.IsOwner(&_Token.CallOpts)
}

// LimitOf is a free data retrieval call binding the contract method 0x546a2ca4.
//
// Solidity: function limitOf(address _owner) constant returns(uint256)
func (_Token *TokenCaller) LimitOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "limitOf", _owner)
	return *ret0, err
}

// LimitOf is a free data retrieval call binding the contract method 0x546a2ca4.
//
// Solidity: function limitOf(address _owner) constant returns(uint256)
func (_Token *TokenSession) LimitOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.LimitOf(&_Token.CallOpts, _owner)
}

// LimitOf is a free data retrieval call binding the contract method 0x546a2ca4.
//
// Solidity: function limitOf(address _owner) constant returns(uint256)
func (_Token *TokenCallerSession) LimitOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.LimitOf(&_Token.CallOpts, _owner)
}

// NameOf is a free data retrieval call binding the contract method 0xf5c57382.
//
// Solidity: function nameOf(address _owner) constant returns(string)
func (_Token *TokenCaller) NameOf(opts *bind.CallOpts, _owner common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "nameOf", _owner)
	return *ret0, err
}

// NameOf is a free data retrieval call binding the contract method 0xf5c57382.
//
// Solidity: function nameOf(address _owner) constant returns(string)
func (_Token *TokenSession) NameOf(_owner common.Address) (string, error) {
	return _Token.Contract.NameOf(&_Token.CallOpts, _owner)
}

// NameOf is a free data retrieval call binding the contract method 0xf5c57382.
//
// Solidity: function nameOf(address _owner) constant returns(string)
func (_Token *TokenCallerSession) NameOf(_owner common.Address) (string, error) {
	return _Token.Contract.NameOf(&_Token.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Token *TokenCaller) NonceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "nonceOf", _owner)
	return *ret0, err
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Token *TokenSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.NonceOf(&_Token.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(address _owner) constant returns(uint256)
func (_Token *TokenCallerSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.NonceOf(&_Token.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// TaxDestination is a free data retrieval call binding the contract method 0x2c547b3d.
//
// Solidity: function taxDestination() constant returns(address)
func (_Token *TokenCaller) TaxDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "taxDestination")
	return *ret0, err
}

// TaxDestination is a free data retrieval call binding the contract method 0x2c547b3d.
//
// Solidity: function taxDestination() constant returns(address)
func (_Token *TokenSession) TaxDestination() (common.Address, error) {
	return _Token.Contract.TaxDestination(&_Token.CallOpts)
}

// TaxDestination is a free data retrieval call binding the contract method 0x2c547b3d.
//
// Solidity: function taxDestination() constant returns(address)
func (_Token *TokenCallerSession) TaxDestination() (common.Address, error) {
	return _Token.Contract.TaxDestination(&_Token.CallOpts)
}

// TaxPercent is a free data retrieval call binding the contract method 0x7541f41c.
//
// Solidity: function taxPercent() constant returns(uint256)
func (_Token *TokenCaller) TaxPercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "taxPercent")
	return *ret0, err
}

// TaxPercent is a free data retrieval call binding the contract method 0x7541f41c.
//
// Solidity: function taxPercent() constant returns(uint256)
func (_Token *TokenSession) TaxPercent() (*big.Int, error) {
	return _Token.Contract.TaxPercent(&_Token.CallOpts)
}

// TaxPercent is a free data retrieval call binding the contract method 0x7541f41c.
//
// Solidity: function taxPercent() constant returns(uint256)
func (_Token *TokenCallerSession) TaxPercent() (*big.Int, error) {
	return _Token.Contract.TaxPercent(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Token *TokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Token *TokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_Token *TokenSession) Mint(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_Token *TokenTransactorSession) Mint(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _account, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address _addr, uint256 _limit) returns()
func (_Token *TokenTransactor) SetLimit(opts *bind.TransactOpts, _addr common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setLimit", _addr, _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address _addr, uint256 _limit) returns()
func (_Token *TokenSession) SetLimit(_addr common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetLimit(&_Token.TransactOpts, _addr, _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address _addr, uint256 _limit) returns()
func (_Token *TokenTransactorSession) SetLimit(_addr common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetLimit(&_Token.TransactOpts, _addr, _limit)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(address _from, string _name, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenTransactor) SetName(opts *bind.TransactOpts, _from common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setName", _from, _name, _r, _s, _v)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(address _from, string _name, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenSession) SetName(_from common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _from, _name, _r, _s, _v)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(address _from, string _name, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenTransactorSession) SetName(_from common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _from, _name, _r, _s, _v)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0xd0c3fe0b.
//
// Solidity: function setTaxDestination(address _taxDestination, uint256 _taxPercent) returns()
func (_Token *TokenTransactor) SetTaxDestination(opts *bind.TransactOpts, _taxDestination common.Address, _taxPercent *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTaxDestination", _taxDestination, _taxPercent)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0xd0c3fe0b.
//
// Solidity: function setTaxDestination(address _taxDestination, uint256 _taxPercent) returns()
func (_Token *TokenSession) SetTaxDestination(_taxDestination common.Address, _taxPercent *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTaxDestination(&_Token.TransactOpts, _taxDestination, _taxPercent)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0xd0c3fe0b.
//
// Solidity: function setTaxDestination(address _taxDestination, uint256 _taxPercent) returns()
func (_Token *TokenTransactorSession) SetTaxDestination(_taxDestination common.Address, _taxPercent *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTaxDestination(&_Token.TransactOpts, _taxDestination, _taxPercent)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(address _from, address _to, uint256 _value, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _from, _to, _value, _r, _s, _v)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(address _from, address _to, uint256 _value, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenSession) Transfer(_from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value, _r, _s, _v)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(address _from, address _to, uint256 _value, bytes32 _r, bytes32 _s, uint8 _v) returns()
func (_Token *TokenTransactorSession) Transfer(_from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value, _r, _s, _v)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Token contract.
type TokenBurnIterator struct {
	Event *TokenBurn // Event containing the contract specifics and raw log

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
func (it *TokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurn)
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
		it.Event = new(TokenBurn)
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
func (it *TokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurn represents a Burn event raised by the Token contract.
type TokenBurn struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 value)
func (_Token *TokenFilterer) FilterBurn(opts *bind.FilterOpts) (*TokenBurnIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &TokenBurnIterator{contract: _Token.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 value)
func (_Token *TokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *TokenBurn) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurn)
				if err := _Token.contract.UnpackLog(event, "Burn", log); err != nil {
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

// TokenLimitChangedIterator is returned from FilterLimitChanged and is used to iterate over the raw logs and unpacked data for LimitChanged events raised by the Token contract.
type TokenLimitChangedIterator struct {
	Event *TokenLimitChanged // Event containing the contract specifics and raw log

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
func (it *TokenLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLimitChanged)
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
		it.Event = new(TokenLimitChanged)
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
func (it *TokenLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLimitChanged represents a LimitChanged event raised by the Token contract.
type TokenLimitChanged struct {
	Addr  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLimitChanged is a free log retrieval operation binding the contract event 0xef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad3.
//
// Solidity: event LimitChanged(address addr, uint256 value)
func (_Token *TokenFilterer) FilterLimitChanged(opts *bind.FilterOpts) (*TokenLimitChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return &TokenLimitChangedIterator{contract: _Token.contract, event: "LimitChanged", logs: logs, sub: sub}, nil
}

// WatchLimitChanged is a free log subscription operation binding the contract event 0xef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad3.
//
// Solidity: event LimitChanged(address addr, uint256 value)
func (_Token *TokenFilterer) WatchLimitChanged(opts *bind.WatchOpts, sink chan<- *TokenLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLimitChanged)
				if err := _Token.contract.UnpackLog(event, "LimitChanged", log); err != nil {
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

// TokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Token contract.
type TokenMintIterator struct {
	Event *TokenMint // Event containing the contract specifics and raw log

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
func (it *TokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMint)
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
		it.Event = new(TokenMint)
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
func (it *TokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMint represents a Mint event raised by the Token contract.
type TokenMint struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x07883703ed0e86588a40d76551c92f8a4b329e3bf19765e0e6749473c1a84665.
//
// Solidity: event Mint(uint256 value)
func (_Token *TokenFilterer) FilterMint(opts *bind.FilterOpts) (*TokenMintIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &TokenMintIterator{contract: _Token.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x07883703ed0e86588a40d76551c92f8a4b329e3bf19765e0e6749473c1a84665.
//
// Solidity: event Mint(uint256 value)
func (_Token *TokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TokenMint) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMint)
				if err := _Token.contract.UnpackLog(event, "Mint", log); err != nil {
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

// TokenNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Token contract.
type TokenNameChangedIterator struct {
	Event *TokenNameChanged // Event containing the contract specifics and raw log

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
func (it *TokenNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNameChanged)
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
		it.Event = new(TokenNameChanged)
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
func (it *TokenNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNameChanged represents a NameChanged event raised by the Token contract.
type TokenNameChanged struct {
	Addr  common.Address
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0x3b0a43ccc1ccd1c76ebb1a8d998fdfe1ded3766582dbbbcdda83889170bec53d.
//
// Solidity: event NameChanged(address addr, string value)
func (_Token *TokenFilterer) FilterNameChanged(opts *bind.FilterOpts) (*TokenNameChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return &TokenNameChangedIterator{contract: _Token.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x3b0a43ccc1ccd1c76ebb1a8d998fdfe1ded3766582dbbbcdda83889170bec53d.
//
// Solidity: event NameChanged(address addr, string value)
func (_Token *TokenFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *TokenNameChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNameChanged)
				if err := _Token.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenTaxDestinationChangedIterator is returned from FilterTaxDestinationChanged and is used to iterate over the raw logs and unpacked data for TaxDestinationChanged events raised by the Token contract.
type TokenTaxDestinationChangedIterator struct {
	Event *TokenTaxDestinationChanged // Event containing the contract specifics and raw log

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
func (it *TokenTaxDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTaxDestinationChanged)
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
		it.Event = new(TokenTaxDestinationChanged)
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
func (it *TokenTaxDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTaxDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTaxDestinationChanged represents a TaxDestinationChanged event raised by the Token contract.
type TokenTaxDestinationChanged struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTaxDestinationChanged is a free log retrieval operation binding the contract event 0x09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa.
//
// Solidity: event TaxDestinationChanged(address addr)
func (_Token *TokenFilterer) FilterTaxDestinationChanged(opts *bind.FilterOpts) (*TokenTaxDestinationChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "TaxDestinationChanged")
	if err != nil {
		return nil, err
	}
	return &TokenTaxDestinationChangedIterator{contract: _Token.contract, event: "TaxDestinationChanged", logs: logs, sub: sub}, nil
}

// WatchTaxDestinationChanged is a free log subscription operation binding the contract event 0x09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa.
//
// Solidity: event TaxDestinationChanged(address addr)
func (_Token *TokenFilterer) WatchTaxDestinationChanged(opts *bind.WatchOpts, sink chan<- *TokenTaxDestinationChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "TaxDestinationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTaxDestinationChanged)
				if err := _Token.contract.UnpackLog(event, "TaxDestinationChanged", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	IsTax bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xf431703d7399230d223d254d7e8479fcd3149952c18616d3c9eb716c6ef5d672.
//
// Solidity: event Transfer(address from, address to, uint256 value, bool isTax)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*TokenTransferIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xf431703d7399230d223d254d7e8479fcd3149952c18616d3c9eb716c6ef5d672.
//
// Solidity: event Transfer(address from, address to, uint256 value, bool isTax)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
