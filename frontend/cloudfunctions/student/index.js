// 云函数入口文件
const cloud = require('wx-server-sdk')
const login = require('./login/index');
const getcourse = require('./getcourse/index');
const selectcourse = require('./selectcourse/index');
const getScourse = require('./getScourse/index');
const deletecourse = require('./deletecourse/index');
cloud.init()

// 云函数入口函数
exports.main = async (event, context) => {
    switch (event.type) {
        case 'login':
            return await login.main(event, context);
        case 'getcourse':
            return await getcourse.main(event, context);
        case 'getScourse':
            return await getScourse.main(event, context);   
        case 'selectcourse':
            return await selectcourse.main(event, context);
        case 'deletecourse':
            return await deletecourse.main(event, context);
    }
}