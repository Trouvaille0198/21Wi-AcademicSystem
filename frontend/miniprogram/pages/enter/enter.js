// pages/enter/enter.js
import Dialog from '../miniprogram_npm/@vant/weapp/dialog/dialog';
let app = getApp()

Page({
    /**
     * 页面的初始数据
     */
    data: {
        array: ['学生', 'admin'],
        index: 0,
        name: '',
        pwd: '',
    },

    InputName: function (e) {
        this.setData({
            name: e.detail.value
        })
    },
    InputPwd: function (e) {
        this.setData({
            pwd: e.detail.value
        })
    },

    bindPickerChange: function (e) {
        console.log('picker发送选择改变，携带值为', e.detail.value)
        this.setData({
            index: e.detail.value
        })
    },

    login: function (e) {
        if (this.data.index == 0)
            var name = "student";
        else
            var name = "admin"
        wx.cloud.callFunction({
            name: name,
            data: {
                type: 'login',
                number: this.data.name,
                password: this.data.pwd
            }
        }).then(res => {
            app.globalData.student = res.result.student
            console.log(app.globalData.student)
            wx.redirectTo({
              url: '../studenthome/studenthome',
            })
        }).catch(error => {
            console.log(error)
            Dialog.alert({
                message: '账号或密码错误！',
                showCancelButton: true
            }).then(() => {
                // on close
            })
            .catch(()=>{
                // on cancel
            });
        })

    },



    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

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