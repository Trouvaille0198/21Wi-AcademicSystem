const cloud = require('wx-server-sdk')
const axios = require('axios')


axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 400:
                    return Promise.resolve(error.response)
                case 404:
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
    var ret = await axios({
        method: 'post',
        url: 'http://1.15.130.83:8080/api/v1/selection',
        headers: {
            'Content-Type': 'application/json'
        },
        data: {
            "courseID":Number(event.courseID),
            "studentID":Number(event.studentID)}
    })
    return ret.data
}