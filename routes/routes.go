package routes

import (
	"HG-Dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/login", controllers.Login)
	router.POST("/api/collect", controllers.CollectSystemInfo)
	router.GET("/api/status", controllers.GetSystemStatus)
	router.POST("/reset-password", controllers.ResetPassword)
	router.POST("/api/add-agent", controllers.AddAgent)
}
