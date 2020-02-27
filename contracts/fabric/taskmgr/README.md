# Fabric 示例任务流合约

简易审批合约提供了基于区块链的简单任务审批接口，支持创建、查询和审批任务，当任务的状态发生变更时，会产生相应的区块链事件。您可以通过`src`目录下的源码自行打包成链码包使用，也可以直接使用我们已经打包好的链码`taskmgr.x.x.cc`。

## 初始化参数

无初始化参数

## 数据结构

#### task

```
type Task struct {
    // 任务名称
	Name  		string `json:"name"`
    // 任务的创建者
	Creator 	string `json:"creator"`
    // 任务需要那些用户同意后，流程才能结束（用户表示方式： ${MSP_ID}.${USER_NAME}）
	Requires    []string `json:"requires"`
    // 当前已经同意该任务的用户（用户表示方式： ${MSP_ID}.${USER_NAME}）
	Approved  	[]string `json:"approved"`
    // 任务的描述
	Description string `json:"description"`
}
```

## 接口

#### create

创建新的任务流

- 参数列表
  1. name: 任务的名称
  2. task: 经过 Json 编码的任务内容，需要具有 requires、description 两个字段
- 返回值： 
  - 成功时，返回成功创建的任务名称
  - 失败时，返回错误原因
- 事件：
  - 触发事件 `event-create-task`，事件内容为经过 Json 编码的任务内容

#### get

获取任务流

- 参数列表
  1. name: 任务的名称
- 返回值： 
  - 成功时，返回经过 Json 编码的任务内容
  - 失败时，返回错误原因

#### approve

同意某个任务

- 参数列表
  1. name: 任务的名称
- 返回值： 
  - 成功时，返回经过 Json 编码的任务内容
  - 失败时，返回错误原因
- 事件：
  - 同意任务后，任务满足结束条件时，触发事件 `event-task-finished`，事件内容为经过 Json 编码的任务内容
  - 同意任务后，任务未满足结束条件时，触发事件 `event-approve-task`，事件内容为经过 Json 编码的任务内容
