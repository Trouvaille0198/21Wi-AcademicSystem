const cloud = require('wx-server-sdk')
const axios = require('axios')
var FormData = require('form-data');

cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    var data = new FormData()
    data.append('score',event.score)
    url = 'http://1.15.130.83:8080/api/v1/student/'+event.studentID.toString()+"/course/"+event.courseID.toString()
    console.log(url);
    var ret = await axios({
        method: 'put',
        url: url,
        data
    })

    return ret.data
}