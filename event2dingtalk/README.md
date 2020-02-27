# 链上事件触发钉钉群通知

本示例为您展示如何利用阿里云区块链云服务集成功能，将链上的事件推送至钉钉；当区块链上发生事件时，您在钉钉群中能够看到相应的通知。

事件流： Blockchain Node  --> Cloud Service Intergration --> Function Compute --> Dingtalk

## 前置条件
1. 您需要先在区块链实例中[安装云服务集成](https://help.aliyun.com/document_detail/153337.html), 并开通函数计算集成功能
2. 本地安装 nodev8 环境(>=8.17.0)
3. 安装函数计算工具 funcraft, [funcraft 介绍及安装方式](https://help.aliyun.com/document_detail/140283.html)

## 操作步骤
1. 创建钉钉群，并添加自定义机器人
  1. 进入钉钉群设置->智能群助手->添加机器，选择“自定义”
  2. 这里以签名为例，安全设置选择加签。可以根据实际情况选择其他方式，可以参考[钉钉群机器人](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)
2. 进入 `fc` 目录，修改以下内容
  1. 按照注释修改 index.js 中的配置参数，填入钉钉回调地址、推送签名密钥
  2. 修改 template.yml 中的函数计算服务名称和函数名称，默认服务名为"octopus"，函数名为"Event2Dingtalk"。更多配置方式可以参考 [funcraft文档](https://github.com/alibaba/funcraft/blob/master/docs/specs/2018-04-03-zh-cn.md)
3. 使用 `fun deploy` 将函数部署到函数计算
4. 在控制台创建云服务集成函数计算触发器
  1. 配置触发器的事件类型为 Contact
  2. 根据帮助，填写函数计算实例的相关信息

5. 按照[发起示例交易](./blockchain2sms/README.md#send-demo-tx)的步骤，在通道上发起示例交易
6. 当审批流创建、结束时，钉钉机器人会在群中发送消息

![钉钉集成示例](http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/pic/151956/cn_zh/1581922247704/65DB59AF-686A-4D9E-844B-F6C66AD9A0A9.png)
