package controllers

import (
	"HG-Dashboard/models"
	"HG-Dashboard/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPassword(c *gin.Context) {
	var req models.PasswordResetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// 替换为你的数据库更新逻辑
	// 假设有一个函数 updatePasswordInDatabase(username, hashedPassword) 来更新密码
	if err := updatePasswordInDatabase(req.Username, hashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}

func updatePasswordInDatabase(username, hashedPassword string) error {
	// 更新数据库中的密码
	// 注意：这里需要使用您的数据库逻辑来实际更新密码
	utils.DB.Exec("UPDATE hg_user SET passwd = ?", hashedPassword)
	utils.DB.Exec("UPDATE hg_user SET name = ?", username)
	return nil
}
