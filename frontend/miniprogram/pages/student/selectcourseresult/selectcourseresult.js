// pages/student/selectcourseresult/selectcourseresult.js
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
    },
    checkboxchange: function (e) {
        const value = e.detail.value
        const student = this.data.student
        console.log(student)
        console.log(value)

        this.setData({
            coursetobeselected: value
        })
    },

    selectcourse: function (e) {
        const ctbs = this.data.coursetobeselected
        const student = this.data.student
        for (let i = 0, lenI = ctbs.length; i < lenI; ++i) {
            wx.cloud.callFunction({
                name: "student",
                data: {
                    type: 'selectcourse',
                    studentID: student.ID,
                    courseID: Number(ctbs[i])
                }
            }).then(res => {
                console.log("选课id和学生")
                console.log(student.ID)
                console.log(ctbs[i])
                console.log(res.result)
            })
        }
        var msg = '课号:'
        for (let i = 0; i < ctbs.length - 1; ++i)
            msg = msg + ctbs[i] + ','
        msg = msg + ctbs[ctbs.length - 1]
        Dialog.alert({
            title: '选课成功！',
            message: msg ,
            showCancelButton: true
        }).then(() => {
            // on close
        });
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