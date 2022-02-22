package v1

import (
	"academic-system/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginAsStu godoc
// @Summary      以学生身份登录
// @Description  以学生身份登录
// @Tags         login
// @Accept       mpfd
// @Produce      json
// @Param        number   formData   string  true  "学号"
// @Param        password   formData   string  true  "密码"
// @Success      200  {object}  model.Student
// @Router       /login/student [post]
func LoginAsStu(c *gin.Context) {
	number := c.PostForm("number")
	password := c.PostForm("password")

	student, err := model.GetStudentByNumber(number)
	if err != nil || student.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "找不到此账号",
		})
		return
	}
	// 简单地验证一下密码
	if student.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})

}

// LoginAsAdmin godoc
// @Summary      以管理员身份登录
// @Description  以管理员身份登录
// @Tags         login
// @Param        number   formData   string  true  "账号"
// @Param        password   formData   string  true  "密码"
// @Accept       mpfd
// @Produce      json
// @Success      200
// @Router       /login/admin [post]
func LoginAsAdmin(c *gin.Context) {
	number := c.PostForm("number")
	password := c.PostForm("password")

	// 简单地验证一下账号和密码
	if number != "123" || password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "账号密码错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})

}
