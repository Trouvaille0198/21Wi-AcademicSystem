package v1

import (
	"academic-system/model"
	"academic-system/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetCoursesByStudent godoc
// @Summary      获取指定学生的所有课程
// @Description  获取指定学生的所有课程
// @Tags         student, course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "student ID"
// @Param        hasScore query bool false "是否有成绩 不写即全部返回"
// @Success      200  {object}  []model.CourseByStuResult
// @Router       /student/{id}/course [get]
func GetCoursesByStudent(c *gin.Context) {
	hasScore, ok := c.GetQuery("hasScore")
	studentID := util.String2Int(c.Param("id"))
	var courses *[]model.CourseByStuResult
	var sqlErr error
	if !ok {
		courses, sqlErr = model.GetCoursesByStudent(studentID)
	} else {
		hasScore, err := strconv.ParseBool(hasScore)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if hasScore {
			courses, sqlErr = model.GetCoursesByStudentWithScore(studentID)
		} else {
			courses, sqlErr = model.GetCoursesByStudentWithoutScore(studentID)
		}
	}

	if sqlErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": sqlErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}
