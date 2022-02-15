package router

import (
	_ "academic-system/docs"
	v1 "academic-system/router/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() // 生成了一个WSGI应用程序实例

	// 跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true 万事大吉
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))
	r.Use(Recovery)

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	rv1 := r.Group("/api/v1")
	{
		// 功能性接口
		rv1.GET("test/ping", v1.Ping)
		rv1.GET("test/generate-examples", v1.GenerateExamples)
		// 登录
		rv1.POST("/login/student", v1.LoginAsStu)
		rv1.POST("/login/admin", v1.LoginAsAdmin)
		// 学生
		rv1.GET("/student", v1.GetAllStudents)
		rv1.GET("/student/:id", v1.GetStudentByID)
		rv1.GET("/student/:id/course", v1.GetCoursesByStudent)
		rv1.POST("/student", v1.CreateStudent)
		rv1.PUT("/student/:student_id/course/:course_id", v1.UpdateSelectionScore)
		// 课程
		rv1.POST("/course", v1.CreateCourse)
		rv1.PUT("/course/:id", v1.UpdateWholeCourse)
		rv1.GET("course", v1.GetCoursesByAttrs)
		// 选课
		rv1.POST("/selection", v1.CreateSelection)
	}
	return r
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Error("gin catch error: ", r)
			c.JSON(http.StatusInternalServerError, "系统内部错误")
		}
	}()
	c.Next()
}
