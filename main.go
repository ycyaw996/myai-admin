package main

import (
	"HG-Dashboard/routes"
	"HG-Dashboard/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化数据库
	utils.InitDB()

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务器
	err := router.Run()
	if err != nil {
		return
	}
}
