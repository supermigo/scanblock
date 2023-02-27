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

// idopoollockUser is an auto generated low-level Go binding around an user-defined struct.
type idopoollockUser struct {
	StartTime *big.Int
	AmountB   *big.Int
}

// idopoolprojectInfo is an auto generated low-level Go binding around an user-defined struct.

// IdoMetaData contains all meta data concerning the Ido contract.
var IdoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"projectType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"groupID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"timeList\",\"type\":\"uint256[]\"}],\"internalType\":\"structidopool.projectInfo\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_adminPoint\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msg\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exec\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"advanceRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryProjectInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"projectType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"groupID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outTokenCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimalB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"timeList\",\"type\":\"uint256[]\"}],\"internalType\":\"structidopool.projectInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"internalType\":\"structidopool.lockUser[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendAddrA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendAddrB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IdoABI is the input ABI used to generate the binding from.
// Deprecated: Use IdoMetaData.ABI instead.
var IdoABI = IdoMetaData.ABI

// Ido is an auto generated Go binding around an Ethereum contract.
type Ido struct {
	IdoCaller     // Read-only binding to the contract
	IdoTransactor // Write-only binding to the contract
	IdoFilterer   // Log filterer for contract events
}

// IdoCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdoSession struct {
	Contract     *Ido              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdoCallerSession struct {
	Contract *IdoCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IdoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdoTransactorSession struct {
	Contract     *IdoTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdoRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdoRaw struct {
	Contract *Ido // Generic contract binding to access the raw methods on
}

// IdoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdoCallerRaw struct {
	Contract *IdoCaller // Generic read-only contract binding to access the raw methods on
}

// IdoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdoTransactorRaw struct {
	Contract *IdoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdo creates a new instance of Ido, bound to a specific deployed contract.
func NewIdo(address common.Address, backend bind.ContractBackend) (*Ido, error) {
	contract, err := bindIdo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ido{IdoCaller: IdoCaller{contract: contract}, IdoTransactor: IdoTransactor{contract: contract}, IdoFilterer: IdoFilterer{contract: contract}}, nil
}

// NewIdoCaller creates a new read-only instance of Ido, bound to a specific deployed contract.
func NewIdoCaller(address common.Address, caller bind.ContractCaller) (*IdoCaller, error) {
	contract, err := bindIdo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdoCaller{contract: contract}, nil
}

// NewIdoTransactor creates a new write-only instance of Ido, bound to a specific deployed contract.
func NewIdoTransactor(address common.Address, transactor bind.ContractTransactor) (*IdoTransactor, error) {
	contract, err := bindIdo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdoTransactor{contract: contract}, nil
}

// NewIdoFilterer creates a new log filterer instance of Ido, bound to a specific deployed contract.
func NewIdoFilterer(address common.Address, filterer bind.ContractFilterer) (*IdoFilterer, error) {
	contract, err := bindIdo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdoFilterer{contract: contract}, nil
}

// bindIdo binds a generic wrapper to an already deployed contract.
func bindIdo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ido *IdoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ido.Contract.IdoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ido *IdoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.Contract.IdoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ido *IdoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ido.Contract.IdoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ido *IdoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ido *IdoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ido *IdoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ido.Contract.contract.Transact(opts, method, params...)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Ido *IdoCaller) AdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "adminAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Ido *IdoSession) AdminAddr() (common.Address, error) {
	return _Ido.Contract.AdminAddr(&_Ido.CallOpts)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_Ido *IdoCallerSession) AdminAddr() (common.Address, error) {
	return _Ido.Contract.AdminAddr(&_Ido.CallOpts)
}

// AdminPoint is a free data retrieval call binding the contract method 0xc96224f6.
//
// Solidity: function adminPoint() view returns(uint256)
func (_Ido *IdoCaller) AdminPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "adminPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminPoint is a free data retrieval call binding the contract method 0xc96224f6.
//
// Solidity: function adminPoint() view returns(uint256)
func (_Ido *IdoSession) AdminPoint() (*big.Int, error) {
	return _Ido.Contract.AdminPoint(&_Ido.CallOpts)
}

// AdminPoint is a free data retrieval call binding the contract method 0xc96224f6.
//
// Solidity: function adminPoint() view returns(uint256)
func (_Ido *IdoCallerSession) AdminPoint() (*big.Int, error) {
	return _Ido.Contract.AdminPoint(&_Ido.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Ido *IdoCaller) Num(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "num")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Ido *IdoSession) Num() (*big.Int, error) {
	return _Ido.Contract.Num(&_Ido.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Ido *IdoCallerSession) Num() (*big.Int, error) {
	return _Ido.Contract.Num(&_Ido.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ido *IdoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ido *IdoSession) Owner() (common.Address, error) {
	return _Ido.Contract.Owner(&_Ido.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ido *IdoCallerSession) Owner() (common.Address, error) {
	return _Ido.Contract.Owner(&_Ido.CallOpts)
}

// QueryNum is a free data retrieval call binding the contract method 0x11ad84bb.
//
// Solidity: function queryNum() view returns(uint256)
func (_Ido *IdoCaller) QueryNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "queryNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryNum is a free data retrieval call binding the contract method 0x11ad84bb.
//
// Solidity: function queryNum() view returns(uint256)
func (_Ido *IdoSession) QueryNum() (*big.Int, error) {
	return _Ido.Contract.QueryNum(&_Ido.CallOpts)
}

// QueryNum is a free data retrieval call binding the contract method 0x11ad84bb.
//
// Solidity: function queryNum() view returns(uint256)
func (_Ido *IdoCallerSession) QueryNum() (*big.Int, error) {
	return _Ido.Contract.QueryNum(&_Ido.CallOpts)
}

// QueryProjectInfo is a free data retrieval call binding the contract method 0x9a35ae52.
//
// Solidity: function queryProjectInfo() view returns((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Ido *IdoCaller) QueryProjectInfo(opts *bind.CallOpts) (idopoolprojectInfo, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "queryProjectInfo")

	if err != nil {
		return *new(idopoolprojectInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(idopoolprojectInfo)).(*idopoolprojectInfo)

	return out0, err

}

// QueryProjectInfo is a free data retrieval call binding the contract method 0x9a35ae52.
//
// Solidity: function queryProjectInfo() view returns((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Ido *IdoSession) QueryProjectInfo() (idopoolprojectInfo, error) {
	return _Ido.Contract.QueryProjectInfo(&_Ido.CallOpts)
}

// QueryProjectInfo is a free data retrieval call binding the contract method 0x9a35ae52.
//
// Solidity: function queryProjectInfo() view returns((address,address,string,uint256,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Ido *IdoCallerSession) QueryProjectInfo() (idopoolprojectInfo, error) {
	return _Ido.Contract.QueryProjectInfo(&_Ido.CallOpts)
}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((uint256,uint256)[])
func (_Ido *IdoCaller) QueryUser(opts *bind.CallOpts) ([]idopoollockUser, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "queryUser")

	if err != nil {
		return *new([]idopoollockUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]idopoollockUser)).(*[]idopoollockUser)

	return out0, err

}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((uint256,uint256)[])
func (_Ido *IdoSession) QueryUser() ([]idopoollockUser, error) {
	return _Ido.Contract.QueryUser(&_Ido.CallOpts)
}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((uint256,uint256)[])
func (_Ido *IdoCallerSession) QueryUser() ([]idopoollockUser, error) {
	return _Ido.Contract.QueryUser(&_Ido.CallOpts)
}

// SendAddrA is a free data retrieval call binding the contract method 0xd41a58a0.
//
// Solidity: function sendAddrA() view returns(address)
func (_Ido *IdoCaller) SendAddrA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "sendAddrA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SendAddrA is a free data retrieval call binding the contract method 0xd41a58a0.
//
// Solidity: function sendAddrA() view returns(address)
func (_Ido *IdoSession) SendAddrA() (common.Address, error) {
	return _Ido.Contract.SendAddrA(&_Ido.CallOpts)
}

// SendAddrA is a free data retrieval call binding the contract method 0xd41a58a0.
//
// Solidity: function sendAddrA() view returns(address)
func (_Ido *IdoCallerSession) SendAddrA() (common.Address, error) {
	return _Ido.Contract.SendAddrA(&_Ido.CallOpts)
}

// SendAddrB is a free data retrieval call binding the contract method 0xf4ac9911.
//
// Solidity: function sendAddrB() view returns(address)
func (_Ido *IdoCaller) SendAddrB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "sendAddrB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SendAddrB is a free data retrieval call binding the contract method 0xf4ac9911.
//
// Solidity: function sendAddrB() view returns(address)
func (_Ido *IdoSession) SendAddrB() (common.Address, error) {
	return _Ido.Contract.SendAddrB(&_Ido.CallOpts)
}

// SendAddrB is a free data retrieval call binding the contract method 0xf4ac9911.
//
// Solidity: function sendAddrB() view returns(address)
func (_Ido *IdoCallerSession) SendAddrB() (common.Address, error) {
	return _Ido.Contract.SendAddrB(&_Ido.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Ido *IdoCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Ido *IdoSession) TokenA() (common.Address, error) {
	return _Ido.Contract.TokenA(&_Ido.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Ido *IdoCallerSession) TokenA() (common.Address, error) {
	return _Ido.Contract.TokenA(&_Ido.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Ido *IdoCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Ido *IdoSession) TokenB() (common.Address, error) {
	return _Ido.Contract.TokenB(&_Ido.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Ido *IdoCallerSession) TokenB() (common.Address, error) {
	return _Ido.Contract.TokenB(&_Ido.CallOpts)
}

// Withdraw is a free data retrieval call binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() view returns(uint256)
func (_Ido *IdoCaller) Withdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ido.contract.Call(opts, &out, "withdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdraw is a free data retrieval call binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() view returns(uint256)
func (_Ido *IdoSession) Withdraw() (*big.Int, error) {
	return _Ido.Contract.Withdraw(&_Ido.CallOpts)
}

// Withdraw is a free data retrieval call binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() view returns(uint256)
func (_Ido *IdoCallerSession) Withdraw() (*big.Int, error) {
	return _Ido.Contract.Withdraw(&_Ido.CallOpts)
}

// AdvanceRelease is a paid mutator transaction binding the contract method 0x29f02852.
//
// Solidity: function advanceRelease() returns()
func (_Ido *IdoTransactor) AdvanceRelease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "advanceRelease")
}

// AdvanceRelease is a paid mutator transaction binding the contract method 0x29f02852.
//
// Solidity: function advanceRelease() returns()
func (_Ido *IdoSession) AdvanceRelease() (*types.Transaction, error) {
	return _Ido.Contract.AdvanceRelease(&_Ido.TransactOpts)
}

// AdvanceRelease is a paid mutator transaction binding the contract method 0x29f02852.
//
// Solidity: function advanceRelease() returns()
func (_Ido *IdoTransactorSession) AdvanceRelease() (*types.Transaction, error) {
	return _Ido.Contract.AdvanceRelease(&_Ido.TransactOpts)
}

// ReceiveOwner is a paid mutator transaction binding the contract method 0x598643ce.
//
// Solidity: function receiveOwner() returns()
func (_Ido *IdoTransactor) ReceiveOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "receiveOwner")
}

// ReceiveOwner is a paid mutator transaction binding the contract method 0x598643ce.
//
// Solidity: function receiveOwner() returns()
func (_Ido *IdoSession) ReceiveOwner() (*types.Transaction, error) {
	return _Ido.Contract.ReceiveOwner(&_Ido.TransactOpts)
}

// ReceiveOwner is a paid mutator transaction binding the contract method 0x598643ce.
//
// Solidity: function receiveOwner() returns()
func (_Ido *IdoTransactorSession) ReceiveOwner() (*types.Transaction, error) {
	return _Ido.Contract.ReceiveOwner(&_Ido.TransactOpts)
}

// ReceiveUser is a paid mutator transaction binding the contract method 0x83955a88.
//
// Solidity: function receiveUser() returns()
func (_Ido *IdoTransactor) ReceiveUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "receiveUser")
}

// ReceiveUser is a paid mutator transaction binding the contract method 0x83955a88.
//
// Solidity: function receiveUser() returns()
func (_Ido *IdoSession) ReceiveUser() (*types.Transaction, error) {
	return _Ido.Contract.ReceiveUser(&_Ido.TransactOpts)
}

// ReceiveUser is a paid mutator transaction binding the contract method 0x83955a88.
//
// Solidity: function receiveUser() returns()
func (_Ido *IdoTransactorSession) ReceiveUser() (*types.Transaction, error) {
	return _Ido.Contract.ReceiveUser(&_Ido.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ido *IdoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ido *IdoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ido.Contract.RenounceOwnership(&_Ido.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ido *IdoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ido.Contract.RenounceOwnership(&_Ido.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Ido *IdoTransactor) SetFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "setFee", fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Ido *IdoSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _Ido.Contract.SetFee(&_Ido.TransactOpts, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Ido *IdoTransactorSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _Ido.Contract.SetFee(&_Ido.TransactOpts, fee)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _amountIn) returns()
func (_Ido *IdoTransactor) Swap(opts *bind.TransactOpts, _amountIn *big.Int) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "swap", _amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _amountIn) returns()
func (_Ido *IdoSession) Swap(_amountIn *big.Int) (*types.Transaction, error) {
	return _Ido.Contract.Swap(&_Ido.TransactOpts, _amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _amountIn) returns()
func (_Ido *IdoTransactorSession) Swap(_amountIn *big.Int) (*types.Transaction, error) {
	return _Ido.Contract.Swap(&_Ido.TransactOpts, _amountIn)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ido *IdoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ido *IdoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ido.Contract.TransferOwnership(&_Ido.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ido *IdoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ido.Contract.TransferOwnership(&_Ido.TransactOpts, newOwner)
}

// IdoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ido contract.
type IdoOwnershipTransferredIterator struct {
	Event *IdoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdoOwnershipTransferred)
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
		it.Event = new(IdoOwnershipTransferred)
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
func (it *IdoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdoOwnershipTransferred represents a OwnershipTransferred event raised by the Ido contract.
type IdoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ido *IdoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ido.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdoOwnershipTransferredIterator{contract: _Ido.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ido *IdoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ido.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdoOwnershipTransferred)
				if err := _Ido.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ido *IdoFilterer) ParseOwnershipTransferred(log types.Log) (*IdoOwnershipTransferred, error) {
	event := new(IdoOwnershipTransferred)
	if err := _Ido.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdoDebugIterator is returned from FilterDebug and is used to iterate over the raw logs and unpacked data for Debug events raised by the Ido contract.
type IdoDebugIterator struct {
	Event *IdoDebug // Event containing the contract specifics and raw log

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
func (it *IdoDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdoDebug)
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
		it.Event = new(IdoDebug)
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
func (it *IdoDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdoDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdoDebug represents a Debug event raised by the Ido contract.
type IdoDebug struct {
	Name   string
	Msg    common.Address
	Addr   common.Address
	Exec   *big.Int
	Number *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebug is a free log retrieval operation binding the contract event 0x5268306319b571359cc341005eac1e0d3b74776dd636ab8b3e1ef2e2990b5450.
//
// Solidity: event debug(string name, address msg, address addr, uint256 exec, uint256 number, uint256 time)
func (_Ido *IdoFilterer) FilterDebug(opts *bind.FilterOpts) (*IdoDebugIterator, error) {

	logs, sub, err := _Ido.contract.FilterLogs(opts, "debug")
	if err != nil {
		return nil, err
	}
	return &IdoDebugIterator{contract: _Ido.contract, event: "debug", logs: logs, sub: sub}, nil
}

// WatchDebug is a free log subscription operation binding the contract event 0x5268306319b571359cc341005eac1e0d3b74776dd636ab8b3e1ef2e2990b5450.
//
// Solidity: event debug(string name, address msg, address addr, uint256 exec, uint256 number, uint256 time)
func (_Ido *IdoFilterer) WatchDebug(opts *bind.WatchOpts, sink chan<- *IdoDebug) (event.Subscription, error) {

	logs, sub, err := _Ido.contract.WatchLogs(opts, "debug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdoDebug)
				if err := _Ido.contract.UnpackLog(event, "debug", log); err != nil {
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

// ParseDebug is a log parse operation binding the contract event 0x5268306319b571359cc341005eac1e0d3b74776dd636ab8b3e1ef2e2990b5450.
//
// Solidity: event debug(string name, address msg, address addr, uint256 exec, uint256 number, uint256 time)
func (_Ido *IdoFilterer) ParseDebug(log types.Log) (*IdoDebug, error) {
	event := new(IdoDebug)
	if err := _Ido.contract.UnpackLog(event, "debug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
