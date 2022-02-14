package v1

import (
	"academic-system/model"
	"academic-system/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateSelectionScore godoc
// @Summary      更新课程成绩
// @Description  更新课程成绩
// @Tags         selection
// @Param 		 student_id   path   int   true   "student ID"
// @Param 		 course_id   path   int   true   "course ID"
// @Param 		 score   formData   int   true   "score"
// @Success      200  {string} string
// @Router       /student/{student_id}/course/{course_id} [put]
func UpdateSelectionScore(c *gin.Context) {
	studentID := util.String2Int(c.Param("student_id"))
	courseID := util.String2Int(c.Param("course_id"))
	score := util.String2Int(c.PostForm("score"))

	selection, err := model.GetSelection(studentID, courseID)
	if err != nil || selection.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "找不到选课记录",
		})
		return
	}

	err = model.UpdateSelectionScore(selection.ID, score)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}
