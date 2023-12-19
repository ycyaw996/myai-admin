package routes

import (
	"HG-Dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/login", controllers.Login)
	router.POST("/api/logout", controllers.Logout)
}
