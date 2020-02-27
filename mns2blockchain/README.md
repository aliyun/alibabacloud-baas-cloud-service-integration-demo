# 消息队列MNS消息直接上链

本示例介绍如何通过函数计算，自动将消息队列MNS中的消息推送到区块链中。

## 前置条件
1. 您需要先在区块链实例中[安装云服务集成](https://help.aliyun.com/document_detail/153337.html)
2. 在本地安装函数计算工具 funcraft, [funcraft 介绍及安装方式](https://help.aliyun.com/document_detail/140283.html)

## 使用方法

1. 进入示例代码目录, 修改示例程序的以下内容
  1. 按照注释修改 index.js 文件上方的配置项，包括REST API服务地址、Refresh Token、通道名称、智能合约名称
  2. 根据业务场景，按照注释修改调用智能合约的方式和参数
  3. 修改 template.yml 中的函数计算服务名称和函数名称，默认服务名为"octopus"，函数名为"MNS2BlockChain"。更多配置方式可以参考 [funcraft文档](https://github.com/alibaba/funcraft/blob/master/docs/specs/2018-04-03-zh-cn.md)
2. 在 mns2blockchain 目录，执行 `fun deploy` 将函数部署到函数计算
3. 进入[函数计算控制台](https://fc.console.aliyun.com/)，为我们刚才创建的函数配置 MNS 触发器(示例程序中 [Event 格式] 需要选择 JSON)
  ![触发器配置](http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/pic/151957/cn_zh/1582165817972/C46853E4-334F-4CC2-8635-E36DB1B540CD.png)


