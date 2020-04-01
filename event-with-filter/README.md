# 事件过滤推送

在配置云服务集成触发器时，您可以为触发器配置过滤器，过滤器遵循 jq 工具的基本[语法](https://stedolan.github.io/jq/manual/v1.6/)，您可以通过表达式对数据进行过滤和简单的处理。

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
  "type": "Contract"
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
  "type": "Tx"
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

### 六、推送到数据库表

通过过滤器我们可以将区块链数据推送到数据库的自定义表中；通过 filter 对事件进行处理，将事件转换成一个 Object 的数组，每个 Object 表示需要插入的一行数据，Object 的 Key 为表的列名， Value 为需要插入的数据：
```
[
  {
    "columnName1": "value1",
    "columnName2": "value2",
    "columnName3": "value3",
  },
  {
    "columnName1": "value3",
    "columnName2": "value4",
    "columnName3": "value5",
  },
  ...
]
```

将智能合约 taskmgr 所有合法交易写入的 Key 导出到以下表结构中

- 表结构，主键： (`event_id`,`namespace`,`key`)
  - event_id, 事件ID
  - namespace, 写入的命名空间（智能合约名称）
  - key, 写入的Key
  - value, 经过Base64编码的写入值
  - create_time, 写入的时间
  - tx_id, 对应的交易ID
  - is_delete, 是否是删除操作
  - creator, 写入操作的发起者
- filter 配置：`select(.content.state=="VALID" and .content.to=="taskmgr") | .id as $eventID | .content.id as $txID | .content.from as $creator | .content.data.header.channel_header.timestamp as $createTime | .content.data.data.actions.[0].payload.action.proposal_response_payload.extension.results.ns_rwset | map(.namespace as $namespace | .rwset.writes | map( .event_id = $eventID | .namespace = $namespace | .tx_id = $txID | .create_time = $createTime | .creator = $creator )) | add`
- 事件类型选择： Tx
- 推送数据示例：
```
[
  {
    "event_id": "tx-128-3a2e0e375c6ee0d9c6c20d756e787db44d50bfeba51f22d8f262dd1556dedba0",
    "namespace": "taskmgr",
    "tx_id": "3a2e0e375c6ee0d9c6c20d756e787db44d50bfeba51f22d8f262dd1556dedba0",
    "is_delete": false,
    "key": "\u0000e2ehmfqasthMSP.octopus_26842_12345678901234\u0000task-1581479306267\u0000",
    "value": "eyJuYW1lIjoidGFzay0xNTgwOTcxODQyOTM3IiwiY3JlYXRvciI6ImUyZWhtZnFhc3RoTVNQLm9jdG9wdXNfMjY4NDJfMTIzNDU2Nzg5MDEyMzQiLCJyZXF1aXJlcyI6WyJlMmVobWZxYXN0aE1TUC5vY3RvcHVzXzI2ODQyXzEyMzQ1Njc4OTAxMjM0Il0sImFwcHJvdmVkIjpudWxsLCJkZXNjcmlwdGlvbiI6IuekuuS+i+S7u+WKoe+8jHJlcXVpcmVzIOmFjee9ruWuoeaJueS7u+WKoeWujOaIkOmcgOimgemCo+S6m+eUqOaIt+WQjOaEj+OAgueUqOaIt+aPj+i/sOS4uiAn57uE57uHTVNQLueUqOaIt+WQjeensCcifQ==",
    "create_time": "2020-02-06T06:50:42.267092Z",
    "creator": "e2ehmfqasthMSP.octopus_26842_12345678901234"
  }
]
```


## 示例消息

### 合约事件（Contract）

```
{
	"type": "Contract",
	"platform": "Fabric",
	"instance_id": "csi-potato20-hs51a74qevzj",
	"network": "channel2",
	"id": "contract-3-233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
	"name": "event-create-task",
	"content": "eyJuY...nIn0="
}
```

### 交易事件（Tx）

```
{
	"type": "Tx",
	"platform": "Fabric",
	"instance_id": "csi-potato20-hs51a74qevzj",
	"network": "channel2",
	"id": "tx-3-233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
	"name": "233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
	"content": {
		"id": "233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
		"state": "VALID",
		"from": "potato20MSP.octopus_26842_12345678901234",
		"to": "taskmgr",
		"input": "[\"create\",\"task-1585319773698\",\"{\\n            \\\"requires\\\": [\\\"e2ehmfqasthMSP.octopus_26842_12345678901234\\\"],\\n            \\\"description\\\": \\\"示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'\\\"\\n        }\"]",
		"events": [
			{
				"type": "Contract",
				"platform": "Fabric",
				"instance_id": "csi-potato20-hs51a74qevzj",
				"network": "",
				"id": "contract-3-233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
				"name": "event-create-task",
				"content": "eyJuY...AnIn0="
			}
		],
		"data": {
			"data": {
				"actions": [
					{
						"header": {
							"creator": {
								"id_bytes": "LS0tL...S0tLS0K",
								"mspid": "potato20MSP"
							},
							"nonce": "m/ytbXHad18UlE2FWTPt5WKTtf8FVUxs"
						},
						"payload": {
							"action": {
								"endorsements": [
									{
										"endorser": "Cgtwb3...0tLQo=",
										"signature": "MEQCIDpiO+qfGGxus+PivghVNv5yrvBW1xjRYKOHNIuiM5XqAiBwyBcteAwprZhsqVZEWD2J4+HGu+StRnYqfxkynBvf8g=="
									}
								],
								"proposal_response_payload": {
									"extension": {
										"chaincode_id": {
											"name": "taskmgr",
											"path": "",
											"version": "1.4"
										},
										"events": {
											"chaincode_id": "taskmgr",
											"event_name": "event-create-task",
											"payload": "eyJuYW1lI...np7AnIn0=",
											"tx_id": "233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03"
										},
										"response": {
											"message": "",
											"payload": "dGFzay0xNTg1MzE5NzczNjk4",
											"status": 200
										},
										"results": {
											"data_model": "KV",
											"ns_rwset": [
												{
													"collection_hashed_rwset": [],
													"namespace": "lscc",
													"rwset": {
														"metadata_writes": [],
														"range_queries_info": [],
														"reads": [
															{
																"key": "taskmgr",
																"version": {
																	"block_num": "2",
																	"tx_num": "0"
																}
															}
														],
														"writes": []
													}
												},
												{
													"collection_hashed_rwset": [],
													"namespace": "taskmgr",
													"rwset": {
														"metadata_writes": [],
														"range_queries_info": [],
														"reads": [
															{
																"key": "\u0000potato20MSP.octopus_26842_12345678901234\u0000task-1585319773698\u0000",
																"version": null
															}
														],
														"writes": [
															{
																"is_delete": false,
																"key": "\u0000potato20MSP.octopus_26842_12345678901234\u0000task-1585319773698\u0000",
																"value": "eyJuYW...AnIn0="
															}
														]
													}
												}
											]
										}
									},
									"proposal_hash": "Ftf//Jn+yGZLqTN3YKRW0kAjjBZccBTDyY77cYDrRzw="
								}
							},
							"chaincode_proposal_payload": {
								"TransientMap": {},
								"input": {
									"chaincode_spec": {
										"chaincode_id": {
											"name": "taskmgr",
											"path": "",
											"version": ""
										},
										"input": {
											"args": [
												"Y3JlYXRl",
												"dGFzay0xNTg1MzE5NzczNjk4",
												"ewogIC...gICB9"
											],
											"decorations": {},
											"is_init": false
										},
										"timeout": 0,
										"type": "GOLANG"
									}
								}
							}
						}
					}
				]
			},
			"header": {
				"channel_header": {
					"channel_id": "channel2",
					"epoch": "0",
					"extension": "EgkSB3Rhc2ttZ3I=",
					"timestamp": "2020-03-27T14:36:13.834822939Z",
					"tls_cert_hash": null,
					"tx_id": "233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03",
					"type": 3,
					"version": 0
				},
				"signature_header": {
					"creator": {
						"id_bytes": "LS0tLS1C...FLS0tLS0K",
						"mspid": "potato20MSP"
					},
					"nonce": "m/ytbXHad18UlE2FWTPt5WKTtf8FVUxs"
				}
			}
		}
	}
}
```

### 块事件（Block）

```
{
	"type": "Block",
	"platform": "Fabric",
	"instance_id": "csi-potato20-hs51a74qevzj",
	"network": "channel2",
	"id": "block-3",
	"name": "3",
	"content": {
		"number": 3,
		"hash": "39e78092b63c530abffdf7b39590c1d56eaccbb64757155a3c85054e757095c1",
		"create_time": 1585319773,
		"previous_hash": "6de5bd1d51b8f174838e4b7300f2ca05fe26e9774544222b4e938c59c849974f",
		"transactions": [
			"233718ea379a7ef798f40bfd6ae77f842145334e3d5520292bea7aadfa9faf03"
		]
	}
}
```

### 配置事件（Config）

```
{
	"type": "Config",
	"platform": "Fabric",
	"instance_id": "csi-potato20-hs51a74qevzj",
	"network": "channel2",
	"id": "config-7-",
	"name": "",
	"content": {
		"id": "",
		"state": "VALID",
		"from": "jc2yf1MSP.orderer1",
		"to": "",
		"input": "{\n\t\"channel_id\": \"channel2\",\n\t\"isolated_data\": {},\n\t\"read_set\": {\n\t\t\"groups\": {\n\t\t\t\"Orderer\": {\n\t\t\t\t\"groups\": {},\n\t\t\t\t\"mod_policy\": \"\",\n\t\t\t\t\"policies\": {},\n\t\t\t\t\"values\": {},\n\t\t\t\t\"version\": \"0\"\n\t\t\t}\n\t\t},\n\t\t\"mod_policy\": \"\",\n\t\t\"policies\": {},\n\t\t\"values\": {},\n\t\t\"version\": \"0\"\n\t},\n\t\"write_set\": {\n\t\t\"groups\": {\n\t\t\t\"Orderer\": {\n\t\t\t\t\"groups\": {},\n\t\t\t\t\"mod_policy\": \"\",\n\t\t\t\t\"policies\": {},\n\t\t\t\t\"values\": {\n\t\t\t\t\t\"BatchSize\": {\n\t\t\t\t\t\t\"mod_policy\": \"Admins\",\n\t\t\t\t\t\t\"value\": {\n\t\t\t\t\t\t\t\"absolute_max_bytes\": 103809024,\n\t\t\t\t\t\t\t\"max_message_count\": 200,\n\t\t\t\t\t\t\t\"preferred_max_bytes\": 5242880\n\t\t\t\t\t\t},\n\t\t\t\t\t\t\"version\": \"2\"\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\t\"version\": \"0\"\n\t\t\t}\n\t\t},\n\t\t\"mod_policy\": \"\",\n\t\t\"policies\": {},\n\t\t\"values\": {},\n\t\t\"version\": \"0\"\n\t}\n}\n",
		"events": null,
		"data": {
			"data": {
				"config": {
					"channel_group": {
						"groups": {
							"Application": {
								"groups": {
									"potato20MSP": {
										"groups": {},
										"mod_policy": "Admins",
										"policies": {
											"Admins": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "potato20MSP",
																	"role": "ADMIN"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											},
											"Readers": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "potato20MSP",
																	"role": "MEMBER"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											},
											"Writers": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "potato20MSP",
																	"role": "MEMBER"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											}
										},
										"values": {
											"AnchorPeers": {
												"mod_policy": "Admins",
												"value": {
													"anchor_peers": [
														{
															"host": "peer1",
															"port": 31111
														}
													]
												},
												"version": "0"
											},
											"MSP": {
												"mod_policy": "Admins",
												"value": {
													"config": {
														"admins": [
															"LS0tL...0tLQo="
														],
														"crypto_config": {
															"identity_identifier_hash_function": "SHA256",
															"signature_hash_family": "SHA2"
														},
														"fabric_node_ous": {
															"admin_ou_identifier": null,
															"client_ou_identifier": {
																"certificate": "LS0tL...0tLS0K",
																"organizational_unit_identifier": "client"
															},
															"enable": true,
															"orderer_ou_identifier": null,
															"peer_ou_identifier": {
																"certificate": "LS0tL...0tLS0K",
																"organizational_unit_identifier": "peer"
															}
														},
														"intermediate_certs": [
															"LS0tL...0tLS0K",
															"LS0tL...0tLS0K"
														],
														"name": "potato20MSP",
														"organizational_unit_identifiers": [],
														"revocation_list": [],
														"root_certs": [
															"LS0tL...0tLS0K"
														],
														"signing_identity": null,
														"tls_intermediate_certs": [
															"LS0tL...0tLS0tCg==",
															"LS0tL...0tLS0K"
														],
														"tls_root_certs": [
															"LS0tL...0tLQo="
														]
													},
													"type": 0
												},
												"version": "0"
											}
										},
										"version": "1"
									}
								},
								"mod_policy": "Admins",
								"policies": {
									"Admins": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "MAJORITY",
												"sub_policy": "Admins"
											}
										},
										"version": "0"
									},
									"Readers": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "ANY",
												"sub_policy": "Readers"
											}
										},
										"version": "0"
									},
									"Writers": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "ANY",
												"sub_policy": "Writers"
											}
										},
										"version": "0"
									}
								},
								"values": {
									"Capabilities": {
										"mod_policy": "Admins",
										"value": {
											"capabilities": {
												"V1_4_2": {}
											}
										},
										"version": "0"
									}
								},
								"version": "1"
							},
							"Orderer": {
								"groups": {
									"jc2yf1MSP": {
										"groups": {},
										"mod_policy": "Admins",
										"policies": {
											"Admins": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "jc2yf1MSP",
																	"role": "ADMIN"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											},
											"Readers": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "jc2yf1MSP",
																	"role": "MEMBER"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											},
											"Writers": {
												"mod_policy": "Admins",
												"policy": {
													"type": 1,
													"value": {
														"identities": [
															{
																"principal": {
																	"msp_identifier": "jc2yf1MSP",
																	"role": "MEMBER"
																},
																"principal_classification": "ROLE"
															}
														],
														"rule": {
															"n_out_of": {
																"n": 1,
																"rules": [
																	{
																		"signed_by": 0
																	}
																]
															}
														},
														"version": 0
													}
												},
												"version": "0"
											}
										},
										"values": {
											"Endpoints": {
												"mod_policy": "/Channel/Orderer/Admins",
												"value": {
													"addresses": [
														"orderer1.jc2yf1.aliyunbaas.top:31010",
														"orderer2.jc2yf1.aliyunbaas.top:31020",
														"orderer3.jc2yf1.aliyunbaas.top:31030"
													]
												},
												"version": "0"
											},
											"MSP": {
												"mod_policy": "Admins",
												"value": {
													"config": {
														"admins": [
															"LS0tL...0tLS0K"
														],
														"crypto_config": {
															"identity_identifier_hash_function": "SHA256",
															"signature_hash_family": "SHA2"
														},
														"fabric_node_ous": null,
														"intermediate_certs": [
															"LS0tL...0tLS0K",
															"LS0tL...0tLS0K"
														],
														"name": "jc2yf1MSP",
														"organizational_unit_identifiers": [],
														"revocation_list": [],
														"root_certs": [
															"LS0tL...0tLS0K"
														],
														"signing_identity": null,
														"tls_intermediate_certs": [
															"LS0tL...0tLS0K",
															"LS0tL...0tLS0K"
														],
														"tls_root_certs": [
															"LS0tL...0tLS0K"
														]
													},
													"type": 0
												},
												"version": "0"
											}
										},
										"version": "0"
									}
								},
								"mod_policy": "Admins",
								"policies": {
									"Admins": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "MAJORITY",
												"sub_policy": "Admins"
											}
										},
										"version": "0"
									},
									"BlockValidation": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "ANY",
												"sub_policy": "Writers"
											}
										},
										"version": "0"
									},
									"Readers": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "ANY",
												"sub_policy": "Readers"
											}
										},
										"version": "0"
									},
									"Writers": {
										"mod_policy": "Admins",
										"policy": {
											"type": 3,
											"value": {
												"rule": "ANY",
												"sub_policy": "Writers"
											}
										},
										"version": "0"
									}
								},
								"values": {
									"BatchSize": {
										"mod_policy": "Admins",
										"value": {
											"absolute_max_bytes": 103809024,
											"max_message_count": 200,
											"preferred_max_bytes": 5242880
										},
										"version": "2"
									},
									"BatchTimeout": {
										"mod_policy": "Admins",
										"value": {
											"timeout": "1s"
										},
										"version": "0"
									},
									"Capabilities": {
										"mod_policy": "Admins",
										"value": {
											"capabilities": {
												"V1_4_2": {}
											}
										},
										"version": "0"
									},
									"ChannelRestrictions": {
										"mod_policy": "Admins",
										"value": null,
										"version": "0"
									},
									"ConsensusType": {
										"mod_policy": "Admins",
										"value": {
											"metadata": {
												"consenters": [
													{
														"client_tls_cert": "LS0tL...0tLS0K",
														"host": "orderer1",
														"port": 31010,
														"server_tls_cert": "LS0tL...0tLS0K"
													},
													{
														"client_tls_cert": "LS0tL...0tLS0K",
														"host": "orderer2",
														"port": 31020,
														"server_tls_cert": "LS0tL...0tLS0K"
													},
													{
														"client_tls_cert": "LS0tL...0tLS0K",
														"host": "orderer3",
														"port": 31030,
														"server_tls_cert": "LS0tL...0tLS0K"
													}
												],
												"options": {
													"election_tick": 10,
													"heartbeat_tick": 1,
													"max_inflight_blocks": 5,
													"snapshot_interval_size": 20971520,
													"tick_interval": "500ms"
												}
											},
											"state": "STATE_NORMAL",
											"type": "etcdraft"
										},
										"version": "0"
									}
								},
								"version": "0"
							}
						},
						"mod_policy": "Admins",
						"policies": {
							"Admins": {
								"mod_policy": "Admins",
								"policy": {
									"type": 3,
									"value": {
										"rule": "MAJORITY",
										"sub_policy": "Admins"
									}
								},
								"version": "0"
							},
							"Readers": {
								"mod_policy": "Admins",
								"policy": {
									"type": 3,
									"value": {
										"rule": "ANY",
										"sub_policy": "Readers"
									}
								},
								"version": "0"
							},
							"Writers": {
								"mod_policy": "Admins",
								"policy": {
									"type": 3,
									"value": {
										"rule": "ANY",
										"sub_policy": "Writers"
									}
								},
								"version": "0"
							}
						},
						"values": {
							"BlockDataHashingStructure": {
								"mod_policy": "Admins",
								"value": {
									"width": 4294967295
								},
								"version": "0"
							},
							"Capabilities": {
								"mod_policy": "Admins",
								"value": {
									"capabilities": {
										"V1_4_2": {}
									}
								},
								"version": "0"
							},
							"Consortium": {
								"mod_policy": "Admins",
								"value": {
									"name": "SampleConsortium"
								},
								"version": "0"
							},
							"HashingAlgorithm": {
								"mod_policy": "Admins",
								"value": {
									"name": "SHA256"
								},
								"version": "0"
							},
							"OrdererAddresses": {
								"mod_policy": "/Channel/Orderer/Admins",
								"value": {
									"addresses": [
										"orderer1:31010",
										"orderer2:31020",
										"orderer3:31030"
									]
								},
								"version": "0"
							}
						},
						"version": "0"
					},
					"sequence": "4"
				},
				"last_update": {
					"payload": {
						"data": {
							"config_update": {
								"channel_id": "channel2",
								"isolated_data": {},
								"read_set": {
									"groups": {
										"Orderer": {
											"groups": {},
											"mod_policy": "",
											"policies": {},
											"values": {},
											"version": "0"
										}
									},
									"mod_policy": "",
									"policies": {},
									"values": {},
									"version": "0"
								},
								"write_set": {
									"groups": {
										"Orderer": {
											"groups": {},
											"mod_policy": "",
											"policies": {},
											"values": {
												"BatchSize": {
													"mod_policy": "Admins",
													"value": {
														"absolute_max_bytes": 103809024,
														"max_message_count": 200,
														"preferred_max_bytes": 5242880
													},
													"version": "2"
												}
											},
											"version": "0"
										}
									},
									"mod_policy": "",
									"policies": {},
									"values": {},
									"version": "0"
								}
							},
							"signatures": [
								{
									"signature": "MEQCIGWQt4357PAFdosHtFLn1dvkFiSxfOoo+dXG0o4Cy6GZAiADcEdd4PyrRXeTvZwLLKD+D3mWmhX90Nr/9WtcgD/vpg==",
									"signature_header": {
										"creator": {
											"id_bytes": "LS0tL...0tLS0K",
											"mspid": "jc2yf1MSP"
										},
										"nonce": "0uCe85p1bTx2CvrXOlijzSoTR+2CqyUy"
									}
								}
							]
						},
						"header": {
							"channel_header": {
								"channel_id": "channel2",
								"epoch": "0",
								"extension": null,
								"timestamp": "2020-03-27T14:43:55Z",
								"tls_cert_hash": null,
								"tx_id": "",
								"type": 2,
								"version": 0
							},
							"signature_header": {
								"creator": {
									"id_bytes": "LS0tL...0tLS0K",
									"mspid": "jc2yf1MSP"
								},
								"nonce": "3qVtmZVkV54mIflbrTokcVCK9pxuaocR"
							}
						}
					},
					"signature": "MEUCIQDPlMNSlZuOMCFTvHvs/FuZGGuWKhOu7U8BUCyHe+OzPwIgPKyz1fC9zTEGW0tnLywFuSIwMzUtLlpaYP3mMVx9krY="
				}
			},
			"header": {
				"channel_header": {
					"channel_id": "channel2",
					"epoch": "0",
					"extension": null,
					"timestamp": "2020-03-27T14:43:55Z",
					"tls_cert_hash": null,
					"tx_id": "",
					"type": 1,
					"version": 0
				},
				"signature_header": {
					"creator": {
						"id_bytes": "LS0tL...0tLS0K",
						"mspid": "jc2yf1MSP"
					},
					"nonce": "VIjuky8xU2uTkkLx4YrgJLoI7uvYrWV1"
				}
			}
		}
	}
}
```
