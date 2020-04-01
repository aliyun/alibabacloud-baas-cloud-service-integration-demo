const axios = require('axios');
const oauth = require('axios-oauth-client');
const tokenProvider = require('axios-token-interceptor');

// ---- 根据您的实例信息配置以下参数
// 区块链 REST API 服务地址
const baasRESTApiEndpoint = 'http://your.gateway.endpoint';
// 在区块链控制台生成的 refresh token
const baasRESTApiRefreshToken = 'ndKBay9/Fe5gHOJUCvQITC4vi+VilgOdin4X8mSdQfCn4qYy+dHP4nTGPKoSsZH6zfb8zYy8n2s0DDLLyjAaL89CyZxSiKVW1gsu92WKUl02G14lQwTESB8EHLZleb8Ip7dsdFQuGV4oe1WYoNtnir48PMwNVyHTbVqG2j+Pc6uOxz1jG887dgHBt0OwdVC/DGD69DvGdgnWGeCO61OXUPz93WbruoQjsUdjWSbA4OJgPxNfWctNdcsuLhUNZRJLJX4VqYY+boJTUnNgj192kp+6hE2P6YiQBGk+HtYkNHwDVg==';
// 区块链网络名称（Fabric 通道名称）
const baasNetworkName = 'channel3'
// 区块链智能合约名称（Fabric 链码名称）
const baasContractName = 'notary'
// ----

const client = axios.create({
    baseURL: baasRESTApiEndpoint + '/api/v1/networks/' + encodeURIComponent(baasNetworkName),
});

const getOwnerCredentials = oauth.client(axios.create(), {
    url: baasRESTApiEndpoint + '/api/v1/token',
    grant_type: 'refresh_token',
    refresh_token: baasRESTApiRefreshToken,
});

client.interceptors.request.use(
    // Wraps axios-token-interceptor with oauth-specific configuration,
    // fetches the token using the desired claim method, and caches
    // until the token expires
    oauth.interceptor(tokenProvider, getOwnerCredentials)
);

async function saveToChain(key, value) {
    // 当内容检测结果为 block 或 review 时，终止上链
    await client.post('transactions/invoke?content_check=block,review', {
        chaincode: baasContractName,
        args: ['set', key, value],
        transients: {
            key: 'value',
        },
    })
        .then(function (response) {
            if (response.data.Success && response.data.Result.status === '200') {
                var txId = response.data.Result.id;
                console.info('Data ' + key + ' pushed to blockchain with transaction ' + txId);
            } else if ( response.data.Error.code === 1180418 ) {
                console.info("Content Moderation check failed: " + response.data.Error.message + ". key: " + key + ", value: " + value);
            } else {
                console.error("Code: " + response.data.Error.code + "\nRequestID: " + response.data.Error.request_id + "\nMessage: " + response.data.Error.message);
            }
        })
        .catch(function (error) {
            console.error(error);
        });
}

async function main() {
    // 正常数据调用 Fabric 智能合约上链
    const key = ''+Date.now();
    await saveToChain(key, key);

    // 拒绝辱骂、政治敏感等非法数据
    await saveToChain(key, '1月16日,在对缅甸联邦共和国进行国事访问前夕,国家主席习近平在缅甸')
}

main();








