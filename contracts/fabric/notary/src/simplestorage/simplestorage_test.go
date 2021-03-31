package main

import (
	"github.com/golang/mock/gomock"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestFunc(f func(t *testing.T, stub *MockChaincodeStubInterface)) func(t *testing.T) {
	return func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		stub := NewMockChaincodeStubInterface(ctrl)
		f(t, stub)
	}
}

func TestSimpleStorageChainCode_Invoke_Get(t *testing.T) {
	cc := new(SimpleStorageChainCode)

	t.Run("Key exists", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("get", []string{"key"})
		stub.EXPECT().GetState("key").Times(1).Return([]byte("val"), nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.OK), resp.Status)
		assert.Equal(t, "val", string(resp.Payload))
	}))

	t.Run("Key not exists", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("get", []string{"key"})
		stub.EXPECT().GetState("key").Times(1).Return(nil, nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.OK), resp.Status)
		assert.Nil(t, resp.Payload)
	}))
}

func TestSimpleStorageChainCode_Invoke_Put(t *testing.T) {
	cc := new(SimpleStorageChainCode)

	t.Run("Key exists", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("put", []string{"key", "val"})
		stub.EXPECT().GetState("key").Times(1).Return([]byte("val"), nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.ERROR), resp.Status)
		assert.Contains(t, resp.Message, "exists")
	}))

	t.Run("Key not exists", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("put", []string{"key", "val"})
		stub.EXPECT().GetState("key").Times(1).Return(nil, nil)
		stub.EXPECT().PutState("key", []byte("val")).Times(1).Return(nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.OK), resp.Status)
		assert.Equal(t, "key", string(resp.Payload))
	}))
}

func TestSimpleStorageChainCode_Invoke_Set(t *testing.T) {
	cc := new(SimpleStorageChainCode)

	t.Run("Set key", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("set", []string{"key", "val"})
		stub.EXPECT().PutState("key", []byte("val")).Times(1).Return(nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.OK), resp.Status)
		assert.Equal(t, "key", string(resp.Payload))
	}))
}

func TestSimpleStorageChainCode_Invoke_History(t *testing.T) {
	cc := new(SimpleStorageChainCode)

	t.Run("Get key history", newTestFunc(func(t *testing.T, stub *MockChaincodeStubInterface) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		iter := NewMockHistoryQueryIteratorInterface(ctrl)

		stub.EXPECT().GetFunctionAndParameters().Times(1).Return("history", []string{"key"})
		stub.EXPECT().GetHistoryForKey("key").Times(1).Return(iter, nil)
		iter.EXPECT().HasNext().Times(1).Return(true)
		iter.EXPECT().HasNext().Times(1).Return(false)
		iter.EXPECT().Next().Return(&queryresult.KeyModification{TxId: "txId", Value: []byte("history")}, nil)
		resp := cc.Invoke(stub)
		assert.Equal(t, int32(shim.OK), resp.Status)
		assert.Equal(t, `[{"tx_id":"txId","value":"aGlzdG9yeQ=="}]`, string(resp.Payload))
	}))
}
