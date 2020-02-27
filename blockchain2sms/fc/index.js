const SMSClient = require('@alicloud/sms-sdk');

// ---- 根据您的实例信息配置以下参数
const accessKeyId = "<Your AK>";
const secretAccessKey = "<Your SK>";
// ----

var client;

// exports.initializer: initializer function
exports.initializer = function (context, callback) {
    client = new SMSClient({accessKeyId, secretAccessKey})

    callback(null, '');
};

// 通过外部系统，获取用户的电话，发送短信通知
async function getUserPhoneFromDB(user) {
    return "151********";
}

async function sendSMSMessage(task) {
    const phone = await getUserPhoneFromDB(task.creator);

    //发送短信, vpc配置options={method:'POST'}，改为POST请求
    // 请根据SMS服务的签名名称及模版Code，修改以下调用参数
    await client.sendSMS({
        PhoneNumbers: phone,
        SignName: 'Octopus',
        TemplateCode: 'SMS_000000',
        TemplateParam: '{"name":"['+task.name+']"}'
    }).then(function (res) {
        let {Code}=res
        if (Code === 'OK') {
            //处理返回参数
            console.log(res)
        }
    }, function (err) {
        console.error(err)
    })
    return true
}

exports.handler = function (event, context, callback) {
    // 从 evnet 中解析出 Event
    try {
        var data = JSON.parse(event);
        if (data.name === 'event-task-finished') {
            var buff = new Buffer(data.content, 'base64');
            var content = buff.toString('utf8');
            console.log(content);
            var task = JSON.parse(content);
            console.log(task);
            sendSMSMessage(task).then((result) => {
                if (result) {
                    callback(null, "");
                } else {
                    callback(new Error('failed'));
                }
            })
        } else {
            callback(null, "SKIP");
        }
    } catch (e) {
        callback(new Error('Parse event failed: ' + e));
        return;
    }
};









