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
