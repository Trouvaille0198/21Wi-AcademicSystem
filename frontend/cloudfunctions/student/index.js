// 云函数入口文件
const cloud = require('wx-server-sdk')
const login = require('./login/index');
cloud.init()

// 云函数入口函数
exports.main = async (event, context) => {
    switch (event.type) {
        case 'login':
            return await login.main(event, context);
    }
}