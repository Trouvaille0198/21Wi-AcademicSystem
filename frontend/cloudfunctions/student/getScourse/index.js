const cloud = require('wx-server-sdk')
const axios = require('axios')


cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    url = 'http://1.15.130.83:8080/api/v1/student/'+event.ID.toString()+"/course"
    if(event.hasScore == false)
        url = url + "?hasScore=false"
    else if(event.hasScore == true)
        url = url + "?hasScore=true"
    else (event.hasScore == null)
        url = url
    var ret = await axios({
        method: 'get',
        url: url
    })

    return ret.data
}