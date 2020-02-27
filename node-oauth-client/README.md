# Nodejs REST-API Client 示例


本示例为您展示如何通过 Nodejs 访问阿里云区块链 REST-API。示例主要演示了如下操作：
1. 获取指定通道区块号为1的区块
2. 调用通道中的 Fabric 智能合约
3. 调用通道中的 Fabric 智能合约的查询接口

## 使用方法

1. 配置本地 node v8 环境(>=8.17.0)

2. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [notary](./contracts/fabric/notary) 部署到通道中

3. 根据您组织实例的 REST-API 地址、生成的 Token 信息、以及业务通道和链码信息，修改文件 `main.js` 中的参数

4. 执行 `npm install` 安装依赖包，再执行 `node main.js` 运行示例程序

**正常结果示例**
```
> node main.js
{ number: 1,
  hash: '1c397c4eb3e0e330c01ec430170f844e46159f16930aa347486b8153b6586548',
  create_time: 1579056958,
  previous_hash: '88ef0ad6ba2df7ba7e53de78575d2d14cee2253fe6897305e50b57ceeecebc78',
  transactions: [],
  data:
   { data: { data: [Array] },
     header:
      { data_hash: 'HDl8TrPg4zDAHsQwFw+ETkYVnxaTCqNHSGuBU7ZYZUg=',
        number: '1',
        previous_hash: 'iO8K1rot97p+U954V10tFM7iJT/miXMF5QtXzu7OvHg=' },
     metadata: { metadata: [Array] } } }
Data 1581931486180 pushed to blockchain with transaction f19217c0db571dc715af8ad99025422f03e5561910371841fe5e69a356d0cb23
{ id: '8fd06f6087c5128b7dbe309658b170366b37e0733994e5e959d30be201c28827',
  status: '200',
  events: [],
  data: 'MTU4MTkzMTQ4NjE4MA==' }
```
