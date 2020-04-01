# 链上数据导出到数据库

## 前置条件
1. 您需要先在区块链实例中[安装云服务集成](https://help.aliyun.com/document_detail/153337.html)
2. 拥有一个公网可连接的数据库实例（本示例以 RDS for MySQL 为例, 版本 >= 5.6）
3. 本地安装 nodev8 环境(>=8.17.0)


## 操作步骤
1. 连接到数据库实例，创建相关的库跟表
    1. 进入 [DMS 控制台](https://dms.console.aliyun.com/#/dms/login)，登陆数据库实例
    2. 在上方菜单栏选择 "SQL 操作 -> SQL 窗口" 开启一个新的 SQL 窗口
    3. 将 SQL 文件 [taskmgr.sql](./blockchain2db/taskmgr.sql) 中的内容复制到 SQL 窗口中执行；6条语句全部执行成功
2. 按照[发起示例交易](#send-demo-tx)的步骤1，在通道上安装并实例化链码 `taskmgr`
3. 在控制台创建云服务集成数据库触发器
    1. 选择链码 `taskmgr` 所在的通道及链码 `taskmgr`
    2. 配置触发器的事件类型为 Tx
    3. 数据库类型选择 MySQL （这里以 MySQL 为例，可以根据实际不同的数据库类型选择）
    4. 根据帮助，依次填写数据实例的地址、用户名和密码
    5. 数据库填写 `octopus`
    6. 表名填写 `write_set`
4. 按照[发起示例交易](#send-demo-tx)的后续步骤，在通道上发起示例交易
5. 之后您可以在数据库 `octopus` 中查看导出的数据
    1. 查看导出的原始交易写入集合 `SELECT * FROM write_set`
   
    |event_id                                                              |namespace|key                                                         |value                                                                                                                                                                                                                                                                        |create_time   |tx_id                                                           |is_delete|creator                                   |
    |----------------------------------------------------------------------|---------|------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------|----------------------------------------------------------------|---------|------------------------------------------|
    |tx-17-4e62ed86004932f03ae366e2c021a6abfd4f5da642e4067a96c1bf96e3251149|taskmgr  |jc2yf010MSP.octopus_26842_12345678901234task-1584944026722|{"name":"task-1584944026722","creator":"jc2yf010MSP.octopus_26842_12345678901234","requires":["e2ehmfqasthMSP.octopus_26842_12345678901234"],"approved":null,"description":"示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'"}                                          |2020/3/23 6:13|4e62ed86004932f03ae366e2c021a6abfd4f5da642e4067a96c1bf96e3251149|0        |jc2yf010MSP.octopus_26842_12345678901234|
    |tx-18-ff6c465a12d19ddba643a073d7571b3a8ce80a3de3bafc4cc62b0809521c483f|taskmgr  |jc2yf010MSP.octopus_26842_12345678901234task-1584944026722|{"name":"task-1584944026722","creator":"jc2yf010MSP.octopus_26842_12345678901234","requires":["e2ehmfqasthMSP.octopus_26842_12345678901234"],"approved":["jc2yf010MSP.octopus_26842_12345678901234"],"description":"示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'"}|2020/3/23 6:13|ff6c465a12d19ddba643a073d7571b3a8ce80a3de3bafc4cc62b0809521c483f|0        |jc2yf010MSP.octopus_26842_12345678901234|
    2.  查看导出的审批任务详情 `SELECT * FROM taskmgr`

    |key                                                                   |task_name|creator                                                     |require                                                                                                                                                                                                                                                                      |approved      |description                                                     |update_time|
    |----------------------------------------------------------------------|---------|------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------|----------------------------------------------------------------|-----------|
    |jc2yf010MSP.octopus_26842_12345678901234task-1584944026722          |task-1584944026722|jc2yf010MSP.octopus_26842_12345678901234                  |["e2ehmfqasthMSP.octopus_26842_12345678901234"]                                                                                                                                                                                                                            |["jc2yf010MSP.octopus_26842_12345678901234"]|示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'               |2020/3/23 14:13|


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

