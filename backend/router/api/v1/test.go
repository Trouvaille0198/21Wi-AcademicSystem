package v1

import (
	"academic-system/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping godoc
// @Summary      ping!
// @Description  pong me!
// @Tags         test
// @Success      200  {string} string
// @Router       /test/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GenerateExamples godoc
// @Summary      生成样例数据
// @Description  生成样例数据
// @Tags         test
// @Success      200  {string} string
// @Router       /test/generate-examples [get]
func GenerateExamples(c *gin.Context) {
	model.CreateStudentsExample()
	model.CreateCoursesExample()
	model.CreateSelectionsExample()

	c.JSON(http.StatusOK, gin.H{
		"message": "已生成样例数据",
	})
}
