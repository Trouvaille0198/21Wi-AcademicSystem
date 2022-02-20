// pages/admin/changescore/changescore.js
let app = getApp()
import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
Page({

    /**
     * 页面的初始数据
     */
    data: {
        studentID: '',
        courseID: '',
        score: 0,
        class: [],
        student: {},
        selectedclass: [],
    },
    onChange(event) {
        this.setData({
            activeNames: event.detail,
        });
    },
    onChange(event) {
        this.setData({
            activeNames: event.detail,
        });
    },
    /**
     * 生命周期函数--监听页面加载
     */
    ChangeStudentID: function (e) {
        var that = this;
        console.log(e.detail)
        that.setData({
            studentID: e.detail
        })
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getScourse',
                ID: that.data.studentID,
                hasScore: null,
            }
        }).then(res => {
            console.log(res.result)
            that.setData({
                //student:student,
                class: res.result.courses
            })
        })


    },
    ChangeCourseID: function (e) {
        console.log(e.detail)
        this.setData({
            courseID: e.detail
        })
    },
    ChangeScore: function (e) {
        console.log(e.detail)
        this.setData({
            score: e.detail
        })
    },
    onLoad: function (options) {

    },
    changescore: function (e) {
        console.log(this.data)
        wx.cloud.callFunction({
            name: "admin",
            data: {
                type: 'changescore',
                score: this.data.score,
                studentID: this.data.studentID,
                courseID: this.data.courseID
            }
        }).then(res => {
            // app.globalData.select_course_result = res.result.course
            // wx.navigateTo({
            //  url: '../selectcourseresult/selectcourseresult',
            //  })
            // console.log(res.result.course)

            console.log(res.result)
            Dialog.alert({
                    context: this,
                    message: res.result.message,
                    showCancelButton: true
                }).then(() => {
                    // on close
                    this.ChangeStudentID();
                })
                .catch(() => {
                    // on cancel
                });
        }).catch(error => {
            console.log(error)
            Dialog.alert({
                    context: this,
                    message: error,
                    showCancelButton: true
                }).then(() => {
                    // on close
                })
                .catch(() => {
                    // on cancel
                });
        })
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