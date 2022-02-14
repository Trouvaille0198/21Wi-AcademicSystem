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
		// 学生
		rv1.GET("/student", v1.GetAllStudents)
		rv1.GET("/student/:id", v1.GetStudentByID)
		rv1.GET("/student/:id/course", v1.GetCoursesByStudent)
	}
	//{
	//	// 用户操作
	//	v1.POST("user/register", api.UserRegister)
	//	v1.POST("user/login", api.UserLogin)
	//	authed := v1.Group("/") //需要登陆保护
	//	{
	//		//任务操作
	//		authed.GET("tasks", api.ListTasks)
	//		authed.POST("task", api.CreateTask)
	//		authed.GET("task/:id", api.ShowTask)
	//		authed.DELETE("task/:id", api.DeleteTask)
	//		authed.PUT("task/:id", api.UpdateTask)
	//		authed.POST("search", api.SearchTasks)
	//	}
	//}
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
