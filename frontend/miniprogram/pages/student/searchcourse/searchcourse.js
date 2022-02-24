// pages/student/searchcourse/searchcourse.js
let app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        activeNames: [],
        student: {},
        class: []
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let student = app.globalData.student
        var that = this
        console.log(student.ID)
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getScourse',
                ID: student.ID,
                hasScore: false,
            }
        }).then(res => {
            console.log(res)
            console.log(student)
            that.setData({
                student: student,
                class: res.result.courses
            })
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