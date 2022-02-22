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
// @Success      200  {object}  []model.CourseByStuResponse
// @Router       /student/{id}/course [get]
func GetCoursesByStudent(c *gin.Context) {
	hasScore, ok := c.GetQuery("hasScore")
	studentID, _ := util.String2Int(c.Param("id"))
	var courses *[]model.CourseByStuResponse
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

// UpdateWholeCourse godoc
// @Summary      整体更新课程信息
// @Description  整体更新课程信息
// @Tags         course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "course ID"
// @Param 		 course   body   model.Course   true   "course 实例"
// @Success      200  {string} string
// @Router       /course/{id} [put]
func UpdateWholeCourse(c *gin.Context) {
	courseID, _ := util.String2Int(c.Param("id"))
	course := model.Course{}
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := model.UpdateWholeCourse(courseID, course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// GetCourseByID godoc
// @Summary      根据id获取课程信息
// @Description  根据id获取课程信息
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "course ID"
// @Success      200  {object}  model.Course
// @Router       /course/{id} [get]
func GetCourseByID(c *gin.Context) {
	courseID, _ := util.String2Int(c.Param("id"))

	course, err := model.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"student": course,
	})
}

// CreateCourse godoc
// @Summary      创建课程
// @Description  创建课程
// @Tags         course
// @Param 		 course   body   model.Course   true   "course 实例"
// @Success      200  {string} string
// @Router       /course [post]
func CreateCourse(c *gin.Context) {
	course := model.Course{}
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	newCourse, err := model.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"course":  newCourse,
	})
}

// GetCoursesByAttrs godoc
// @Summary      获取课程
// @Description  获取课程 可以自由添加筛选属性
// @Tags         course
// @Param 		 course   query   model.Course   false   "course 实例"
// @Success      200  {string} string
// @Router       /course [get]
func GetCoursesByAttrs(c *gin.Context) {
	course := model.Course{}
	if err := c.ShouldBindQuery(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	coursesResult, err := model.GetCourseByAttrs(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"course":  coursesResult,
	})
}
