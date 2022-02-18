// pages/student/selectcourseresult/selectcourseresult.js
let app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        activeNames: [],
        scr: {},
        selectedbox:[],
        student:{}
    },
    checkboxchange:function(e){
        let sc = this.data.scr
        const value = e.detail.value
        // console.log(sc)
        // console.log(value)
        for (let i = 0, lenI = sc.length; i < lenI; ++i) {
            for (let j = 0, lenJ = value.length; j < lenJ; ++j) {
              if (sc[i].ID === Number(value[j])) {
                sc[i].checked = true
                console.log(sc)
                break
              }
            }
        }
        this.setData({
            scr: sc
        })
    },
    selectcourse:function(e){
        const sc = this.data.scr
        for (let i = 0, lenI = sc.length; i < lenI; ++i) {
            if(sc[i].checked == true)
            wx.cloud.callFunction({
                name: "student",
                data: {
                    type: 'selectcourse',
                    studentID: student.ID,
                    courseID:sc[i].ID
                }
            }).then(res => {
                console.log(res.result)
            })
        }
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let scr = app.globalData.select_course_result
        for (let i = 0, lenI = scr.length; i < lenI; ++i) {
            scr[i].checked=false
        }
        this.setData({
            scr: scr
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