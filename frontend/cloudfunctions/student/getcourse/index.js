const cloud = require('wx-server-sdk')
const axios = require('axios')
var FormData = require('form-data');


cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    url = 'http://1.15.130.83:8080/api/v1/course?'
    if(event.credit != 0){
        url = url + `credit=${event.credit}&`
    }
    if(event.department != ''){
        url = url + `department=${event.department}&`
    }
    if(event.name != ''){
        url = url + `name=${event.name}&`
    }
    if(event.number != ''){
        url = url + `number=${event.number}&`
    }
    if(event.teacher_name != ''){
        url = url + `teacher_name=${event.teacher_name}&`
    }
    if(event.term != ''){
        url = url + `term=${event.term}&`
    }

    var ret = await axios({
        method: 'get',
        url: url
    })

    return ret.data
}