// pages/student/retreatclass/retreatclass.js
import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
let app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        activeNames: [],
        scr: {},
        selectedbox: [],
        student: {},
        coursetobeselected: [],
        class: {},
        value: []
    },
    checkboxchange: function (e) {
        const value = e.detail.value

        console.log(value)
        this.setData({
            coursetobeselected: value
        })
    },

    retreatcourse: function (e) {
        console.log(1)
        const ctbs = this.data.coursetobeselected
        const student = this.data.student
        console.log(ctbs)
        for (let i = 0, lenI = ctbs.length; i < lenI; ++i) {
            console.log(student.ID)
            console.log(Number(ctbs[i]))
            wx.cloud.callFunction({
                name: "student",
                data: {
                    type: 'deletecourse',
                    studentID: student.ID,
                    courseID: Number(ctbs[i])
                }
            }).then(res => {
                // var msg = '课号:'
                // for (let i = 0; i < ctbs.length - 1; ++i)
                //     msg = msg + ctbs[i] + ','
                // msg = msg + ctbs[ctbs.length - 1]
                var msg = res
                console.log(res)
                Dialog.alert({
                    // title: '退课成功！',
                    message: msg.result.message ,
                    showCancelButton: true
                }).then(() => {
                    // on close
                });
            }).catch(error =>{
                console.log(error)
                Dialog.alert({
                    title: error,
                    // message: msg ,
                    showCancelButton: true
                }).then(() => {
                    // on close
                });
            })
        }
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let student = app.globalData.student
        let scr = app.globalData.select_course_result
        for (let i = 0, lenI = scr.length; i < lenI; ++i) {
            scr[i].checked = false
        }
        this.setData({
            scr: scr,
            student: student
        })
        var that = this
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getScourse',
                ID: student.ID,
                hasScore: false,
            }
        }).then(res => {
            console.log(res.result.courses)
            that.setData({
                student: student,
                class: res.result.courses
            })
        })
        console.log(this.data.class)
    },


    onChange(event) {
        this.setData({
            activeNames: event.detail,
        });
    },
    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})