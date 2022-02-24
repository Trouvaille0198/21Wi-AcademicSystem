// pages/student/selectcourse/selectcourse.js
let app = getApp()
import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';

Page({

    /**
     * 页面的初始数据
     */
    data: {
        student: {},
        credit: 0,
        department: '',
        name: '',
        number: '',
        teacher_name: '',
        term: ''
    },

    /**
     * 生命周期函数--监听页面加载
     */
    ChangeCredit:function(e){
        console.log(typeof(Number(e.detail)))
        this.setData({
            credit: Number(e.detail)
        })
    },
    ChangeDepartment:function(e){
        console.log(e.detail)
        this.setData({
            department: e.detail
        })
    },
    ChangeName:function(e){
        console.log(e.detail.value)
        this.setData({
            name: e.detail.value
        })
    },
    ChangeNumber:function(e){
        console.log(e.detail)
        this.setData({
            number: e.detail
        })
    },
    ChangeTeacher_name:function(e){
        console.log(e.detail)
        this.setData({
            teacher_name: e.detail
        })
    },
    ChangeTerm:function(e){
        console.log(e.detail)
        this.setData({
            term: e.detail
        })
    },
    onLoad: function (options) {
        this.setData({
            student: app.globalData.student
        })
    },
    GetCourse:function(e){
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getcourse',
                credit: this.data.credit,
                department: this.data.department,
                name: this.data.name,
                number: this.data.number,
                teacher_name: this.data.teacher_name,
                term: this.data.term
            }
        }).then(res => {
            app.globalData.select_course_result = res.result.course
            wx.navigateTo({
              url: '../selectcourseresult/selectcourseresult',
            })
            console.log(res.result.course)
        }).catch(error => {
            console.log(error)
            // Dialog.alert({
            //     message: '出错了',
            //     showCancelButton: true
            // }).then(() => {
            //     // on close
            // })
            // .catch(()=>{
            //     // on cancel
            // });
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