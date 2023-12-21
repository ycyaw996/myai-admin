package main

import (
	"HG-Dashboard/routes"
	"HG-Dashboard/utils"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化数据库
	configPath := flag.String("config", "./config/config.json", "path to config file")
	flag.Parse()
	utils.InitDB(*configPath)

	// CORS 配置
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:9528"}, // 允许的源
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Token"},
	}))

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务器
	err := router.Run(":5484")
	if err != nil {
		return
	}
}
