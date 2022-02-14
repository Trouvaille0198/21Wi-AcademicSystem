package main

import (
	"academic-system/model"
	"academic-system/router"
)

// @title           Academic System
// @version         1.0
// @description     学生选课系统

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	model.Setup() // 初始化gorm数据库连接

	r := router.NewRouter()
	_ = r.Run(":8080")
}
