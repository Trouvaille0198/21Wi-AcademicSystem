const cloud = require('wx-server-sdk')
const axios = require('axios')


cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    var ret = await axios({
        method: 'post',
        url: 'http://1.15.130.83:8080/api/v1/selection',
        data: `{"courseID":"${event.courseID}","studentID":"${event.studentID}"}`
    })
    console.log(ret.data)

    return ret.data
}