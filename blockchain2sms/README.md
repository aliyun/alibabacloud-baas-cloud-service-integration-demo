# 链上事件触发短信通知

本示例为您展示如何利用阿里云区块链云服务集成功能，监控链上发生的事件；当区块链上发生特定的事件时，能自动通过短信的方式通知到相应用户。

## 前置条件
1. 您需要先在区块链实例中[安装云服务集成](https://help.aliyun.com/document_detail/153337.html), 并开通函数计算集成功能
2. 本地安装 nodev8 环境(>=8.17.0)
3. 安装函数计算工具 funcraft, [funcraft 介绍及安装方式](https://help.aliyun.com/document_detail/140283.html)

## 操作步骤

1. 进入目录 `fc`, 修改以下内容
  1. 按照注释修改 index.js 中的配置参数，填入具有调用短信服务权限的Access Key和 Access Key Secret
  2. 根据您短信服务的配置，修改短信签名(SignName)、模版CODE(TemplateCode)和模版参数(TemplateParam)
  3. 修改 template.yml 中的函数计算服务名称和函数名称，默认服务名为"octopus"，函数名为"BlockChain2SMS"。更多配置方式可以参考 [funcraft文档](https://github.com/alibaba/funcraft/blob/master/docs/specs/2018-04-03-zh-cn.md)
2. 使用 `fun deploy` 将函数部署到函数计算
3. 在控制台创建云服务集成函数计算触发器
  1. 配置触发器的事件类型为 Contact
  2. 根据帮助，填写函数计算实例的相关信息

4. 按照[发起示例交易](#send-demo-tx)的步骤，在通道上发起示例交易
5. 当审批流满足条件结束时，您会收到短信通知。我们也可以在[短信服务控制台](https://dysms.console.aliyun.com/dysms.htm#/statistic/record)看到如下发送记录

![短信集成](http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/pic/151956/cn_zh/1581922484192/09981890-B46C-4E62-9029-8B4F5AFC142A.png)


## 发起示例交易
<span id="send-demo-tx"></span>

1. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [taskmgr](./contracts/fabric/taskmgr) 部署到通道中
2.  进入 `blockchain2sms` 目录，参考文档[使用REST API](https://help.aliyun.com/document_detail/151690.html)按照注释修改 main.js 中的配置参数，填入REST API地址、Refresh Token、通道名和智能合约名称
3. 执行 `npm install` 安装依赖包，通过 `node main.js` 发起示例交易

**成功示例输出**
```
Data pushed to blockchain with transaction 701c7006f26aed8457273a00bbfcc8cea4d75eac958996e07837036ea7e2fdac
{ id: '701c7006f26aed8457273a00bbfcc8cea4d75eac958996e07837036ea7e2fdac',
  status: '200',
  events:
   [ { type: 'Contract',
       platform: 'Fabric',
       instance_id: 'csi-e2ehmfqasth-bcw7tzao2dzeo',
       network: '',
       id: '',
       name: 'event-create-task',
       content: 'eyJuY...TAz' }
Data pushed to blockchain with transaction 60a06a189415db587b49cbf91b46467bce1ea16490b19f6dfc8d520aa31240bc
{ id: '60a06a189415db587b49cbf91b46467bce1ea16490b19f6dfc8d520aa31240bc',
  status: '200',
  events:
   [ { type: 'Contract',
       platform: 'Fabric',
       instance_id: 'csi-e2ehmfqasth-bcw7tzao2dzeo',
       network: '',
       id: '',
       name: 'event-task-finished',
       content: 'eyJuY...p7In0=' } ],
  data: 'eyJu...In0=' }
Data pushed to blockchain with transaction 7810e0496a1c91a16102736b00a6f26da0baa42874a907fff26bff1b7eb3bf27
{ id: '7810e0496a1c91a16102736b00a6f26da0baa42874a907fff26bff1b7eb3bf27',
  status: '200',
  events: [],
  data: 'eyJu...nIn0=' }
```
