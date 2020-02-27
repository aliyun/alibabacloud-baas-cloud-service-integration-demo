const axios = require('axios');
const oauth = require('axios-oauth-client');
const tokenProvider = require('axios-token-interceptor');

// ---- 根据您的实例信息配置以下参数
// 区块链 REST API 服务地址
const baasRESTApiEndpoint = 'http://your.gateway.endpoint';
// 在区块链控制台生成的 refresh token
const baasRESTApiRefreshToken = 'PAr7YkKH2pUiLHMA526h0Ar5MZHXheRw/fnzAUv1cz1PJBXZOQBQkirnDyLxK9kV3Q8cWRfRkerAUONQNUOD/IcBEQ2kUxrXXt3+550ZzchQFdNk1D5+qKwN5K8+tL1zTmNIgUvnQffpZ7lyKpnpMusAJqKZJqbX8dEHH3ye9o3ZKkPnDm2K5vbliAlI1eUrlytrd1HIdBBVU+uyHyL6xkaq+SQ7xDhtgSL7Gj367LFhb76VqU29ePWt6feLY1uo7YyvPRspxFf5ev9aEneGTzUx8cm/DGRqkzo=';
// 区块链网络名称（Fabric 通道名称）
const baasNetworkName = 'channel3'
// 区块链智能合约名称（Fabric 链码名称）
const baasContractName = 'taskmgr'
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

async function callChainCode(args) {
    await client.post('transactions/invoke', {
        chaincode: baasContractName,
        args: args,
        transients: {
            key: 'value',
        },
    })
        .then(function (response) {
            if (response.data.Success && response.data.Result.status === '200') {
                var txId = response.data.Result.id;
                console.info('Data pushed to blockchain with transaction ' + txId);
                console.info(response.data.Result);
            } else {
                console.error("Code: " + response.data.Error.code + "\nRequestID: " + response.data.Error.request_id + "\nMessage: " + response.data.Error.message);
            }
        })
        .catch(function (error) {
            console.error(error);
        });
}

async function main() {
    // 创建审批流
    const taskName = 'task-'+Date.now();
    await callChainCode([
        'create',
        taskName,
        `{
            "requires": ["e2ehmfqasthMSP.octopus_26842_12345678901234"],
            "description": "示例任务，requires 配置审批任务完成需要那些用户同意。用户描述为 '组织MSP.用户名称'"
        }`,
    ]);

    // 同意审批
    await callChainCode([
        'approve',
        taskName,
    ]);

    // 读取任务
    await callChainCode([
        'get',
        taskName,
    ]);
}

main();








