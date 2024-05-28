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
	_ = abi.ConvertType
)

// YiSinEBookMetaData contains all meta data concerning the YiSinEBook contract.
var YiSinEBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bookInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRentTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"booksOnRent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"}],\"name\":\"burnEBook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isAddressHaveTokenId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentTime\",\"type\":\"uint256\"}],\"name\":\"rentBook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"renterRentInfoIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"}],\"name\":\"returnBook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupplyBook\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bookAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"uploadEBook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YiSinEBookABI is the input ABI used to generate the binding from.
// Deprecated: Use YiSinEBookMetaData.ABI instead.
var YiSinEBookABI = YiSinEBookMetaData.ABI

// YiSinEBook is an auto generated Go binding around an Ethereum contract.
type YiSinEBook struct {
	YiSinEBookCaller     // Read-only binding to the contract
	YiSinEBookTransactor // Write-only binding to the contract
	YiSinEBookFilterer   // Log filterer for contract events
}

// YiSinEBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type YiSinEBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YiSinEBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YiSinEBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YiSinEBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YiSinEBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YiSinEBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YiSinEBookSession struct {
	Contract     *YiSinEBook       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YiSinEBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YiSinEBookCallerSession struct {
	Contract *YiSinEBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YiSinEBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YiSinEBookTransactorSession struct {
	Contract     *YiSinEBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YiSinEBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type YiSinEBookRaw struct {
	Contract *YiSinEBook // Generic contract binding to access the raw methods on
}

// YiSinEBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YiSinEBookCallerRaw struct {
	Contract *YiSinEBookCaller // Generic read-only contract binding to access the raw methods on
}

// YiSinEBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YiSinEBookTransactorRaw struct {
	Contract *YiSinEBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYiSinEBook creates a new instance of YiSinEBook, bound to a specific deployed contract.
func NewYiSinEBook(address common.Address, backend bind.ContractBackend) (*YiSinEBook, error) {
	contract, err := bindYiSinEBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YiSinEBook{YiSinEBookCaller: YiSinEBookCaller{contract: contract}, YiSinEBookTransactor: YiSinEBookTransactor{contract: contract}, YiSinEBookFilterer: YiSinEBookFilterer{contract: contract}}, nil
}

// NewYiSinEBookCaller creates a new read-only instance of YiSinEBook, bound to a specific deployed contract.
func NewYiSinEBookCaller(address common.Address, caller bind.ContractCaller) (*YiSinEBookCaller, error) {
	contract, err := bindYiSinEBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookCaller{contract: contract}, nil
}

// NewYiSinEBookTransactor creates a new write-only instance of YiSinEBook, bound to a specific deployed contract.
func NewYiSinEBookTransactor(address common.Address, transactor bind.ContractTransactor) (*YiSinEBookTransactor, error) {
	contract, err := bindYiSinEBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookTransactor{contract: contract}, nil
}

// NewYiSinEBookFilterer creates a new log filterer instance of YiSinEBook, bound to a specific deployed contract.
func NewYiSinEBookFilterer(address common.Address, filterer bind.ContractFilterer) (*YiSinEBookFilterer, error) {
	contract, err := bindYiSinEBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookFilterer{contract: contract}, nil
}

// bindYiSinEBook binds a generic wrapper to an already deployed contract.
func bindYiSinEBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YiSinEBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YiSinEBook *YiSinEBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YiSinEBook.Contract.YiSinEBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YiSinEBook *YiSinEBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YiSinEBook.Contract.YiSinEBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YiSinEBook *YiSinEBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YiSinEBook.Contract.YiSinEBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YiSinEBook *YiSinEBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YiSinEBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YiSinEBook *YiSinEBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YiSinEBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YiSinEBook *YiSinEBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YiSinEBook.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YiSinEBook *YiSinEBookCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YiSinEBook *YiSinEBookSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.BalanceOf(&_YiSinEBook.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YiSinEBook *YiSinEBookCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.BalanceOf(&_YiSinEBook.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YiSinEBook *YiSinEBookCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YiSinEBook *YiSinEBookSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _YiSinEBook.Contract.BalanceOfBatch(&_YiSinEBook.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YiSinEBook *YiSinEBookCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _YiSinEBook.Contract.BalanceOfBatch(&_YiSinEBook.CallOpts, accounts, ids)
}

// BookInfos is a free data retrieval call binding the contract method 0x0fc5fc4a.
//
// Solidity: function bookInfos(uint256 ) view returns(address writer, uint256 supplyAmount, uint256 rentPrice, uint256 maxRentTime)
func (_YiSinEBook *YiSinEBookCaller) BookInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Writer       common.Address
	SupplyAmount *big.Int
	RentPrice    *big.Int
	MaxRentTime  *big.Int
}, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "bookInfos", arg0)

	outstruct := new(struct {
		Writer       common.Address
		SupplyAmount *big.Int
		RentPrice    *big.Int
		MaxRentTime  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Writer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SupplyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RentPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxRentTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BookInfos is a free data retrieval call binding the contract method 0x0fc5fc4a.
//
// Solidity: function bookInfos(uint256 ) view returns(address writer, uint256 supplyAmount, uint256 rentPrice, uint256 maxRentTime)
func (_YiSinEBook *YiSinEBookSession) BookInfos(arg0 *big.Int) (struct {
	Writer       common.Address
	SupplyAmount *big.Int
	RentPrice    *big.Int
	MaxRentTime  *big.Int
}, error) {
	return _YiSinEBook.Contract.BookInfos(&_YiSinEBook.CallOpts, arg0)
}

// BookInfos is a free data retrieval call binding the contract method 0x0fc5fc4a.
//
// Solidity: function bookInfos(uint256 ) view returns(address writer, uint256 supplyAmount, uint256 rentPrice, uint256 maxRentTime)
func (_YiSinEBook *YiSinEBookCallerSession) BookInfos(arg0 *big.Int) (struct {
	Writer       common.Address
	SupplyAmount *big.Int
	RentPrice    *big.Int
	MaxRentTime  *big.Int
}, error) {
	return _YiSinEBook.Contract.BookInfos(&_YiSinEBook.CallOpts, arg0)
}

// BooksOnRent is a free data retrieval call binding the contract method 0xd42881de.
//
// Solidity: function booksOnRent(uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookCaller) BooksOnRent(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "booksOnRent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BooksOnRent is a free data retrieval call binding the contract method 0xd42881de.
//
// Solidity: function booksOnRent(uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookSession) BooksOnRent(arg0 *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.BooksOnRent(&_YiSinEBook.CallOpts, arg0)
}

// BooksOnRent is a free data retrieval call binding the contract method 0xd42881de.
//
// Solidity: function booksOnRent(uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookCallerSession) BooksOnRent(arg0 *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.BooksOnRent(&_YiSinEBook.CallOpts, arg0)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_YiSinEBook *YiSinEBookCaller) CheckUpkeep(opts *bind.CallOpts, checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "checkUpkeep", checkData)

	outstruct := new(struct {
		UpkeepNeeded bool
		PerformData  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpkeepNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PerformData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_YiSinEBook *YiSinEBookSession) CheckUpkeep(checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _YiSinEBook.Contract.CheckUpkeep(&_YiSinEBook.CallOpts, checkData)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_YiSinEBook *YiSinEBookCallerSession) CheckUpkeep(checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _YiSinEBook.Contract.CheckUpkeep(&_YiSinEBook.CallOpts, checkData)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_YiSinEBook *YiSinEBookCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_YiSinEBook *YiSinEBookSession) Fee() (*big.Int, error) {
	return _YiSinEBook.Contract.Fee(&_YiSinEBook.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_YiSinEBook *YiSinEBookCallerSession) Fee() (*big.Int, error) {
	return _YiSinEBook.Contract.Fee(&_YiSinEBook.CallOpts)
}

// IsAddressHaveTokenId is a free data retrieval call binding the contract method 0x61e5b62c.
//
// Solidity: function isAddressHaveTokenId(address signer, uint256 id) view returns(bool)
func (_YiSinEBook *YiSinEBookCaller) IsAddressHaveTokenId(opts *bind.CallOpts, signer common.Address, id *big.Int) (bool, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "isAddressHaveTokenId", signer, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAddressHaveTokenId is a free data retrieval call binding the contract method 0x61e5b62c.
//
// Solidity: function isAddressHaveTokenId(address signer, uint256 id) view returns(bool)
func (_YiSinEBook *YiSinEBookSession) IsAddressHaveTokenId(signer common.Address, id *big.Int) (bool, error) {
	return _YiSinEBook.Contract.IsAddressHaveTokenId(&_YiSinEBook.CallOpts, signer, id)
}

// IsAddressHaveTokenId is a free data retrieval call binding the contract method 0x61e5b62c.
//
// Solidity: function isAddressHaveTokenId(address signer, uint256 id) view returns(bool)
func (_YiSinEBook *YiSinEBookCallerSession) IsAddressHaveTokenId(signer common.Address, id *big.Int) (bool, error) {
	return _YiSinEBook.Contract.IsAddressHaveTokenId(&_YiSinEBook.CallOpts, signer, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YiSinEBook *YiSinEBookCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YiSinEBook *YiSinEBookSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _YiSinEBook.Contract.IsApprovedForAll(&_YiSinEBook.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YiSinEBook *YiSinEBookCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _YiSinEBook.Contract.IsApprovedForAll(&_YiSinEBook.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YiSinEBook *YiSinEBookCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YiSinEBook *YiSinEBookSession) Owner() (common.Address, error) {
	return _YiSinEBook.Contract.Owner(&_YiSinEBook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YiSinEBook *YiSinEBookCallerSession) Owner() (common.Address, error) {
	return _YiSinEBook.Contract.Owner(&_YiSinEBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YiSinEBook *YiSinEBookCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YiSinEBook *YiSinEBookSession) Paused() (bool, error) {
	return _YiSinEBook.Contract.Paused(&_YiSinEBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YiSinEBook *YiSinEBookCallerSession) Paused() (bool, error) {
	return _YiSinEBook.Contract.Paused(&_YiSinEBook.CallOpts)
}

// RentInfos is a free data retrieval call binding the contract method 0x8aa087f1.
//
// Solidity: function rentInfos(uint256 , uint256 ) view returns(address renter, uint256 tokenId, uint256 endTime)
func (_YiSinEBook *YiSinEBookCaller) RentInfos(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Renter  common.Address
	TokenId *big.Int
	EndTime *big.Int
}, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "rentInfos", arg0, arg1)

	outstruct := new(struct {
		Renter  common.Address
		TokenId *big.Int
		EndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Renter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RentInfos is a free data retrieval call binding the contract method 0x8aa087f1.
//
// Solidity: function rentInfos(uint256 , uint256 ) view returns(address renter, uint256 tokenId, uint256 endTime)
func (_YiSinEBook *YiSinEBookSession) RentInfos(arg0 *big.Int, arg1 *big.Int) (struct {
	Renter  common.Address
	TokenId *big.Int
	EndTime *big.Int
}, error) {
	return _YiSinEBook.Contract.RentInfos(&_YiSinEBook.CallOpts, arg0, arg1)
}

// RentInfos is a free data retrieval call binding the contract method 0x8aa087f1.
//
// Solidity: function rentInfos(uint256 , uint256 ) view returns(address renter, uint256 tokenId, uint256 endTime)
func (_YiSinEBook *YiSinEBookCallerSession) RentInfos(arg0 *big.Int, arg1 *big.Int) (struct {
	Renter  common.Address
	TokenId *big.Int
	EndTime *big.Int
}, error) {
	return _YiSinEBook.Contract.RentInfos(&_YiSinEBook.CallOpts, arg0, arg1)
}

// RenterRentInfoIndex is a free data retrieval call binding the contract method 0xbcb3e9ae.
//
// Solidity: function renterRentInfoIndex(address , uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookCaller) RenterRentInfoIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "renterRentInfoIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RenterRentInfoIndex is a free data retrieval call binding the contract method 0xbcb3e9ae.
//
// Solidity: function renterRentInfoIndex(address , uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookSession) RenterRentInfoIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.RenterRentInfoIndex(&_YiSinEBook.CallOpts, arg0, arg1)
}

// RenterRentInfoIndex is a free data retrieval call binding the contract method 0xbcb3e9ae.
//
// Solidity: function renterRentInfoIndex(address , uint256 ) view returns(uint256)
func (_YiSinEBook *YiSinEBookCallerSession) RenterRentInfoIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _YiSinEBook.Contract.RenterRentInfoIndex(&_YiSinEBook.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YiSinEBook *YiSinEBookCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YiSinEBook *YiSinEBookSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YiSinEBook.Contract.SupportsInterface(&_YiSinEBook.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YiSinEBook *YiSinEBookCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YiSinEBook.Contract.SupportsInterface(&_YiSinEBook.CallOpts, interfaceId)
}

// TotalSupplyBook is a free data retrieval call binding the contract method 0x8e109a27.
//
// Solidity: function totalSupplyBook() view returns(uint256)
func (_YiSinEBook *YiSinEBookCaller) TotalSupplyBook(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "totalSupplyBook")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyBook is a free data retrieval call binding the contract method 0x8e109a27.
//
// Solidity: function totalSupplyBook() view returns(uint256)
func (_YiSinEBook *YiSinEBookSession) TotalSupplyBook() (*big.Int, error) {
	return _YiSinEBook.Contract.TotalSupplyBook(&_YiSinEBook.CallOpts)
}

// TotalSupplyBook is a free data retrieval call binding the contract method 0x8e109a27.
//
// Solidity: function totalSupplyBook() view returns(uint256)
func (_YiSinEBook *YiSinEBookCallerSession) TotalSupplyBook() (*big.Int, error) {
	return _YiSinEBook.Contract.TotalSupplyBook(&_YiSinEBook.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YiSinEBook *YiSinEBookCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _YiSinEBook.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YiSinEBook *YiSinEBookSession) Uri(arg0 *big.Int) (string, error) {
	return _YiSinEBook.Contract.Uri(&_YiSinEBook.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YiSinEBook *YiSinEBookCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _YiSinEBook.Contract.Uri(&_YiSinEBook.CallOpts, arg0)
}

// BurnEBook is a paid mutator transaction binding the contract method 0x4c019c38.
//
// Solidity: function burnEBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookTransactor) BurnEBook(opts *bind.TransactOpts, bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "burnEBook", bookId)
}

// BurnEBook is a paid mutator transaction binding the contract method 0x4c019c38.
//
// Solidity: function burnEBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookSession) BurnEBook(bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.BurnEBook(&_YiSinEBook.TransactOpts, bookId)
}

// BurnEBook is a paid mutator transaction binding the contract method 0x4c019c38.
//
// Solidity: function burnEBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) BurnEBook(bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.BurnEBook(&_YiSinEBook.TransactOpts, bookId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YiSinEBook *YiSinEBookTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YiSinEBook *YiSinEBookSession) Pause() (*types.Transaction, error) {
	return _YiSinEBook.Contract.Pause(&_YiSinEBook.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_YiSinEBook *YiSinEBookTransactorSession) Pause() (*types.Transaction, error) {
	return _YiSinEBook.Contract.Pause(&_YiSinEBook.TransactOpts)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_YiSinEBook *YiSinEBookTransactor) PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "performUpkeep", performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_YiSinEBook *YiSinEBookSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.PerformUpkeep(&_YiSinEBook.TransactOpts, performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.PerformUpkeep(&_YiSinEBook.TransactOpts, performData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YiSinEBook *YiSinEBookTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YiSinEBook *YiSinEBookSession) RenounceOwnership() (*types.Transaction, error) {
	return _YiSinEBook.Contract.RenounceOwnership(&_YiSinEBook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YiSinEBook *YiSinEBookTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YiSinEBook.Contract.RenounceOwnership(&_YiSinEBook.TransactOpts)
}

// RentBook is a paid mutator transaction binding the contract method 0xd43a4276.
//
// Solidity: function rentBook(uint256 bookId, uint256 rentTime) payable returns()
func (_YiSinEBook *YiSinEBookTransactor) RentBook(opts *bind.TransactOpts, bookId *big.Int, rentTime *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "rentBook", bookId, rentTime)
}

// RentBook is a paid mutator transaction binding the contract method 0xd43a4276.
//
// Solidity: function rentBook(uint256 bookId, uint256 rentTime) payable returns()
func (_YiSinEBook *YiSinEBookSession) RentBook(bookId *big.Int, rentTime *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.RentBook(&_YiSinEBook.TransactOpts, bookId, rentTime)
}

// RentBook is a paid mutator transaction binding the contract method 0xd43a4276.
//
// Solidity: function rentBook(uint256 bookId, uint256 rentTime) payable returns()
func (_YiSinEBook *YiSinEBookTransactorSession) RentBook(bookId *big.Int, rentTime *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.RentBook(&_YiSinEBook.TransactOpts, bookId, rentTime)
}

// ReturnBook is a paid mutator transaction binding the contract method 0xca5140c9.
//
// Solidity: function returnBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookTransactor) ReturnBook(opts *bind.TransactOpts, bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "returnBook", bookId)
}

// ReturnBook is a paid mutator transaction binding the contract method 0xca5140c9.
//
// Solidity: function returnBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookSession) ReturnBook(bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.ReturnBook(&_YiSinEBook.TransactOpts, bookId)
}

// ReturnBook is a paid mutator transaction binding the contract method 0xca5140c9.
//
// Solidity: function returnBook(uint256 bookId) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) ReturnBook(bookId *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.ReturnBook(&_YiSinEBook.TransactOpts, bookId)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YiSinEBook *YiSinEBookTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YiSinEBook *YiSinEBookSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SafeBatchTransferFrom(&_YiSinEBook.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SafeBatchTransferFrom(&_YiSinEBook.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YiSinEBook *YiSinEBookTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YiSinEBook *YiSinEBookSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SafeTransferFrom(&_YiSinEBook.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SafeTransferFrom(&_YiSinEBook.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YiSinEBook *YiSinEBookTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YiSinEBook *YiSinEBookSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SetApprovalForAll(&_YiSinEBook.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SetApprovalForAll(&_YiSinEBook.TransactOpts, operator, approved)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_YiSinEBook *YiSinEBookTransactor) SetURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_YiSinEBook *YiSinEBookSession) SetURI(newuri string) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SetURI(&_YiSinEBook.TransactOpts, newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) SetURI(newuri string) (*types.Transaction, error) {
	return _YiSinEBook.Contract.SetURI(&_YiSinEBook.TransactOpts, newuri)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YiSinEBook *YiSinEBookTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YiSinEBook *YiSinEBookSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YiSinEBook.Contract.TransferOwnership(&_YiSinEBook.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YiSinEBook.Contract.TransferOwnership(&_YiSinEBook.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YiSinEBook *YiSinEBookTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YiSinEBook *YiSinEBookSession) Unpause() (*types.Transaction, error) {
	return _YiSinEBook.Contract.Unpause(&_YiSinEBook.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_YiSinEBook *YiSinEBookTransactorSession) Unpause() (*types.Transaction, error) {
	return _YiSinEBook.Contract.Unpause(&_YiSinEBook.TransactOpts)
}

// UploadEBook is a paid mutator transaction binding the contract method 0xa358aaab.
//
// Solidity: function uploadEBook(uint256 bookAmount, address uploader, uint256 price, uint256 time) returns()
func (_YiSinEBook *YiSinEBookTransactor) UploadEBook(opts *bind.TransactOpts, bookAmount *big.Int, uploader common.Address, price *big.Int, time *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.contract.Transact(opts, "uploadEBook", bookAmount, uploader, price, time)
}

// UploadEBook is a paid mutator transaction binding the contract method 0xa358aaab.
//
// Solidity: function uploadEBook(uint256 bookAmount, address uploader, uint256 price, uint256 time) returns()
func (_YiSinEBook *YiSinEBookSession) UploadEBook(bookAmount *big.Int, uploader common.Address, price *big.Int, time *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.UploadEBook(&_YiSinEBook.TransactOpts, bookAmount, uploader, price, time)
}

// UploadEBook is a paid mutator transaction binding the contract method 0xa358aaab.
//
// Solidity: function uploadEBook(uint256 bookAmount, address uploader, uint256 price, uint256 time) returns()
func (_YiSinEBook *YiSinEBookTransactorSession) UploadEBook(bookAmount *big.Int, uploader common.Address, price *big.Int, time *big.Int) (*types.Transaction, error) {
	return _YiSinEBook.Contract.UploadEBook(&_YiSinEBook.TransactOpts, bookAmount, uploader, price, time)
}

// YiSinEBookApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the YiSinEBook contract.
type YiSinEBookApprovalForAllIterator struct {
	Event *YiSinEBookApprovalForAll // Event containing the contract specifics and raw log

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
func (it *YiSinEBookApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookApprovalForAll)
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
		it.Event = new(YiSinEBookApprovalForAll)
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
func (it *YiSinEBookApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookApprovalForAll represents a ApprovalForAll event raised by the YiSinEBook contract.
type YiSinEBookApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YiSinEBook *YiSinEBookFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*YiSinEBookApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookApprovalForAllIterator{contract: _YiSinEBook.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YiSinEBook *YiSinEBookFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *YiSinEBookApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookApprovalForAll)
				if err := _YiSinEBook.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YiSinEBook *YiSinEBookFilterer) ParseApprovalForAll(log types.Log) (*YiSinEBookApprovalForAll, error) {
	event := new(YiSinEBookApprovalForAll)
	if err := _YiSinEBook.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YiSinEBook contract.
type YiSinEBookOwnershipTransferredIterator struct {
	Event *YiSinEBookOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YiSinEBookOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookOwnershipTransferred)
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
		it.Event = new(YiSinEBookOwnershipTransferred)
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
func (it *YiSinEBookOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookOwnershipTransferred represents a OwnershipTransferred event raised by the YiSinEBook contract.
type YiSinEBookOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YiSinEBook *YiSinEBookFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YiSinEBookOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookOwnershipTransferredIterator{contract: _YiSinEBook.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YiSinEBook *YiSinEBookFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YiSinEBookOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookOwnershipTransferred)
				if err := _YiSinEBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_YiSinEBook *YiSinEBookFilterer) ParseOwnershipTransferred(log types.Log) (*YiSinEBookOwnershipTransferred, error) {
	event := new(YiSinEBookOwnershipTransferred)
	if err := _YiSinEBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the YiSinEBook contract.
type YiSinEBookPausedIterator struct {
	Event *YiSinEBookPaused // Event containing the contract specifics and raw log

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
func (it *YiSinEBookPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookPaused)
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
		it.Event = new(YiSinEBookPaused)
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
func (it *YiSinEBookPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookPaused represents a Paused event raised by the YiSinEBook contract.
type YiSinEBookPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_YiSinEBook *YiSinEBookFilterer) FilterPaused(opts *bind.FilterOpts) (*YiSinEBookPausedIterator, error) {

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &YiSinEBookPausedIterator{contract: _YiSinEBook.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_YiSinEBook *YiSinEBookFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *YiSinEBookPaused) (event.Subscription, error) {

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookPaused)
				if err := _YiSinEBook.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_YiSinEBook *YiSinEBookFilterer) ParsePaused(log types.Log) (*YiSinEBookPaused, error) {
	event := new(YiSinEBookPaused)
	if err := _YiSinEBook.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the YiSinEBook contract.
type YiSinEBookTransferBatchIterator struct {
	Event *YiSinEBookTransferBatch // Event containing the contract specifics and raw log

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
func (it *YiSinEBookTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookTransferBatch)
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
		it.Event = new(YiSinEBookTransferBatch)
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
func (it *YiSinEBookTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookTransferBatch represents a TransferBatch event raised by the YiSinEBook contract.
type YiSinEBookTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YiSinEBook *YiSinEBookFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YiSinEBookTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookTransferBatchIterator{contract: _YiSinEBook.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YiSinEBook *YiSinEBookFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *YiSinEBookTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookTransferBatch)
				if err := _YiSinEBook.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YiSinEBook *YiSinEBookFilterer) ParseTransferBatch(log types.Log) (*YiSinEBookTransferBatch, error) {
	event := new(YiSinEBookTransferBatch)
	if err := _YiSinEBook.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the YiSinEBook contract.
type YiSinEBookTransferSingleIterator struct {
	Event *YiSinEBookTransferSingle // Event containing the contract specifics and raw log

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
func (it *YiSinEBookTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookTransferSingle)
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
		it.Event = new(YiSinEBookTransferSingle)
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
func (it *YiSinEBookTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookTransferSingle represents a TransferSingle event raised by the YiSinEBook contract.
type YiSinEBookTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YiSinEBook *YiSinEBookFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YiSinEBookTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookTransferSingleIterator{contract: _YiSinEBook.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YiSinEBook *YiSinEBookFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *YiSinEBookTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookTransferSingle)
				if err := _YiSinEBook.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YiSinEBook *YiSinEBookFilterer) ParseTransferSingle(log types.Log) (*YiSinEBookTransferSingle, error) {
	event := new(YiSinEBookTransferSingle)
	if err := _YiSinEBook.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the YiSinEBook contract.
type YiSinEBookURIIterator struct {
	Event *YiSinEBookURI // Event containing the contract specifics and raw log

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
func (it *YiSinEBookURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookURI)
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
		it.Event = new(YiSinEBookURI)
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
func (it *YiSinEBookURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookURI represents a URI event raised by the YiSinEBook contract.
type YiSinEBookURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YiSinEBook *YiSinEBookFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*YiSinEBookURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &YiSinEBookURIIterator{contract: _YiSinEBook.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YiSinEBook *YiSinEBookFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *YiSinEBookURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookURI)
				if err := _YiSinEBook.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YiSinEBook *YiSinEBookFilterer) ParseURI(log types.Log) (*YiSinEBookURI, error) {
	event := new(YiSinEBookURI)
	if err := _YiSinEBook.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YiSinEBookUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the YiSinEBook contract.
type YiSinEBookUnpausedIterator struct {
	Event *YiSinEBookUnpaused // Event containing the contract specifics and raw log

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
func (it *YiSinEBookUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YiSinEBookUnpaused)
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
		it.Event = new(YiSinEBookUnpaused)
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
func (it *YiSinEBookUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YiSinEBookUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YiSinEBookUnpaused represents a Unpaused event raised by the YiSinEBook contract.
type YiSinEBookUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_YiSinEBook *YiSinEBookFilterer) FilterUnpaused(opts *bind.FilterOpts) (*YiSinEBookUnpausedIterator, error) {

	logs, sub, err := _YiSinEBook.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &YiSinEBookUnpausedIterator{contract: _YiSinEBook.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_YiSinEBook *YiSinEBookFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *YiSinEBookUnpaused) (event.Subscription, error) {

	logs, sub, err := _YiSinEBook.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YiSinEBookUnpaused)
				if err := _YiSinEBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_YiSinEBook *YiSinEBookFilterer) ParseUnpaused(log types.Log) (*YiSinEBookUnpaused, error) {
	event := new(YiSinEBookUnpaused)
	if err := _YiSinEBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
