# 事件过滤推送

在配置云服务集成触发器时，您可以为触发器配置过滤器，过滤器遵循 jq 工具的基本语法，您可以通过表达式对数据进行过滤和简单的处理。

下面介绍如何使用过滤器对区块链事件进行过滤和简单的预处理，本示例通过函数计算触发器来演示功能，您需要按如下步骤准备环境：

1. 安装函数计算工具 funcraft, [funcraft 介绍及安装方式](https://help.aliyun.com/document_detail/140283.html)
2. 使用 `fun deploy` 将函数部署到函数计算。默认服务名为"octopus"，函数名为"LoggerFunc"。更多配置方式可以参考 [funcraft文档](https://github.com/alibaba/funcraft/blob/master/docs/specs/2018-04-03-zh-cn.md)
3. 进入[函数计算控制台](https://fc.console.aliyun.com/fc/)，为我们刚才创建的函数 `LoggerFunc` 开通日志查询
  ![开通日志查询](http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/pic/151966/cn_zh/1582169498182/1E7BDA41-7698-406E-A88B-B186B5D81676.png)

4. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [taskmgr](./contracts/fabric/taskmgr) 部署到通道中
5. 配置云服务集成触发器
  1. 根据帮助，填写函数计算实例的相关信息
  2. 在高级选项中，根据以下不同的示例场景，填写事件过滤表达式
6. 进入 `../blockchain2sms` 目录，参考文档[使用REST API](https://help.aliyun.com/document_detail/151690.html)按照注释修改 main.js 中的配置参数，填入REST API地址、Refresh Token、通道名和智能合约名称
7. 执行 `npm install` 安装依赖包，通过 `node main.js` 发起示例交易
8. 查看函数计算的日志，观察收到的事件内容


## 示例场景

### 一、只推送满足特定条件事件

只推送智能合约事件名称为 event-create-task 的事件

- 事件类型选择： Contract
- filter 配置： `select(.name=="event-create-task")`
- 推送数据示例：

```
{
  "content": "eyJuYW1lIjoidGFzay0xNTgwOTcxODQyOTM3IiwiY3JlYXRvciI6ImUyZWhtZnFhc3RoTVNQLm9jdG9wdXNfMjY4NDJfMTIzNDU2Nzg5MDEyMzQiLCJyZXF1aXJlcyI6WyJlMmVobWZxYXN0aE1TUC5vY3RvcHVzXzI2ODQyXzEyMzQ1Njc4OTAxMjM0Il0sImFwcHJvdmVkIjpudWxsLCJkZXNjcmlwdGlvbiI6IuekuuS+i+S7u+WKoe+8jHJlcXVpcmVzIOmFjee9ruWuoeaJueS7u+WKoeWujOaIkOmcgOimgemCo+S6m+eUqOaIt+WQjOaEj+OAgueUqOaIt+aPj+i/sOS4uiAn57uE57uHTVNQLueUqOaIt+WQjeensCcifQ==",
  "id": "contract-131-ca144c3385f56a429b9e8874173f7c97fbb49de69f931aec583ec50222a3a2ed",
  "instance_id": "csi-e2ehmfqasth-bcw7tzao2dzeo",
  "name": "event-create-task",
  "network": "channel3",
  "platform": "Fabric",
  "type": 8
}
```

### 二、只推送合法的特定交易事件

只推送发送给智能合约 taskmgr 的交易，且交易状态为合法

- 事件类型选择： Tx
- filter 配置：`select(.content.to=="taskmgr" and .content.state=="VALID")`
- 推送数据示例：

```
{
  "content": {
    "events": [
      "event-create-task"
    ],
    "from": "e2ehmfqasthMSP.octopus_26842_12345678901234",
    "id": "3a2e0e375c6ee0d9c6c20d756e787db44d50bfeba51f22d8f262dd1556dedba0",
    "input": "[\"create\",\"task-1580971842937\",\"{\\n \\\"requires\\\": [\\\"e2ehmfqasthMSP.octopus_26842_12345678901234\\\"],\\n \\\"description\\\": \\\"示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'\\\"\\n }\"]",
    "state": "VALID",
    "to": "taskmgr"
  },
  "id": "tx-128-3a2e0e375c6ee0d9c6c20d756e787db44d50bfeba51f22d8f262dd1556dedba0",
  "instance_id": "csi-e2ehmfqasth-bcw7tzao2dzeo",
  "name": "3a2e0e375c6ee0d9c6c20d756e787db44d50bfeba51f22d8f262dd1556dedba0",
  "network": "channel3",
  "platform": "Fabric",
  "type": 2
}
```

### 三、只推送事件的部分内容

只推送智能合约事件的内容，且事件的名称为 event-create-task

- 事件类型选择： Contract
- filter 配置：`select(.name=="event-create-task") | .content`
- 推送数据示例：

```
"eyJuYW1lIjoidGFzay0xNTgwOTcxODQyOTM3IiwiY3JlYXRvciI6ImUyZWhtZnFhc3RoTVNQLm9jdG9wdXNfMjY4NDJfMTIzNDU2Nzg5MDEyMzQiLCJyZXF1aXJlcyI6WyJlMmVobWZxYXN0aE1TUC5vY3RvcHVzXzI2ODQyXzEyMzQ1Njc4OTAxMjM0Il0sImFwcHJvdmVkIjpudWxsLCJkZXNjcmlwdGlvbiI6IuekuuS+i+S7u+WKoe+8jHJlcXVpcmVzIOmFjee9ruWuoeaJueS7u+WKoeWujOaIkOmcgOimgemCo+S6m+eUqOaIt+WQjOaEj+OAgueUqOaIt+aPj+i/sOS4uiAn57uE57uHTVNQLueUqOaIt+WQjeensCcifQ=="
```

### 四、获取区块链的写入情况

只推送智能合约 taskmgr 中所有 Key 的写入情况，且状态为合法的交易

- 事件类型选择： Tx
- filter 配置：`select(.content.state=="VALID" and .content.to=="taskmgr") | .content.data.data.actions.[0].payload.action.proposal_response_payload.extension.results.ns_rwset[] | select(.namespace=="taskmgr") | .rwset.writes`
- 推送数据示例：

```
[
  {
    "is_delete": false,
    "key": "\u0000e2ehmfqasthMSP.octopus_26842_12345678901234\u0000task-1581479306267\u0000",
    "value": "eyJuYW1lIjoidGFzay0xNTgwOTcxODQyOTM3IiwiY3JlYXRvciI6ImUyZWhtZnFhc3RoTVNQLm9jdG9wdXNfMjY4NDJfMTIzNDU2Nzg5MDEyMzQiLCJyZXF1aXJlcyI6WyJlMmVobWZxYXN0aE1TUC5vY3RvcHVzXzI2ODQyXzEyMzQ1Njc4OTAxMjM0Il0sImFwcHJvdmVkIjpudWxsLCJkZXNjcmlwdGlvbiI6IuekuuS+i+S7u+WKoe+8jHJlcXVpcmVzIOmFjee9ruWuoeaJueS7u+WKoeWujOaIkOmcgOimgemCo+S6m+eUqOaIt+WQjOaEj+OAgueUqOaIt+aPj+i/sOS4uiAn57uE57uHTVNQLueUqOaIt+WQjeensCcifQ=="
  }
]
```

### 五、对特定Key进行监控

监控智能合约 taskmgr 中特定 Key 的 Value 变化情况，且状态为合法的交易

- 事件类型选择： Tx
- filter 配置：`select(.content.state=="VALID" and .content.to=="taskmgr") | .content.data.data.actions.[0].payload.action.proposal_response_payload.extension.results.ns_rwset[] | select(.namespace=="taskmgr") | .rwset.writes[] | select(.key=="special_key")`
- 推送数据示例：

```
{
  "is_delete": false,
  "key": "special_key",
  "value": "MTA="
}
```



