package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)

type TaskManagementChainCode struct{}

const (
	EVENT_CREATE   = "event-create-task"
	EVENT_APPROVE  = "event-approve-task"
	EVENT_FINISHED = "event-task-finished"
)

type Task struct {
	Name        string   `json:"name"`
	Creator     string   `json:"creator"`
	Requires    []string `json:"requires"`
	Approved    []string `json:"approved"`
	Description string   `json:"description"`
}

func (t *Task) FromJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return fmt.Errorf("unmarshal json task failed: %s", err)
	}
	return nil
}

func (t *Task) ToJson() ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("marshal task to json failed: %s", err)
	}
	return data, nil
}

func (t *Task) IsFinished() bool {
	for _, user := range t.Requires {
		found := false
		for _, sig := range t.Approved {
			if user == sig {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (t *Task) Approve(userIdentity string) {
	for _, sig := range t.Approved {
		if sig == userIdentity {
			return
		}
	}
	t.Approved = append(t.Approved, userIdentity)
}

func (cc *TaskManagementChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (cc *TaskManagementChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fnc, args := stub.GetFunctionAndParameters()
	switch fnc {
	case "get":
		return cc.get(stub, args)
	case "create":
		return cc.create(stub, args)
	case "approve":
		return cc.approve(stub, args)
	default:
		return shim.Error("Invalid function name, support 'get', 'create', 'approve'")
	}
}

func getUserIdentity(stub shim.ChaincodeStubInterface) (string, error) {
	user, err := cid.New(stub)
	if err != nil {
		return "", fmt.Errorf("get invoke user identity failed: %s", err)
	}
	userName, _, err := user.GetAttributeValue("hf.EnrollmentID")
	if err != nil {
		return "", fmt.Errorf("get invoke user name failed: %s", err)
	}
	mspId, err := user.GetMSPID()
	if err != nil {
		return "", fmt.Errorf("get invoke user msp id failed: %s", err)
	}
	return fmt.Sprintf("%s.%s", mspId, userName), nil
}

func (cc *TaskManagementChainCode) getTaskKey(stub shim.ChaincodeStubInterface, name string) (string, error) {
	userIdentity, err := getUserIdentity(stub)
	if err != nil {
		return "", err
	}
	key, err := stub.CreateCompositeKey(userIdentity, []string{name})
	if err != nil {
		return "", fmt.Errorf("create composite key failed: %s", err)
	}
	return key, nil
}

func (cc *TaskManagementChainCode) getTask(stub shim.ChaincodeStubInterface, name string) (*Task, error) {
	key, err := cc.getTaskKey(stub, name)
	if err != nil {
		return nil, err
	}
	value, err := stub.GetState(key)
	if err != nil {
		return nil, fmt.Errorf("get task %s from ledger failed: %s", name, err)
	}
	if value == nil {
		return nil, fmt.Errorf("task %s not found in ledger", name)
	}
	task := &Task{}
	err = task.FromJson(value)
	if err != nil {
		return nil, err
	}
	task.Name = name
	return task, nil
}

func (cc *TaskManagementChainCode) updateTask(stub shim.ChaincodeStubInterface, task *Task) ([]byte, error) {
	key, err := cc.getTaskKey(stub, task.Name)
	if err != nil {
		return nil, err
	}
	data, err := task.ToJson()
	if err != nil {
		return nil, nil
	}
	err = stub.PutState(key, data)
	if err != nil {
		return nil, fmt.Errorf("update task %s in ledger failed: %s", task.Name, err)
	}
	return data, nil
}

func (cc *TaskManagementChainCode) create(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Invalid argument, require <name> <json task>")
	}
	name := args[0]
	taskJson := args[1]
	task := &Task{}
	err := task.FromJson([]byte(taskJson))
	if err != nil {
		msg := fmt.Sprintf("Invalid json task: %s", err)
		logger.Println(msg)
		return shim.Error(msg)
	}
	task.Name = name
	userIdentity, err := getUserIdentity(stub)
	if err != nil {
		msg := fmt.Sprintf("Get user identity failed: %s", err)
		logger.Println(msg)
		return shim.Error(msg)
	}
	task.Creator = userIdentity

	_, err = cc.getTask(stub, name)
	if err == nil {
		msg := fmt.Sprintf("Task %s already exists in ledger", name)
		logger.Printf(msg)
		return shim.Error(msg)
	}

	data, err := cc.updateTask(stub, task)
	if err != nil {
		msg := fmt.Sprintf("Create task %s failed: %s", name, err)
		logger.Printf(msg)
		return shim.Error(msg)
	}

	err = stub.SetEvent(EVENT_CREATE, data)
	if err != nil {
		msg := fmt.Sprintf("Set task create event failed: %s", err)
		logger.Printf(msg)
		return shim.Error(msg)
	}
	logger.Printf("Task %s created\n", name)

	return shim.Success([]byte(name))
}

func (cc *TaskManagementChainCode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument, require <name>")
	}
	name := args[0]
	task, err := cc.getTask(stub, name)
	if err != nil {
		msg := fmt.Sprintf("Get task %s failed: %s", name, err)
		logger.Printf(msg)
		return shim.Error(msg)
	}

	data, err := task.ToJson()
	if err != nil {
		logger.Printf(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(data)
}

func (cc *TaskManagementChainCode) approve(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument, require <name>")
	}
	name := args[0]
	task, err := cc.getTask(stub, name)
	if err != nil {
		msg := fmt.Sprintf("Get task %s failed: %s", name, err)
		logger.Printf(msg)
		return shim.Error(msg)
	}
	userIdentity, err := getUserIdentity(stub)
	if err != nil {
		msg := fmt.Sprintf("Get user identity failed: %s", err)
		logger.Printf(msg)
		return shim.Error(msg)
	}
	task.Approve(userIdentity)
	data, err := cc.updateTask(stub, task)
	if err != nil {
		msg := fmt.Sprintf("Update task %s failed: %s", name, err)
		logger.Printf(msg)
		return shim.Error(msg)
	}

	eventName := EVENT_APPROVE
	if task.IsFinished() {
		eventName = EVENT_FINISHED
	}
	err = stub.SetEvent(eventName, data)
	if err != nil {
		msg := fmt.Sprintf("Set task event failed: %s", err)
		logger.Printf(msg)
		return shim.Error(msg)
	}

	logger.Printf("User %s approved taks %s\n", userIdentity, name)
	return shim.Success(data)
}

func main() {
	err := shim.Start(new(TaskManagementChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
