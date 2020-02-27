package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-oauth-client/swagger"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)
func getOAuthClient(tokenEndpoint string, refresh_token string) *http.Client {
	config := &oauth2.Config{
		Endpoint:     oauth2.Endpoint{
			TokenURL: tokenEndpoint,
		},
	}
	token := &oauth2.Token{
		AccessToken:  "none",
		TokenType:    "bearer",
		RefreshToken: refresh_token,
		Expiry:       time.Now().Add(-time.Second),
	}
	return config.Client(context.Background(), token)
}

func assertNoError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err.Error())
		os.Exit(1)
	}
}

func main() {
	// ---- 根据您的实例信息配置以下参数
	const (
		// 区块链 REST API 服务地址
		baasRESTApiEndpoint = "http://your.gateway.endpoint"
		// 在区块链控制台生成的 refresh token
		baasRESTApiRefreshToken = "Yt+MCV4NzwEaZC6t3C848xipJfLCW1O7t5BY3gJSbetvUTAh/HWxfmTeqT7CpGE8cBWWYYcgz46DMzEAbqLzMvj0aG1dC8IkVLT6ZkEKxX1ioS6ZomFeSlcGcsnNzcPe605uwqZXcGNdjDh6WzUBuBWgP9gh6i0EMs2aaNyvkOLiOQx/+aWm04dLP9Q2L/eJAhIWkWjjOtOd4oMVHP7FOJoZux6zRfz4hSbCYyvl4A6WV3tfrD3FUnVq7E74CJKjMgWfTbjD6+QzkKsbbs/hD/qGF9bAMPmVe1o="
		// 区块链网络名称（Fabric 通道名称）
		baasNetworkName = "channel3"
		// 区块链智能合约名称（Fabric 链码名称）
		baasContractname = "notary"
	)
	cli := getOAuthClient(
		baasRESTApiEndpoint + "/api/v1/token",
		baasRESTApiRefreshToken,
	)
	urlPrefix := baasRESTApiEndpoint + "/api/v1/networks/" + url.QueryEscape(baasNetworkName)

	// 查询链上区块
	resp, err := cli.Get(urlPrefix + "/blocks/1")
	assertNoError(err, "Send get request to Octopus failed")
	data, err := ioutil.ReadAll(resp.Body)
	assertNoError(err, "Read get response body failed")
	getResp := &swagger.InlineResponse2002{}
	err = json.Unmarshal(data, getResp)
	assertNoError(err, "Unmarshal response failed")
	fmt.Printf("Block response body: %s\n", string(data))
	fmt.Printf("Block data: %s\n", getResp.Result.Data)

	// 调用 Fabric 智能合约
	key := fmt.Sprintf("%d", time.Now().Unix())
	resp, err = cli.Post(urlPrefix + "/transactions/invoke", "application/json", strings.NewReader(fmt.Sprintf(`
	{
        "chaincode": "%s",
        "args": ["put", "%s", "%s"],
        "transients": {
            "key": "value"
        }
	}
	`, baasContractname, key, key)))
	assertNoError(err, "Send invoke request to Octopus failed")
	data, err = ioutil.ReadAll(resp.Body)
	assertNoError(err, "Read invoke response body failed")
	invokeResp := &swagger.InlineResponse200{}
	err = json.Unmarshal(data, invokeResp)
	assertNoError(err, "Unmarshal response failed")
	fmt.Printf("Invoke response body: %s\n", string(data))
	fmt.Printf("Invoke contract response: %s\n", string(invokeResp.Result.Data))

	// 查询 Fabric 智能合约
	resp, err = cli.Post(urlPrefix + "/transactions/query", "application/json", strings.NewReader(fmt.Sprintf(`
	{
        "chaincode": "%s",
        "args": ["get", "%s"],
        "transients": {
            "key": "value"
        }
	}
	`, baasContractname, key)))
	assertNoError(err, "Send query request to Octopus failed")
	data, err = ioutil.ReadAll(resp.Body)
	assertNoError(err, "Read query response body failed")
	queryResp := &swagger.InlineResponse200{}
	err = json.Unmarshal(data, queryResp)
	assertNoError(err, "Unmarshal response failed")
	fmt.Printf("Query response body: %s\n", string(data))
	fmt.Printf("Query contract response: %s\n", string(queryResp.Result.Data))
}
