package routes

import (
	"HG-Dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/login", controllers.Login)
	router.POST("/api/logout", controllers.Logout)
	router.POST("/api/getInfo", controllers.GetUserInfo)
	router.POST("/api/getAIConfig", controllers.GetAIConfigHandler)
}
