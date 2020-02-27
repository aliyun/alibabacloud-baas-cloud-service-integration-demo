// exports.initializer: initializer function
exports.initializer = function (context, callback) {
    callback(null, '');
};

exports.handler = function (event, context, callback) {
    console.info(''+event);
    callback(null, "OK");
};

