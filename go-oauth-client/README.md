# Golang REST-API Client 示例

本示例为您展示如何通过 Golang 访问阿里云区块链 REST-API。示例主要演示了如下操作：
1. 获取指定通道区块号为1的区块
2. 调用通道中的 Fabric 智能合约
3. 调用通道中的 Fabric 智能合约的查询接口

## 使用方法

1. 配置本地 Golang 1.13.x 环境

2. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [notary](./contracts/fabric/notary) 部署到通道中

3. 根据您组织实例的 REST-API 地址、生成的 Token 信息、以及业务通道和链码信息，修改文件 `src/go-oauth-client/main.go` 中的参数

4. 进入目录 `src/go-oauth-client` 执行 `go run main.go` 运行示例程序

**正常结果示例**

```
> go run main.go
Block response body: {"Success":true,"Result":{"number":1,"hash":"1c397c4eb3e0e330c01ec430170f844e46159f16930aa347486b8153b6586548","create_time":1579056958,"previous_hash":"88ef0ad6ba2df7ba7e53de78575d2d14cee2253fe6897305e50b57ceeecebc78","transactions":[],"data":{"data":{"data":[{"payload":{"data":{"config":{"channel_group": ...}}}}]}}}
Invoke response body: {"Success":true,"Result":{"id":"e0a11c3b953fa1759ef715214bb8bd24c0a7e762b739eb5f645ead921314fef4","status":"200","events":[],"data":"MTU4MTkzNDAwOA=="},"Error":{"code":200,"message":"Success","request_id":"c3c92ad0-a0ef-4a82-b596-4ac50a893ef6"}}
Invoke contract response: "MTU4MTkzNDAwOA=="
Query response body: {"Success":true,"Result":{"id":"c3d4c2928dfa7129641863b18dfeee4b18c7227929c841ff7d8bd25148f6c5f6","status":"200","events":[],"data":"MTU4MTkzNDAwOA=="},"Error":{"code":200,"message":"Success","request_id":"f1bf52b2-8c60-491e-9a40-8fecb96667ea"}}
Query contract response: "MTU4MTkzNDAwOA=="
```