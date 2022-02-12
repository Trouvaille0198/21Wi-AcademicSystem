package main

import (
	"academic-system/models"
	"academic-system/routers"
)

func main() {
	models.Setup() // 初始化gorm数据库连接

	r := routers.NewRouter()
	_ = r.Run(":8080")
}
