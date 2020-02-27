const axios = require('axios');
const oauth = require('axios-oauth-client');
const tokenProvider = require('axios-token-interceptor');

// ---- 根据您的实例信息配置以下参数
// 区块链 REST API 服务地址
const baasRESTApiEndpoint = 'http://your.gateway.endpoint';
// 在区块链控制台生成的 refresh token
const baasRESTApiRefreshToken = 'yO+zDBVJDCnX2DL2Q4yubX1VZfkLD+cwje1lldgtU1UOtWNsn344n/oMF7pWYLM87f+F/4kssyaRIKZYKfl5ha4EXdJpXx1sVLhXE/evMIsYpcn58ZnJkwg+axCKK+50Sp2dqu77HKCj27g+qDI9vRz9jbvajIGrYLmXQ/wN4vMYKCESqWIH3owkM9tE8+1/Di95zpMR0knhpMGhAqTlKmi4jJ6SE7RB+YCRxv+HZUzPirZoXe6laxYGfWlOYmd9jzyXjG2CD6biFMMa5umOjE6jfn2djoykE9Hgx29YG/BFgA==';
// 区块链网络名称（Fabric 通道名称）
const baasNetworkName = 'channel3'
// 区块链智能合约名称（Fabric 链码名称）
const baasContractname = 'sacc'
// ----

var client;

// exports.initializer: initializer function
exports.initializer = function (context, callback) {
    client = axios.create({
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

    callback(null, '');
}

exports.handler = function (event, context, callback) {
    // 从 evnet 中解析出 MNS 消息
    var mnsMessage, mnsMessageId;
    try {
        var data = JSON.parse(event);
        mnsMessage = data.Message;
        mnsMessageId = data.MessageId;
    } catch (e) {
        callback(new Error('Parse event failed: ' + e));
        return;
    }

    // 处理 MNS 消息
    mnsMessage = mnsMessage;

    // 调用 Fabric 智能合约, 将MNS消息上链；根据业务场景，修改调用智能合约的方式和参数
    client.post('transactions/invoke', {
        chaincode: baasContractname,
        args: ['set', ''+Date.now(), mnsMessage],
        transients: {
            key: 'value',
        },
    })
    .then(function (response) {
        if (response.data && response.data.Success && response.data.Result.status === '200') {
            var txId = response.data.Result.id;
            console.info('MNS message ' + mnsMessageId + ' pushed to blockchain with transaction ' + txId);
            callback(null, txId);
        } else {
            callback(new Error("Code: " + response.data.Error.code + "\nRequestID: " + response.data.Error.request_id + "\nMessage: " + response.data.Error.message));
        }
    })
    .catch(function (error) {
        console.error(error);
        callback(new Error(error));
    });
};








