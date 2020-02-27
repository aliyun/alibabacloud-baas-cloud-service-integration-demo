const axios = require('axios');
const crypto = require('crypto');

// ---- 根据您的实例信息配置以下参数
// 钉钉消息推送地址
const dingtalkCallbackUrl = "https://oapi.dingtalk.com/robot/send?access_token=aa814fe2f6642f9327b92dbe0b3c32b32b68d85c19afab02b3813e44613aed90";
// 钉钉消息推送签名密钥
const dingtalkSignKey = "SECb72495657da89e260ddc128b0ed7117a6923616eeda9a6f7ff5b32fd3569a873";
// 钉钉 Outgoing 消息验证密钥
const dingtalkVerifyKey = "";
// ----


var client;

// exports.initializer: initializer function
exports.initializer = function (context, callback) {
    client = axios.create();
    callback(null, '');
};

function newDingtalkSignature(timestamp) {
    var oriString = timestamp + "\n" + dingtalkSignKey;
    var hash = crypto.createHmac('sha256', dingtalkSignKey)
        .update(oriString, 'utf8')
        .digest('base64');
    return encodeURIComponent(hash);
}

async function sendDingtalkMessage(name, task) {
    var timestamp = Date.now();
    var signature = newDingtalkSignature(timestamp);

    //发送钉钉通知
    await client.post(dingtalkCallbackUrl+'&timestamp='+timestamp+'&sign='+signature, {
        msgtype: 'text',
        text: {
            content: '[事件] 任务流 ' + task.name + ' 触发了事件 ' + name + ', 请留意'
        }
    })
        .then(function (response) {
            if (response.data.errcode != 0) {
                throw new Error(JSON.stringify(response.data));
            }
        });
    return true
}

exports.handler = function (event, context, callback) {
    // 从 evnet 中解析出 Event
    try {
        var data = JSON.parse(event);
        if (data.name === 'event-task-finished' || data.name === 'event-create-task' || data.name === 'event-approve-task') {
            var buff = new Buffer(data.content, 'base64');
            var content = buff.toString('utf8');
            console.log(content);
            var task = JSON.parse(content);
            console.log(task);
            sendDingtalkMessage(data.name, task).then((result) => {
                if (result) {
                    callback(null, "");
                } else {
                    callback(new Error('failed'));
                }
            }).catch(function (error) {
                callback(error);
            });
        } else {
            callback(null, "SKIP");
        }
    } catch (e) {
        callback(new Error('Parse event failed: ' + e));
        return;
    }
};






