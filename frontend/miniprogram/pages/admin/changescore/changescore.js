// pages/admin/changescore/changescore.js
let app = getApp()
import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
Page({

    /**
     * 页面的初始数据
     */
    data: {
        studentID: '',
        ID: '',
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
        this.setData({
            ID: e.detail
        })
        console.log(this.data.ID)

        wx.cloud.callFunction({
            name: "admin",
            data: {
                type: 'getStudentID',
                ID: this.data.ID,
            }
        }).then(res => {
            console.log(res.result[0].ID)
            that.setData({
                //student:student,
                studentID: res.result[0].ID
            })
        });
        console.log(that.data.studentID);
    },


    searchCourse:function(e){
        console.log(1)
        console.log(this.data)
        var that = this;
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getScourse',
                ID: this.data.studentID,
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
        if(Number(this.data.score) > 100 || Number(this.data.score)<0)
        {
            Dialog.alert({
                message: "成绩不合法！",
                showCancelButton: true
            }).then(() => {
                // on close
                this.ChangeStudentID();
            })
            .catch(() => {
                // on cancel
            });
            return
        }
        wx.cloud.callFunction({
            name: "admin",
            data: {
                type: 'changescore',
                score: this.data.score,
                studentID: this.data.studentID,
                courseID: this.data.courseID
            }
        }).then(res => {
            console.log(res.result.message)
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
                    message: res.result.message,
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