# Java REST-API Client 示例


本示例为您展示如何通过 Java 访问阿里云区块链 REST-API。示例主要演示了如下操作：
1. 获取指定通道区块号为1的区块
2. 调用通道中的 Fabric 智能合约
3. 调用通道中的 Fabric 智能合约的查询接口

## 使用方法

1. 配置本地 java 1.8 环境

2. 按照文档[部署链码](https://help.aliyun.com/document_detail/85739.html)将示例链码 [notary](./contracts/fabric/notary) 部署到通道中

3. 根据您组织实例的 REST-API 地址、生成的 Token 信息、以及业务通道和链码信息，修改文件 `src/main/resources/application.properties`

4. 执行 `mvn spring-boot:run` 运行示例程序

**正常结果示例**

```
> mvn spring-boot:run
  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v2.0.6.RELEASE)

2020-02-17 18:00:09.056  INFO 79141 --- [           main] com.aliyun.baas.MainApplication          : Starting MainApplication on Bright.local with PID 79141 (java-oauth-client/target/classes started by bright in java-oauth-client)
2020-02-17 18:00:09.059  INFO 79141 --- [           main] com.aliyun.baas.MainApplication          : No active profile set, falling back to default profiles: default
2020-02-17 18:00:09.092  INFO 79141 --- [           main] s.c.a.AnnotationConfigApplicationContext : Refreshing org.springframework.context.annotation.AnnotationConfigApplicationContext@5bf85360: startup date [Mon Feb 17 18:00:09 CST 2020]; root of context hierarchy
2020-02-17 18:00:09.775  INFO 79141 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on startup
2020-02-17 18:00:09.792  INFO 79141 --- [           main] com.aliyun.baas.MainApplication          : Started MainApplication in 0.958 seconds (JVM running for 3.321)
<200,class InlineResponse2002 {
    success: true
    result: class Block {
        number: 1
        hash: 1c397c4eb3e0e330c01ec430170f844e46159f16930aa347486b8153b6586548
        previousHash: 88ef0ad6ba2df7ba7e53de78575d2d14cee2253fe6897305e50b57ceeecebc78
        createTime: 1579056958
        transactions: []
        data: {data={data=[{payload={data={config={channel_group= ... }}}}]}}
    }
    error: class Error {
        code: 200
        message: Success
        requestId: edf8fe52-7cef-447f-a04a-7b8c1db56487
    }
},{Server=[nginx], Date=[Mon, 17 Feb 2020 10:00:10 GMT], Content-Type=[application/json; charset=UTF-8], Transfer-Encoding=[chunked], Connection=[keep-alive]}>
<200,class InlineResponse200 {
    success: true
    result: class Response {
        id: a5f5503f12b92a4c59e079e1baf49b517785e2d9988f90dfb234f6c3954a2389
        status: 200
        event: null
        data: MTU4MTkzMzYxMDEzNg==
    }
    error: class Error {
        code: 200
        message: Success
        requestId: 71a9f95f-ea5b-4dea-b4a1-a608ae429fb4
    }
},{Server=[nginx], Date=[Mon, 17 Feb 2020 10:00:11 GMT], Content-Type=[application/json; charset=UTF-8], Content-Length=[249], Connection=[keep-alive]}>
<200,class InlineResponse200 {
    success: true
    result: class Response {
        id: ba50180c9fe38c9be115f20775b78e80b7a1205c34ef34b66fab635efedc3b49
        status: 200
        event: null
        data: MTU4MTkzMzYxMDEzNg==
    }
    error: class Error {
        code: 200
        message: Success
        requestId: c703a36b-3589-4a8b-87a0-5e5bf56b2396
    }
},{Server=[nginx], Date=[Mon, 17 Feb 2020 10:00:11 GMT], Content-Type=[application/json; charset=UTF-8], Content-Length=[249], Connection=[keep-alive]}>
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  4.217 s
[INFO] Finished at: 2020-02-17T18:00:11+08:00
[INFO] ------------------------------------------------------------------------
2020-02-17 18:00:11.573  INFO 79141 --- [       Thread-3] s.c.a.AnnotationConfigApplicationContext : Closing org.springframework.context.annotation.AnnotationConfigApplicationContext@5bf85360: startup date [Mon Feb 17 18:00:09 CST 2020]; root of context hierarchy
2020-02-17 18:00:11.574  INFO 79141 --- [       Thread-3] o.s.j.e.a.AnnotationMBeanExporter        : Unregistering JMX-exposed beans on shutdown
```

