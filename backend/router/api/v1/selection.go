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
// @Tags         selection, admin
// @Param 		 student_id   path   string   true   "student ID"
// @Param 		 course_id   path   string   true   "course ID"
// @Param 		 score   formData   string   true   "score"
// @Success      200  {string} string
// @Router       /student/{student_id}/course/{course_id} [put]
func UpdateSelectionScore(c *gin.Context) {
	studentID, err1 := util.String2Int(c.Param("student_id"))
	courseID, err2 := util.String2Int(c.Param("course_id"))
	score, err3 := util.String2Int(c.PostForm("score"))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "student_id 不合法！",
		})
		return
	}
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "course_id 不合法！",
		})
		return
	}
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "score 不合法！",
		})
		return
	}

	selection, err := model.GetSelection(studentID, courseID)
	if err != nil || selection.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "找不到选课记录",
		})
		return
	}

	err = model.UpdateSelectionScore(selection.ID, score)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "更新失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// CreateSelection godoc
// @Summary      创建选课
// @Description  创建选课
// @Tags         selection
// @Param 		 selection   body   model.Selection   true   "选课情况"
// @Success      200  {string} string
// @Router       /selection [post]
func CreateSelection(c *gin.Context) {
	var selection model.Selection
	if err := c.ShouldBindJSON(&selection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newSelection, err := model.CreateSelection(selection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "创建失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "创建成功",
		"selection": newSelection,
	})
}

// DeleteSelection godoc
// @Summary      删除选课
// @Description  删除选课
// @Tags         selection
// @Param 		 selection   body   model.Selection   true   "选课情况"
// @Success      200  {string} string
// @Router       /selection [delete]
func DeleteSelection(c *gin.Context) {
	var selection model.Selection
	if err := c.ShouldBindJSON(&selection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := model.DeleteSelection(selection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "删除失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
