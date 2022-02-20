const cloud = require('wx-server-sdk')
const axios = require('axios')


cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    var ret = await axios({
        method: 'post',
        url: 'http://1.15.130.83:8080/api/v1/course',
        headers: {
            'Content-Type': 'application/json'
        },
        data: {
            "credit":Number(event.credit),
            "department":event.department,
            "number":event.number,
            "teacher_name":event.teacher_name,
            "term":event.term,
            "name":event.name
        }
    })
    return ret.data
}