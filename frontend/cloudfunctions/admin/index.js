// 云函数入口文件
const cloud = require('wx-server-sdk')
const login = require('./login/index');
const addcourse = require('./addcourse/index');
const changescore = require('./changescore/index');
const getScourse = require('./getScourse/index');
const getStudentID = require('./getStudentID/index');
cloud.init()

// 云函数入口函数
exports.main = async (event, context) => {
    switch (event.type) {
        case 'login':
            return await login.main(event, context);
        case 'addcourse':
            return await addcourse.main(event, context);
        case 'changescore':
            return await changescore.main(event, context);  
        case 'getScourse':
            return await getScourse.main(event,context); 
        case 'getStudentID':
            return await getStudentID.main(event,context);
    }
}