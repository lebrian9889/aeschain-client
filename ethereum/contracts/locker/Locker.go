// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package locker

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

// LockerMetaData contains all meta data concerning the Locker contract.
var LockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"RequiredApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"UnAuthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nounce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"LockNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nounce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"ReleaseNFT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LockerABI is the input ABI used to generate the binding from.
// Deprecated: Use LockerMetaData.ABI instead.
var LockerABI = LockerMetaData.ABI

// Locker is an auto generated Go binding around an Ethereum contract.
type Locker struct {
	LockerCaller     // Read-only binding to the contract
	LockerTransactor // Write-only binding to the contract
	LockerFilterer   // Log filterer for contract events
}

// LockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockerSession struct {
	Contract     *Locker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockerCallerSession struct {
	Contract *LockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockerTransactorSession struct {
	Contract     *LockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockerRaw struct {
	Contract *Locker // Generic contract binding to access the raw methods on
}

// LockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockerCallerRaw struct {
	Contract *LockerCaller // Generic read-only contract binding to access the raw methods on
}

// LockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockerTransactorRaw struct {
	Contract *LockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocker creates a new instance of Locker, bound to a specific deployed contract.
func NewLocker(address common.Address, backend bind.ContractBackend) (*Locker, error) {
	contract, err := bindLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locker{LockerCaller: LockerCaller{contract: contract}, LockerTransactor: LockerTransactor{contract: contract}, LockerFilterer: LockerFilterer{contract: contract}}, nil
}

// NewLockerCaller creates a new read-only instance of Locker, bound to a specific deployed contract.
func NewLockerCaller(address common.Address, caller bind.ContractCaller) (*LockerCaller, error) {
	contract, err := bindLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockerCaller{contract: contract}, nil
}

// NewLockerTransactor creates a new write-only instance of Locker, bound to a specific deployed contract.
func NewLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*LockerTransactor, error) {
	contract, err := bindLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockerTransactor{contract: contract}, nil
}

// NewLockerFilterer creates a new log filterer instance of Locker, bound to a specific deployed contract.
func NewLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*LockerFilterer, error) {
	contract, err := bindLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockerFilterer{contract: contract}, nil
}

// bindLocker binds a generic wrapper to an already deployed contract.
func bindLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.LockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transact(opts, method, params...)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address _nftAddress, uint256 _tokenId, string _data) returns()
func (_Locker *LockerTransactor) Lock(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _data string) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "lock", _nftAddress, _tokenId, _data)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address _nftAddress, uint256 _tokenId, string _data) returns()
func (_Locker *LockerSession) Lock(_nftAddress common.Address, _tokenId *big.Int, _data string) (*types.Transaction, error) {
	return _Locker.Contract.Lock(&_Locker.TransactOpts, _nftAddress, _tokenId, _data)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address _nftAddress, uint256 _tokenId, string _data) returns()
func (_Locker *LockerTransactorSession) Lock(_nftAddress common.Address, _tokenId *big.Int, _data string) (*types.Transaction, error) {
	return _Locker.Contract.Lock(&_Locker.TransactOpts, _nftAddress, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Locker *LockerTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Locker *LockerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Locker.Contract.OnERC721Received(&_Locker.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Locker *LockerTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Locker.Contract.OnERC721Received(&_Locker.TransactOpts, operator, from, tokenId, data)
}

// Release is a paid mutator transaction binding the contract method 0xd1c89de8.
//
// Solidity: function release(address _nftAddress, uint256 _tokenId, address _receiver, string _data) returns()
func (_Locker *LockerTransactor) Release(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _receiver common.Address, _data string) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "release", _nftAddress, _tokenId, _receiver, _data)
}

// Release is a paid mutator transaction binding the contract method 0xd1c89de8.
//
// Solidity: function release(address _nftAddress, uint256 _tokenId, address _receiver, string _data) returns()
func (_Locker *LockerSession) Release(_nftAddress common.Address, _tokenId *big.Int, _receiver common.Address, _data string) (*types.Transaction, error) {
	return _Locker.Contract.Release(&_Locker.TransactOpts, _nftAddress, _tokenId, _receiver, _data)
}

// Release is a paid mutator transaction binding the contract method 0xd1c89de8.
//
// Solidity: function release(address _nftAddress, uint256 _tokenId, address _receiver, string _data) returns()
func (_Locker *LockerTransactorSession) Release(_nftAddress common.Address, _tokenId *big.Int, _receiver common.Address, _data string) (*types.Transaction, error) {
	return _Locker.Contract.Release(&_Locker.TransactOpts, _nftAddress, _tokenId, _receiver, _data)
}

// LockerLockNFTIterator is returned from FilterLockNFT and is used to iterate over the raw logs and unpacked data for LockNFT events raised by the Locker contract.
type LockerLockNFTIterator struct {
	Event *LockerLockNFT // Event containing the contract specifics and raw log

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
func (it *LockerLockNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerLockNFT)
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
		it.Event = new(LockerLockNFT)
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
func (it *LockerLockNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerLockNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerLockNFT represents a LockNFT event raised by the Locker contract.
type LockerLockNFT struct {
	NftAddress common.Address
	From       common.Address
	TokenId    *big.Int
	Nounce     *big.Int
	Data       string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockNFT is a free log retrieval operation binding the contract event 0x4fee5ebb8064020f4f38f427e9d749e3a49779a1805930493ac35ca840f5edd6.
//
// Solidity: event LockNFT(address _nftAddress, address _from, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) FilterLockNFT(opts *bind.FilterOpts) (*LockerLockNFTIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "LockNFT")
	if err != nil {
		return nil, err
	}
	return &LockerLockNFTIterator{contract: _Locker.contract, event: "LockNFT", logs: logs, sub: sub}, nil
}

// WatchLockNFT is a free log subscription operation binding the contract event 0x4fee5ebb8064020f4f38f427e9d749e3a49779a1805930493ac35ca840f5edd6.
//
// Solidity: event LockNFT(address _nftAddress, address _from, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) WatchLockNFT(opts *bind.WatchOpts, sink chan<- *LockerLockNFT) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "LockNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerLockNFT)
				if err := _Locker.contract.UnpackLog(event, "LockNFT", log); err != nil {
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

// ParseLockNFT is a log parse operation binding the contract event 0x4fee5ebb8064020f4f38f427e9d749e3a49779a1805930493ac35ca840f5edd6.
//
// Solidity: event LockNFT(address _nftAddress, address _from, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) ParseLockNFT(log types.Log) (*LockerLockNFT, error) {
	event := new(LockerLockNFT)
	if err := _Locker.contract.UnpackLog(event, "LockNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerReleaseNFTIterator is returned from FilterReleaseNFT and is used to iterate over the raw logs and unpacked data for ReleaseNFT events raised by the Locker contract.
type LockerReleaseNFTIterator struct {
	Event *LockerReleaseNFT // Event containing the contract specifics and raw log

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
func (it *LockerReleaseNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerReleaseNFT)
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
		it.Event = new(LockerReleaseNFT)
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
func (it *LockerReleaseNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerReleaseNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerReleaseNFT represents a ReleaseNFT event raised by the Locker contract.
type LockerReleaseNFT struct {
	NftAddress common.Address
	To         common.Address
	TokenId    *big.Int
	Nounce     *big.Int
	Data       string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReleaseNFT is a free log retrieval operation binding the contract event 0x70bd022d34c1d3f31f5de67a2f3f26a19fd99a9f868ee153dabe6cd7fb7e50ec.
//
// Solidity: event ReleaseNFT(address _nftAddress, address _to, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) FilterReleaseNFT(opts *bind.FilterOpts) (*LockerReleaseNFTIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "ReleaseNFT")
	if err != nil {
		return nil, err
	}
	return &LockerReleaseNFTIterator{contract: _Locker.contract, event: "ReleaseNFT", logs: logs, sub: sub}, nil
}

// WatchReleaseNFT is a free log subscription operation binding the contract event 0x70bd022d34c1d3f31f5de67a2f3f26a19fd99a9f868ee153dabe6cd7fb7e50ec.
//
// Solidity: event ReleaseNFT(address _nftAddress, address _to, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) WatchReleaseNFT(opts *bind.WatchOpts, sink chan<- *LockerReleaseNFT) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "ReleaseNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerReleaseNFT)
				if err := _Locker.contract.UnpackLog(event, "ReleaseNFT", log); err != nil {
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

// ParseReleaseNFT is a log parse operation binding the contract event 0x70bd022d34c1d3f31f5de67a2f3f26a19fd99a9f868ee153dabe6cd7fb7e50ec.
//
// Solidity: event ReleaseNFT(address _nftAddress, address _to, uint256 _tokenId, uint256 _nounce, string _data)
func (_Locker *LockerFilterer) ParseReleaseNFT(log types.Log) (*LockerReleaseNFT, error) {
	event := new(LockerReleaseNFT)
	if err := _Locker.contract.UnpackLog(event, "ReleaseNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
