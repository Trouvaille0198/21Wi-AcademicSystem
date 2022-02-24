// pages/student/searchscore/searchscore.js
let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        array: ['22-春季学期','22-冬季学期','21-秋季学期','21-冬季学期','21-春季学期'],
        index: 0,
        class: [],
        student: {},
        selectedclass: [],
    },
    onChange(event) {
        this.setData({
            activeNames: event.detail,
        });
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let student = app.globalData.student
        var that = this
        wx.cloud.callFunction({
            name: "student",
            data: {
                type: 'getScourse',
                ID: student.ID,
                hasScore: true,
            }
        }).then(res => {
            that.setData({
                student: student,
                class: res.result.courses
            })
        })
        var class1 = this.data.class
        var array = this.data.array
        var selectedclass = []
        for(let i=0;i<class1.length;++i){
            if(class1[i].term == array[index])
                selectedclass.push(class1[i])
        }
        console.log(selectedclass)
        this.setData({
            selectedclass:selectedclass
        })
    },
    bindPickerChange:function(e){
        var class1 = this.data.class
        var array = this.data.array
        var selectedclass = []
        for(let i=0;i<class1.length;++i){
            if(class1[i].term == array[e.detail.value])
                selectedclass.push(class1[i])
        }
        console.log(selectedclass)
        this.setData({
            index: e.detail.value,
            selectedclass:selectedclass
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