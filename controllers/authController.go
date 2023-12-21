package controllers

import (
	"HG-Dashboard/models"
	"HG-Dashboard/utils"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var loginInfo models.LoginInfo
	if err := c.BindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		storedPassword string
		userRole       int
	)
	// 查询数据库以获取存储的密码和用户角色
	err := utils.DB.QueryRow("SELECT password, role FROM myai_user WHERE username = ?", loginInfo.Username).Scan(&storedPassword, &userRole)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 用户凭证验证成功后，创建 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginInfo.Username,
		"role":     userRole, // 将用户角色添加到 Token
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("234wer234fwgr")) // 替换为您的密钥
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 将 Token 和用户角色发送回客户端
	c.JSON(http.StatusOK, gin.H{"code": 20000, "token": tokenString, "role": userRole})
}

var jwtKey = []byte("234wer234fwgr")

func Logout(c *gin.Context) {
	tokenString := c.GetHeader("X-Token")

	// 解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 发送成功响应
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}

func GetUserInfo(c *gin.Context) {
	tokenString := c.GetHeader("X-Token")

	// 解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 发送成功响应
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": gin.H{
		"username": claims["username"],
		"role":     claims["role"],
	}})
}
