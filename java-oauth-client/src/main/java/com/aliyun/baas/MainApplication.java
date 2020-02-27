package com.aliyun.baas;

import com.aliyun.baas.rest.module.InlineResponse200;
import com.aliyun.baas.rest.module.InlineResponse2002;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.security.oauth2.client.OAuth2RestOperations;

import java.net.URLEncoder;

@SpringBootApplication
public class MainApplication implements CommandLineRunner {

    @Autowired
    @Qualifier("baasRestTemplate")
    private OAuth2RestOperations restTemplate;

    @Value("${baas.rest.endpoint}")
    private String endpoint;

    @Value("${baas.rest.network_name}")
    private String network;

    @Value("${baas.rest.contract_name}")
    private String contarct;

    public static void main(String[] args)
    {
        SpringApplication.run(MainApplication.class, args);
    }

    @Override
    public void run(String... args) throws Exception {
        String urlPrefix = endpoint + "/api/v1/networks/" + URLEncoder.encode(network, "UTF-8");

        // 查询链上区块
        ParameterizedTypeReference<InlineResponse2002> getResp = new ParameterizedTypeReference<InlineResponse2002>() {};
        ResponseEntity<InlineResponse2002> getResult = restTemplate.exchange(urlPrefix + "/blocks/1", HttpMethod.GET, null, getResp);
        System.out.println(getResult);

        // 调用 Fabric 智能合约
        String key = String.format("%d", System.currentTimeMillis());
        ParameterizedTypeReference<InlineResponse200> invokeResp = new ParameterizedTypeReference<InlineResponse200>() {};
        HttpEntity<String> invokeReq = new HttpEntity<String>(String.format("{\n" +
                "        \"chaincode\": \"%s\",\n" +
                "        \"args\": [\"put\", \"%s\", \"%s\"],\n" +
                "        \"transients\": {\n" +
                "            \"key\": \"value\"\n" +
                "        }\n" +
                "\t}", contarct, key, key));
        ResponseEntity<InlineResponse200> invokeResult = restTemplate.exchange(urlPrefix + "/transactions/invoke", HttpMethod.POST, invokeReq, invokeResp);
        System.out.println(invokeResult);

        // 查询 Fabric 智能合约
        ParameterizedTypeReference<InlineResponse200> queryResp = new ParameterizedTypeReference<InlineResponse200>() {};
        HttpEntity<String> queryReq = new HttpEntity<String>(String.format("{\n" +
                "        \"chaincode\": \"%s\",\n" +
                "        \"args\": [\"get\", \"%s\"],\n" +
                "        \"transients\": {\n" +
                "            \"key\": \"value\"\n" +
                "        }\n" +
                "\t}", contarct, key));
        ResponseEntity<InlineResponse200> queryResult = restTemplate.exchange(urlPrefix + "/transactions/query", HttpMethod.POST, queryReq, queryResp);
        System.out.println(queryResult);
    }
}
