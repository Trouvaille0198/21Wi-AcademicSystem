const cloud = require('wx-server-sdk')
const axios = require('axios')
var FormData = require('form-data');

axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 409:
                    return Promise.resolve(error.response)
                case 500:
                    return Promise.resolve(error.response)
            }
        }
        return Promise.reject(error.response) // 返回接口返回的错误信息
    })

cloud.init({
    "env": "cloud1-8gezryve5b08e376"
})

// 云函数入口函数
exports.main = async (event, context) => {
    console.log("1111")
    var data = new FormData()
    data.append('number', event.number)
    data.append('password', event.password)
    var ret = await axios({
        method: 'post',
        url: 'http://1.15.130.83:8080/api/v1/login/student',
        headers: data.getHeaders(),
        data
    })
    console.log(ret.data)

    return ret.data
}