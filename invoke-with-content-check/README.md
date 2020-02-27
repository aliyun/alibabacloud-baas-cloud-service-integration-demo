# 数据合规上链

使用REST API将数据上链时，我们可以通过阿里云集成内容安全服务，对上链内容进行前置检查，避免不合规数据上链。

本示例以 nodejs 为例介绍如何使用REST API在数据上链前进行内容合规检查：

## 前置条件
1. 您需要先在区块链实例中[安装云服务集成](https://help.aliyun.com/document_detail/153337.html), 并开通内容安全集成功能
3. 本地安装 nodev8 环境(>=8.17.0)

## 使用方法
1. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [notary](./contracts/fabric/notary) 部署到通道中

2. 参考文档[使用REST API](https://help.aliyun.com/document_detail/151690.html), 修改示例中的 main.js

3. 执行 `npm install` 安装 node 依赖项 

4. 使用 `node main.js` 启动示例

## 示例输出

```
Data 1581905807512 pushed to blockchain with transaction 284d0b2b89db5bc5489127de863d0bb9b9d0a5f05bae67762567a3aff113822a
Content Moderation check failed: Send transaction failed: CONTENT_CHECK returned error VERIFY_FAILED: Content Moderation Check failed, suggestion block. key: 1581905807512, value: 妈跟非洲野驴怎么生下你这个骡子? 你爹是怎么操非洲野猪操了你个鳖货? 我用放大镜也木
Content Moderation check failed: Send transaction failed: CONTENT_CHECK returned error VERIFY_FAILED: Content Moderation Check failed, suggestion block. key: 1581905807512, value: 1月16日,在对缅甸联邦共和国进行国事访问前夕,国家主席习近平在缅甸
```

## 高级

如果默认的内容安全检查策略不能满足您业务的要求，您可以登陆[内容安全控制台](https://yundun.console.aliyun.com/?p=cts#/api/statistics?tabIndex=1)，参考文档[自定义文本库](https://help.aliyun.com/document_detail/66057.html)来扩充检查规则或对上链内容检查的行为进行调整。