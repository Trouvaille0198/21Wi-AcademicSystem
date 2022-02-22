const cloud = require('wx-server-sdk')
const axios = require('axios')


cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    var ret = await axios({
        method: 'delete',
        url: 'http://1.15.130.83:8080/api/v1/selection',
        headers: {
            'Content-Type': 'application/json'
        },
        data: {
            "courseID":Number(event.courseID),
            "studentID":Number(event.studentID)},
            "score": 0
    })
    return ret.data
}