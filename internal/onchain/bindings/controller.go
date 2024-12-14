// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ControllerDataDoc is an auto generated low-level Go binding around an user-defined struct.
type ControllerDataDoc struct {
	Id          string
	HashContent string
}

// ControllerUploadSession is an auto generated low-level Go binding around an user-defined struct.
type ControllerUploadSession struct {
	Id        *big.Int
	User      common.Address
	Proof     string
	Confirmed bool
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pcspAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hashContent\",\"type\":\"string\"}],\"name\":\"GeneDataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"GeneNFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"PCSPRewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"UploadData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contentHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geneNFT\",\"outputs\":[{\"internalType\":\"contractGeneNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashContent\",\"type\":\"string\"}],\"internalType\":\"structController.DataDoc\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getSession\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"internalType\":\"structController.UploadSession\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pcspToken\",\"outputs\":[{\"internalType\":\"contractPostCovidStrokePrevention\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"uploadData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCaller) GeneNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "geneNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCallerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCaller) GetDoc(opts *bind.CallOpts, docId string) (ControllerDataDoc, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getDoc", docId)

	if err != nil {
		return *new(ControllerDataDoc), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerDataDoc)).(*ControllerDataDoc)

	return out0, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCallerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerCaller) GetSession(opts *bind.CallOpts, sessionId *big.Int) (ControllerUploadSession, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getSession", sessionId)

	if err != nil {
		return *new(ControllerUploadSession), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerUploadSession)).(*ControllerUploadSession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerCallerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCaller) PcspToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "pcspToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCallerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns()
func (_Controller *ControllerTransactor) Confirm(opts *bind.TransactOpts, docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "confirm", docId, contentHash, proof, sessionId, riskScore)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns()
func (_Controller *ControllerSession) Confirm(docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, proof, sessionId, riskScore)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns()
func (_Controller *ControllerTransactorSession) Confirm(docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, proof, sessionId, riskScore)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactor) UploadData(opts *bind.TransactOpts, docId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "uploadData", docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactorSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// ControllerGeneDataSubmittedIterator is returned from FilterGeneDataSubmitted and is used to iterate over the raw logs and unpacked data for GeneDataSubmitted events raised by the Controller contract.
type ControllerGeneDataSubmittedIterator struct {
	Event *ControllerGeneDataSubmitted // Event containing the contract specifics and raw log

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
func (it *ControllerGeneDataSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerGeneDataSubmitted)
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
		it.Event = new(ControllerGeneDataSubmitted)
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
func (it *ControllerGeneDataSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerGeneDataSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerGeneDataSubmitted represents a GeneDataSubmitted event raised by the Controller contract.
type ControllerGeneDataSubmitted struct {
	DocId       common.Hash
	User        common.Address
	HashContent string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGeneDataSubmitted is a free log retrieval operation binding the contract event 0xa44c9df778cfa6cc467846bb2c56756453c6c25c46977927881c4d88775d10e5.
//
// Solidity: event GeneDataSubmitted(string indexed docId, address indexed user, string hashContent)
func (_Controller *ControllerFilterer) FilterGeneDataSubmitted(opts *bind.FilterOpts, docId []string, user []common.Address) (*ControllerGeneDataSubmittedIterator, error) {

	var docIdRule []interface{}
	for _, docIdItem := range docId {
		docIdRule = append(docIdRule, docIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "GeneDataSubmitted", docIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ControllerGeneDataSubmittedIterator{contract: _Controller.contract, event: "GeneDataSubmitted", logs: logs, sub: sub}, nil
}

// WatchGeneDataSubmitted is a free log subscription operation binding the contract event 0xa44c9df778cfa6cc467846bb2c56756453c6c25c46977927881c4d88775d10e5.
//
// Solidity: event GeneDataSubmitted(string indexed docId, address indexed user, string hashContent)
func (_Controller *ControllerFilterer) WatchGeneDataSubmitted(opts *bind.WatchOpts, sink chan<- *ControllerGeneDataSubmitted, docId []string, user []common.Address) (event.Subscription, error) {

	var docIdRule []interface{}
	for _, docIdItem := range docId {
		docIdRule = append(docIdRule, docIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "GeneDataSubmitted", docIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerGeneDataSubmitted)
				if err := _Controller.contract.UnpackLog(event, "GeneDataSubmitted", log); err != nil {
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

// ParseGeneDataSubmitted is a log parse operation binding the contract event 0xa44c9df778cfa6cc467846bb2c56756453c6c25c46977927881c4d88775d10e5.
//
// Solidity: event GeneDataSubmitted(string indexed docId, address indexed user, string hashContent)
func (_Controller *ControllerFilterer) ParseGeneDataSubmitted(log types.Log) (*ControllerGeneDataSubmitted, error) {
	event := new(ControllerGeneDataSubmitted)
	if err := _Controller.contract.UnpackLog(event, "GeneDataSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerGeneNFTMintedIterator is returned from FilterGeneNFTMinted and is used to iterate over the raw logs and unpacked data for GeneNFTMinted events raised by the Controller contract.
type ControllerGeneNFTMintedIterator struct {
	Event *ControllerGeneNFTMinted // Event containing the contract specifics and raw log

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
func (it *ControllerGeneNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerGeneNFTMinted)
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
		it.Event = new(ControllerGeneNFTMinted)
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
func (it *ControllerGeneNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerGeneNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerGeneNFTMinted represents a GeneNFTMinted event raised by the Controller contract.
type ControllerGeneNFTMinted struct {
	Owner   common.Address
	TokenId *big.Int
	DocId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneNFTMinted is a free log retrieval operation binding the contract event 0xca300ff523ecd04c68657ab684566141f6c8911842cf0cc1eb346135ffe1c556.
//
// Solidity: event GeneNFTMinted(address indexed owner, uint256 indexed tokenId, string docId)
func (_Controller *ControllerFilterer) FilterGeneNFTMinted(opts *bind.FilterOpts, owner []common.Address, tokenId []*big.Int) (*ControllerGeneNFTMintedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "GeneNFTMinted", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerGeneNFTMintedIterator{contract: _Controller.contract, event: "GeneNFTMinted", logs: logs, sub: sub}, nil
}

// WatchGeneNFTMinted is a free log subscription operation binding the contract event 0xca300ff523ecd04c68657ab684566141f6c8911842cf0cc1eb346135ffe1c556.
//
// Solidity: event GeneNFTMinted(address indexed owner, uint256 indexed tokenId, string docId)
func (_Controller *ControllerFilterer) WatchGeneNFTMinted(opts *bind.WatchOpts, sink chan<- *ControllerGeneNFTMinted, owner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "GeneNFTMinted", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerGeneNFTMinted)
				if err := _Controller.contract.UnpackLog(event, "GeneNFTMinted", log); err != nil {
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

// ParseGeneNFTMinted is a log parse operation binding the contract event 0xca300ff523ecd04c68657ab684566141f6c8911842cf0cc1eb346135ffe1c556.
//
// Solidity: event GeneNFTMinted(address indexed owner, uint256 indexed tokenId, string docId)
func (_Controller *ControllerFilterer) ParseGeneNFTMinted(log types.Log) (*ControllerGeneNFTMinted, error) {
	event := new(ControllerGeneNFTMinted)
	if err := _Controller.contract.UnpackLog(event, "GeneNFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerPCSPRewardedIterator is returned from FilterPCSPRewarded and is used to iterate over the raw logs and unpacked data for PCSPRewarded events raised by the Controller contract.
type ControllerPCSPRewardedIterator struct {
	Event *ControllerPCSPRewarded // Event containing the contract specifics and raw log

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
func (it *ControllerPCSPRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerPCSPRewarded)
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
		it.Event = new(ControllerPCSPRewarded)
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
func (it *ControllerPCSPRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerPCSPRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerPCSPRewarded represents a PCSPRewarded event raised by the Controller contract.
type ControllerPCSPRewarded struct {
	User      common.Address
	Amount    *big.Int
	RiskScore *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPCSPRewarded is a free log retrieval operation binding the contract event 0x3df4c48fd3dd34cbad69335f4fcb5d270a4b874251179300e51a6261d845d790.
//
// Solidity: event PCSPRewarded(address indexed user, uint256 amount, uint256 riskScore)
func (_Controller *ControllerFilterer) FilterPCSPRewarded(opts *bind.FilterOpts, user []common.Address) (*ControllerPCSPRewardedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "PCSPRewarded", userRule)
	if err != nil {
		return nil, err
	}
	return &ControllerPCSPRewardedIterator{contract: _Controller.contract, event: "PCSPRewarded", logs: logs, sub: sub}, nil
}

// WatchPCSPRewarded is a free log subscription operation binding the contract event 0x3df4c48fd3dd34cbad69335f4fcb5d270a4b874251179300e51a6261d845d790.
//
// Solidity: event PCSPRewarded(address indexed user, uint256 amount, uint256 riskScore)
func (_Controller *ControllerFilterer) WatchPCSPRewarded(opts *bind.WatchOpts, sink chan<- *ControllerPCSPRewarded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "PCSPRewarded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerPCSPRewarded)
				if err := _Controller.contract.UnpackLog(event, "PCSPRewarded", log); err != nil {
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

// ParsePCSPRewarded is a log parse operation binding the contract event 0x3df4c48fd3dd34cbad69335f4fcb5d270a4b874251179300e51a6261d845d790.
//
// Solidity: event PCSPRewarded(address indexed user, uint256 amount, uint256 riskScore)
func (_Controller *ControllerFilterer) ParsePCSPRewarded(log types.Log) (*ControllerPCSPRewarded, error) {
	event := new(ControllerPCSPRewarded)
	if err := _Controller.contract.UnpackLog(event, "PCSPRewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerUploadDataIterator is returned from FilterUploadData and is used to iterate over the raw logs and unpacked data for UploadData events raised by the Controller contract.
type ControllerUploadDataIterator struct {
	Event *ControllerUploadData // Event containing the contract specifics and raw log

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
func (it *ControllerUploadDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerUploadData)
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
		it.Event = new(ControllerUploadData)
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
func (it *ControllerUploadDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerUploadDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerUploadData represents a UploadData event raised by the Controller contract.
type ControllerUploadData struct {
	DocId     string
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUploadData is a free log retrieval operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) FilterUploadData(opts *bind.FilterOpts) (*ControllerUploadDataIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return &ControllerUploadDataIterator{contract: _Controller.contract, event: "UploadData", logs: logs, sub: sub}, nil
}

// WatchUploadData is a free log subscription operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) WatchUploadData(opts *bind.WatchOpts, sink chan<- *ControllerUploadData) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerUploadData)
				if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
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

// ParseUploadData is a log parse operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) ParseUploadData(log types.Log) (*ControllerUploadData, error) {
	event := new(ControllerUploadData)
	if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
