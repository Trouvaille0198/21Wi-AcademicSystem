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
    // var ret = await axios({
    //     method: 'delete',
    //     url: 'http://1.15.130.83:8080/api/v1/selection',
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    //     config:{
    //         data: {
    //             "courseID":3,
    //             "studentID":3,
    //             "score": 10
    //         }
    //     }
    // })
    var ret = await axios.delete('http://1.15.130.83:8080/api/v1/selection', {
        data: {
            "courseID": 3,
            "studentID": 3,
            "score": 10
        }
    })
    return ret.data
}