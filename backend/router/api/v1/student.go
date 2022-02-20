package v1

import (
	"academic-system/model"
	"academic-system/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStudentByID godoc
// @Summary      根据id获取学生信息
// @Description  根据id获取学生信息
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "student ID"
// @Success      200  {object}  model.Student
// @Router       /student/{id} [get]
func GetStudentByID(c *gin.Context) {
	studentID := util.String2Int(c.Param("id"))

	student, err := model.GetStudentByID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

// GetStudentsByAttrs godoc
// @Summary      获取学生
// @Description  获取学生 可以自由添加筛选属性
// @Tags         student
// @Param 		 student   query   model.Student   false   "student 实例"
// @Success      200  {string} string
// @Router       /student [get]
func GetStudentsByAttrs(c *gin.Context) {
	student := model.Student{}
	if err := c.ShouldBindQuery(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	studentsResult, err := model.GetStudentByAttrs(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"student": studentsResult,
	})
}

// CreateStudent godoc
// @Summary      创建学生
// @Description  创建学生
// @Tags         student
// @Param 		 student   body   model.Student    true   "student 实例"
// @Success      200  {string} string
// @Router       /student [post]
func CreateStudent(c *gin.Context) {
	student := model.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		if err.Error() == "UNIQUE constraint failed: students.number" {
			c.JSON(http.StatusConflict, gin.H{
				"message": "学号已存在！",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	newStudent, err := model.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"student": newStudent,
	})
}
