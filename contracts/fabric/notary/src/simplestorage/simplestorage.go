package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/hyperledger/fabric-protos-go/peer"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)

type SimpleStorageChainCode struct{}

func (s *SimpleStorageChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SimpleStorageChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fnc, args := stub.GetFunctionAndParameters()
	switch fnc {
	case "get":
		return s.get(stub, args)
	case "put":
		return s.put(stub, args)
	case "set":
		return s.set(stub, args)
	case "history":
		return s.history(stub, args)
	default:
		return shim.Error("Invalid function name, support 'get', 'put', 'set', 'history'")
	}
}

func (s *SimpleStorageChainCode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument, require <key>")
	}
	key := args[0]
	value, err := stub.GetState(key)
	if err != nil {
		logger.Printf("Get key %s from ledger failed: %s\n", key, err)
		return shim.Error(fmt.Sprintf("GetState error %v", err))
	}
	logger.Printf("Got key %s from ledger\n", key)
	return shim.Success(value)
}

func (s *SimpleStorageChainCode) put(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Invalid argument, require <key> and <value>")
	}
	key := args[0]
	value := args[1]
	val, err := stub.GetState(key)
	if err != nil {
		logger.Printf("Check key %s whether exists in ledger failed: %s\n", key, err)
		return shim.Error(fmt.Sprintf("GetState error %v", err))
	}
	if val != nil {
		logger.Printf("Put key %s failed: already exists\n", key)
		return shim.Error(fmt.Sprintf("Key %s already exists", key))
	}
	err = stub.PutState(key, []byte(value))
	if err != nil {
		logger.Printf("Put key %s failed: %s\n", key, err)
		return shim.Error(fmt.Sprintf("PutState error %v", err))
	}
	logger.Printf("Put key %s into ledger\n", key)
	return shim.Success([]byte(key))
}

func (s *SimpleStorageChainCode) set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Invalid argument, require <key> and <value>")
	}
	key := args[0]
	value := args[1]
	err := stub.PutState(key, []byte(value))
	if err != nil {
		logger.Printf("Set key %s failed: %s\n", key, err)
		return shim.Error(fmt.Sprintf("PutState error %v", err))
	}
	logger.Printf("Set key %s in ledger\n", key)
	return shim.Success([]byte(key))
}

func (s *SimpleStorageChainCode) history(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument, require <key>")
	}
	key := args[0]
	iter, err := stub.GetHistoryForKey(key)
	if err != nil {
		logger.Printf("Get history for key %s failed: %s\n", key, err)
		return shim.Error(fmt.Sprintf("GetHistoryForKey error %v", err))
	}
	var result []*queryresult.KeyModification
	for iter.HasNext() {
		item, err := iter.Next()
		if err != nil {
			logger.Printf("Iter to next hisotry failed: %s\n", err)
			return shim.Error(fmt.Sprintf("Iter to next error %v", err))
		}
		result = append(result, item)
	}
	data, err := json.Marshal(result)
	if err != nil {
		logger.Printf("Marshl history result to json failed: %s\n", err)
		return shim.Error(fmt.Sprintf("Marshl history result error %v", err))
	}
	logger.Printf("Got key %s hisotry in ledger\n", key)
	return shim.Success(data)
}

func main() {
	err := shim.Start(new(SimpleStorageChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
