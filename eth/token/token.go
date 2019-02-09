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

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, _taxDestination common.Address) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, _taxDestination)
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
// Solidity: function balanceOf(_owner address) constant returns(uint256)
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
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
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
// Solidity: function limitOf(_owner address) constant returns(uint256)
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
// Solidity: function limitOf(_owner address) constant returns(uint256)
func (_Token *TokenSession) LimitOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.LimitOf(&_Token.CallOpts, _owner)
}

// LimitOf is a free data retrieval call binding the contract method 0x546a2ca4.
//
// Solidity: function limitOf(_owner address) constant returns(uint256)
func (_Token *TokenCallerSession) LimitOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.LimitOf(&_Token.CallOpts, _owner)
}

// NameOf is a free data retrieval call binding the contract method 0xf5c57382.
//
// Solidity: function nameOf(_owner address) constant returns(string)
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
// Solidity: function nameOf(_owner address) constant returns(string)
func (_Token *TokenSession) NameOf(_owner common.Address) (string, error) {
	return _Token.Contract.NameOf(&_Token.CallOpts, _owner)
}

// NameOf is a free data retrieval call binding the contract method 0xf5c57382.
//
// Solidity: function nameOf(_owner address) constant returns(string)
func (_Token *TokenCallerSession) NameOf(_owner common.Address) (string, error) {
	return _Token.Contract.NameOf(&_Token.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(_owner address) constant returns(uint256)
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
// Solidity: function nonceOf(_owner address) constant returns(uint256)
func (_Token *TokenSession) NonceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.NonceOf(&_Token.CallOpts, _owner)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(_owner address) constant returns(uint256)
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

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_account address, _value uint256) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", _account, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_account address, _value uint256) returns()
func (_Token *TokenSession) Burn(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _account, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_account address, _value uint256) returns()
func (_Token *TokenTransactorSession) Burn(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_account address, _value uint256) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_account address, _value uint256) returns()
func (_Token *TokenSession) Mint(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_account address, _value uint256) returns()
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

// SetLimit is a paid mutator transaction binding the contract method 0x8e6e3e96.
//
// Solidity: function setLimit(_addr address, _limit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactor) SetLimit(opts *bind.TransactOpts, _addr common.Address, _limit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setLimit", _addr, _limit, _r, _s, _v)
}

// SetLimit is a paid mutator transaction binding the contract method 0x8e6e3e96.
//
// Solidity: function setLimit(_addr address, _limit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenSession) SetLimit(_addr common.Address, _limit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetLimit(&_Token.TransactOpts, _addr, _limit, _r, _s, _v)
}

// SetLimit is a paid mutator transaction binding the contract method 0x8e6e3e96.
//
// Solidity: function setLimit(_addr address, _limit uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactorSession) SetLimit(_addr common.Address, _limit *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetLimit(&_Token.TransactOpts, _addr, _limit, _r, _s, _v)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(_addr address, _name string, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactor) SetName(opts *bind.TransactOpts, _addr common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setName", _addr, _name, _r, _s, _v)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(_addr address, _name string, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenSession) SetName(_addr common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _addr, _name, _r, _s, _v)
}

// SetName is a paid mutator transaction binding the contract method 0xe69babe8.
//
// Solidity: function setName(_addr address, _name string, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactorSession) SetName(_addr common.Address, _name string, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _addr, _name, _r, _s, _v)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0x1163c3eb.
//
// Solidity: function setTaxDestination(_taxDestination address) returns()
func (_Token *TokenTransactor) SetTaxDestination(opts *bind.TransactOpts, _taxDestination common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTaxDestination", _taxDestination)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0x1163c3eb.
//
// Solidity: function setTaxDestination(_taxDestination address) returns()
func (_Token *TokenSession) SetTaxDestination(_taxDestination common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTaxDestination(&_Token.TransactOpts, _taxDestination)
}

// SetTaxDestination is a paid mutator transaction binding the contract method 0x1163c3eb.
//
// Solidity: function setTaxDestination(_taxDestination address) returns()
func (_Token *TokenTransactorSession) SetTaxDestination(_taxDestination common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTaxDestination(&_Token.TransactOpts, _taxDestination)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(_from address, _to address, _value uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _from, _to, _value, _r, _s, _v)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(_from address, _to address, _value uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenSession) Transfer(_from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value, _r, _s, _v)
}

// Transfer is a paid mutator transaction binding the contract method 0x1f6b6fa7.
//
// Solidity: function transfer(_from address, _to address, _value uint256, _r bytes32, _s bytes32, _v uint8) returns()
func (_Token *TokenTransactorSession) Transfer(_from common.Address, _to common.Address, _value *big.Int, _r [32]byte, _s [32]byte, _v uint8) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value, _r, _s, _v)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
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
// Solidity: event LimitChanged(addr address, value uint256)
func (_Token *TokenFilterer) FilterLimitChanged(opts *bind.FilterOpts) (*TokenLimitChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return &TokenLimitChangedIterator{contract: _Token.contract, event: "LimitChanged", logs: logs, sub: sub}, nil
}

// WatchLimitChanged is a free log subscription operation binding the contract event 0xef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad3.
//
// Solidity: event LimitChanged(addr address, value uint256)
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
// Solidity: event NameChanged(addr address, value string)
func (_Token *TokenFilterer) FilterNameChanged(opts *bind.FilterOpts) (*TokenNameChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return &TokenNameChangedIterator{contract: _Token.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x3b0a43ccc1ccd1c76ebb1a8d998fdfe1ded3766582dbbbcdda83889170bec53d.
//
// Solidity: event NameChanged(addr address, value string)
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
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: event TaxDestinationChanged(addr address)
func (_Token *TokenFilterer) FilterTaxDestinationChanged(opts *bind.FilterOpts) (*TokenTaxDestinationChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "TaxDestinationChanged")
	if err != nil {
		return nil, err
	}
	return &TokenTaxDestinationChangedIterator{contract: _Token.contract, event: "TaxDestinationChanged", logs: logs, sub: sub}, nil
}

// WatchTaxDestinationChanged is a free log subscription operation binding the contract event 0x09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa.
//
// Solidity: event TaxDestinationChanged(addr address)
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
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*TokenTransferIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, value uint256)
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
