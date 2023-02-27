// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// idopoolprojectInfo is an auto generated low-level Go binding around an user-defined struct.
type idopoolprojectInfo struct {
	TokenA           common.Address
	TokenB           common.Address
	ProjectText      string
	ProjectType      *big.Int
	GroupID          string
	StartTime        *big.Int
	EndTime          *big.Int
	InTokenCapacity  *big.Int
	InTokenAmount    *big.Int
	OutTokenCapacity *big.Int
	MaxExchange      *big.Int
	Exchange         *big.Int
	DecimalA         *big.Int
	DecimalB         *big.Int
	LockNum          *big.Int
	TimeList         []*big.Int
}

// IdfactoryMetaData contains all meta data concerning the Idfactory contract.
var IdfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msg\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"group\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mode\",\"type\":\"uint256\"}],\"name\":\"debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IDOVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPoint0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPoint1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"projectType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"groupID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"timeList\",\"type\":\"uint256[]\"}],\"internalType\":\"structidopool.projectInfo\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"createIDO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setUserTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdfactoryMetaData.ABI instead.
var IdfactoryABI = IdfactoryMetaData.ABI

// Idfactory is an auto generated Go binding around an Ethereum contract.
type Idfactory struct {
	IdfactoryCaller     // Read-only binding to the contract
	IdfactoryTransactor // Write-only binding to the contract
	IdfactoryFilterer   // Log filterer for contract events
}

// IdfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdfactorySession struct {
	Contract     *Idfactory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdfactoryCallerSession struct {
	Contract *IdfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IdfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdfactoryTransactorSession struct {
	Contract     *IdfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IdfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdfactoryRaw struct {
	Contract *Idfactory // Generic contract binding to access the raw methods on
}

// IdfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdfactoryCallerRaw struct {
	Contract *IdfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IdfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdfactoryTransactorRaw struct {
	Contract *IdfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdfactory creates a new instance of Idfactory, bound to a specific deployed contract.
func NewIdfactory(address common.Address, backend bind.ContractBackend) (*Idfactory, error) {
	contract, err := bindIdfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Idfactory{IdfactoryCaller: IdfactoryCaller{contract: contract}, IdfactoryTransactor: IdfactoryTransactor{contract: contract}, IdfactoryFilterer: IdfactoryFilterer{contract: contract}}, nil
}

// NewIdfactoryCaller creates a new read-only instance of Idfactory, bound to a specific deployed contract.
func NewIdfactoryCaller(address common.Address, caller bind.ContractCaller) (*IdfactoryCaller, error) {
	contract, err := bindIdfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdfactoryCaller{contract: contract}, nil
}

// NewIdfactoryTransactor creates a new write-only instance of Idfactory, bound to a specific deployed contract.
func NewIdfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdfactoryTransactor, error) {
	contract, err := bindIdfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdfactoryTransactor{contract: contract}, nil
}

// NewIdfactoryFilterer creates a new log filterer instance of Idfactory, bound to a specific deployed contract.
func NewIdfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdfactoryFilterer, error) {
	contract, err := bindIdfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdfactoryFilterer{contract: contract}, nil
}

// bindIdfactory binds a generic wrapper to an already deployed contract.
func bindIdfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idfactory *IdfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idfactory.Contract.IdfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idfactory *IdfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.Contract.IdfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idfactory *IdfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idfactory.Contract.IdfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idfactory *IdfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idfactory *IdfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idfactory *IdfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idfactory.Contract.contract.Transact(opts, method, params...)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Idfactory *IdfactoryCaller) AdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idfactory.contract.Call(opts, &out, "adminAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Idfactory *IdfactorySession) AdminAddr() (common.Address, error) {
	return _Idfactory.Contract.AdminAddr(&_Idfactory.CallOpts)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Idfactory *IdfactoryCallerSession) AdminAddr() (common.Address, error) {
	return _Idfactory.Contract.AdminAddr(&_Idfactory.CallOpts)
}

// AdminPoint0 is a free data retrieval call binding the contract method 0xaaa12fef.
//
// Solidity: function adminPoint0() view returns(uint256)
func (_Idfactory *IdfactoryCaller) AdminPoint0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idfactory.contract.Call(opts, &out, "adminPoint0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminPoint0 is a free data retrieval call binding the contract method 0xaaa12fef.
//
// Solidity: function adminPoint0() view returns(uint256)
func (_Idfactory *IdfactorySession) AdminPoint0() (*big.Int, error) {
	return _Idfactory.Contract.AdminPoint0(&_Idfactory.CallOpts)
}

// AdminPoint0 is a free data retrieval call binding the contract method 0xaaa12fef.
//
// Solidity: function adminPoint0() view returns(uint256)
func (_Idfactory *IdfactoryCallerSession) AdminPoint0() (*big.Int, error) {
	return _Idfactory.Contract.AdminPoint0(&_Idfactory.CallOpts)
}

// AdminPoint1 is a free data retrieval call binding the contract method 0x31d60d72.
//
// Solidity: function adminPoint1() view returns(uint256)
func (_Idfactory *IdfactoryCaller) AdminPoint1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idfactory.contract.Call(opts, &out, "adminPoint1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminPoint1 is a free data retrieval call binding the contract method 0x31d60d72.
//
// Solidity: function adminPoint1() view returns(uint256)
func (_Idfactory *IdfactorySession) AdminPoint1() (*big.Int, error) {
	return _Idfactory.Contract.AdminPoint1(&_Idfactory.CallOpts)
}

// AdminPoint1 is a free data retrieval call binding the contract method 0x31d60d72.
//
// Solidity: function adminPoint1() view returns(uint256)
func (_Idfactory *IdfactoryCallerSession) AdminPoint1() (*big.Int, error) {
	return _Idfactory.Contract.AdminPoint1(&_Idfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idfactory *IdfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idfactory *IdfactorySession) Owner() (common.Address, error) {
	return _Idfactory.Contract.Owner(&_Idfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idfactory *IdfactoryCallerSession) Owner() (common.Address, error) {
	return _Idfactory.Contract.Owner(&_Idfactory.CallOpts)
}

// ProxyStatus is a free data retrieval call binding the contract method 0x71aeac72.
//
// Solidity: function proxyStatus() view returns(bool)
func (_Idfactory *IdfactoryCaller) ProxyStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Idfactory.contract.Call(opts, &out, "proxyStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProxyStatus is a free data retrieval call binding the contract method 0x71aeac72.
//
// Solidity: function proxyStatus() view returns(bool)
func (_Idfactory *IdfactorySession) ProxyStatus() (bool, error) {
	return _Idfactory.Contract.ProxyStatus(&_Idfactory.CallOpts)
}

// ProxyStatus is a free data retrieval call binding the contract method 0x71aeac72.
//
// Solidity: function proxyStatus() view returns(bool)
func (_Idfactory *IdfactoryCallerSession) ProxyStatus() (bool, error) {
	return _Idfactory.Contract.ProxyStatus(&_Idfactory.CallOpts)
}

// IDOVerify is a paid mutator transaction binding the contract method 0x6f703643.
//
// Solidity: function IDOVerify() returns()
func (_Idfactory *IdfactoryTransactor) IDOVerify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "IDOVerify")
}

// IDOVerify is a paid mutator transaction binding the contract method 0x6f703643.
//
// Solidity: function IDOVerify() returns()
func (_Idfactory *IdfactorySession) IDOVerify() (*types.Transaction, error) {
	return _Idfactory.Contract.IDOVerify(&_Idfactory.TransactOpts)
}

// IDOVerify is a paid mutator transaction binding the contract method 0x6f703643.
//
// Solidity: function IDOVerify() returns()
func (_Idfactory *IdfactoryTransactorSession) IDOVerify() (*types.Transaction, error) {
	return _Idfactory.Contract.IDOVerify(&_Idfactory.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Idfactory *IdfactoryTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Idfactory *IdfactorySession) Close() (*types.Transaction, error) {
	return _Idfactory.Contract.Close(&_Idfactory.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Idfactory *IdfactoryTransactorSession) Close() (*types.Transaction, error) {
	return _Idfactory.Contract.Close(&_Idfactory.TransactOpts)
}

// CreateIDO is a paid mutator transaction binding the contract method 0xa00320be.
//
// Solidity: function createIDO((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) data) returns()
func (_Idfactory *IdfactoryTransactor) CreateIDO(opts *bind.TransactOpts, data idopoolprojectInfo) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "createIDO", data)
}

// CreateIDO is a paid mutator transaction binding the contract method 0xa00320be.
//
// Solidity: function createIDO((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) data) returns()
func (_Idfactory *IdfactorySession) CreateIDO(data idopoolprojectInfo) (*types.Transaction, error) {
	return _Idfactory.Contract.CreateIDO(&_Idfactory.TransactOpts, data)
}

// CreateIDO is a paid mutator transaction binding the contract method 0xa00320be.
//
// Solidity: function createIDO((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) data) returns()
func (_Idfactory *IdfactoryTransactorSession) CreateIDO(data idopoolprojectInfo) (*types.Transaction, error) {
	return _Idfactory.Contract.CreateIDO(&_Idfactory.TransactOpts, data)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Idfactory *IdfactoryTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Idfactory *IdfactorySession) Open() (*types.Transaction, error) {
	return _Idfactory.Contract.Open(&_Idfactory.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Idfactory *IdfactoryTransactorSession) Open() (*types.Transaction, error) {
	return _Idfactory.Contract.Open(&_Idfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idfactory *IdfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idfactory *IdfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Idfactory.Contract.RenounceOwnership(&_Idfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idfactory *IdfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Idfactory.Contract.RenounceOwnership(&_Idfactory.TransactOpts)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x2c801cb1.
//
// Solidity: function setFee0(uint256 fee) returns()
func (_Idfactory *IdfactoryTransactor) SetFee0(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "setFee0", fee)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x2c801cb1.
//
// Solidity: function setFee0(uint256 fee) returns()
func (_Idfactory *IdfactorySession) SetFee0(fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.Contract.SetFee0(&_Idfactory.TransactOpts, fee)
}

// SetFee0 is a paid mutator transaction binding the contract method 0x2c801cb1.
//
// Solidity: function setFee0(uint256 fee) returns()
func (_Idfactory *IdfactoryTransactorSession) SetFee0(fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.Contract.SetFee0(&_Idfactory.TransactOpts, fee)
}

// SetFee1 is a paid mutator transaction binding the contract method 0xa681148d.
//
// Solidity: function setFee1(uint256 fee) returns()
func (_Idfactory *IdfactoryTransactor) SetFee1(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "setFee1", fee)
}

// SetFee1 is a paid mutator transaction binding the contract method 0xa681148d.
//
// Solidity: function setFee1(uint256 fee) returns()
func (_Idfactory *IdfactorySession) SetFee1(fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.Contract.SetFee1(&_Idfactory.TransactOpts, fee)
}

// SetFee1 is a paid mutator transaction binding the contract method 0xa681148d.
//
// Solidity: function setFee1(uint256 fee) returns()
func (_Idfactory *IdfactoryTransactorSession) SetFee1(fee *big.Int) (*types.Transaction, error) {
	return _Idfactory.Contract.SetFee1(&_Idfactory.TransactOpts, fee)
}

// SetUserTime is a paid mutator transaction binding the contract method 0x1d8473ab.
//
// Solidity: function setUserTime(address addr) returns()
func (_Idfactory *IdfactoryTransactor) SetUserTime(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "setUserTime", addr)
}

// SetUserTime is a paid mutator transaction binding the contract method 0x1d8473ab.
//
// Solidity: function setUserTime(address addr) returns()
func (_Idfactory *IdfactorySession) SetUserTime(addr common.Address) (*types.Transaction, error) {
	return _Idfactory.Contract.SetUserTime(&_Idfactory.TransactOpts, addr)
}

// SetUserTime is a paid mutator transaction binding the contract method 0x1d8473ab.
//
// Solidity: function setUserTime(address addr) returns()
func (_Idfactory *IdfactoryTransactorSession) SetUserTime(addr common.Address) (*types.Transaction, error) {
	return _Idfactory.Contract.SetUserTime(&_Idfactory.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idfactory *IdfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Idfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idfactory *IdfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Idfactory.Contract.TransferOwnership(&_Idfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idfactory *IdfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Idfactory.Contract.TransferOwnership(&_Idfactory.TransactOpts, newOwner)
}

// IdfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Idfactory contract.
type IdfactoryOwnershipTransferredIterator struct {
	Event *IdfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdfactoryOwnershipTransferred)
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
		it.Event = new(IdfactoryOwnershipTransferred)
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
func (it *IdfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Idfactory contract.
type IdfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Idfactory *IdfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Idfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdfactoryOwnershipTransferredIterator{contract: _Idfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Idfactory *IdfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Idfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdfactoryOwnershipTransferred)
				if err := _Idfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Idfactory *IdfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*IdfactoryOwnershipTransferred, error) {
	event := new(IdfactoryOwnershipTransferred)
	if err := _Idfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdfactoryDebugIterator is returned from FilterDebug and is used to iterate over the raw logs and unpacked data for Debug events raised by the Idfactory contract.
type IdfactoryDebugIterator struct {
	Event *IdfactoryDebug // Event containing the contract specifics and raw log

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
func (it *IdfactoryDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdfactoryDebug)
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
		it.Event = new(IdfactoryDebug)
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
func (it *IdfactoryDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdfactoryDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdfactoryDebug represents a Debug event raised by the Idfactory contract.
type IdfactoryDebug struct {
	Name  string
	Msg   common.Address
	Owner common.Address
	Admin common.Address
	Text  string
	Group string
	Mode  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDebug is a free log retrieval operation binding the contract event 0xe8e7ed9aca3d2eb136e604edb16874d315c68555c75fd0f87e736ccbfa32f0ec.
//
// Solidity: event debug(string name, address msg, address owner, address admin, string text, string group, uint256 mode)
func (_Idfactory *IdfactoryFilterer) FilterDebug(opts *bind.FilterOpts) (*IdfactoryDebugIterator, error) {

	logs, sub, err := _Idfactory.contract.FilterLogs(opts, "debug")
	if err != nil {
		return nil, err
	}
	return &IdfactoryDebugIterator{contract: _Idfactory.contract, event: "debug", logs: logs, sub: sub}, nil
}

// WatchDebug is a free log subscription operation binding the contract event 0xe8e7ed9aca3d2eb136e604edb16874d315c68555c75fd0f87e736ccbfa32f0ec.
//
// Solidity: event debug(string name, address msg, address owner, address admin, string text, string group, uint256 mode)
func (_Idfactory *IdfactoryFilterer) WatchDebug(opts *bind.WatchOpts, sink chan<- *IdfactoryDebug) (event.Subscription, error) {

	logs, sub, err := _Idfactory.contract.WatchLogs(opts, "debug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdfactoryDebug)
				if err := _Idfactory.contract.UnpackLog(event, "debug", log); err != nil {
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

// ParseDebug is a log parse operation binding the contract event 0xe8e7ed9aca3d2eb136e604edb16874d315c68555c75fd0f87e736ccbfa32f0ec.
//
// Solidity: event debug(string name, address msg, address owner, address admin, string text, string group, uint256 mode)
func (_Idfactory *IdfactoryFilterer) ParseDebug(log types.Log) (*IdfactoryDebug, error) {
	event := new(IdfactoryDebug)
	if err := _Idfactory.contract.UnpackLog(event, "debug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
