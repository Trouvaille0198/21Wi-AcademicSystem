package routers

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() // 生成了一个WSGI应用程序实例
	r.Use(Recovery)

	//v1 := r.Group("api/v1")
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
