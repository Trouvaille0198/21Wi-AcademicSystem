// pages/enter/enter.js
const axios = require('axios')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        array:['学生','admin'],
        index:0
    },
    
    bindPickerChange: function(e) {
        console.log('picker发送选择改变，携带值为', e.detail.value)
        this.setData({
          index: e.detail.value
        })
    },

    login:function(e){
       wx.cloud.callFunction({
           name:"student",
           data:{
               type:'login',
               number:'0198',
               password:"123"
           }
       }).then(res=>{
        console.log("res")
           console.log(res)
       }).catch(error=>{
        console.log("err")
           console.log(error)
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