package controllers

import (
	"HG-Dashboard/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	mu               sync.RWMutex
	latestSystemInfo models.SystemInfo
)

func CollectSystemInfo(c *gin.Context) {
	var info models.SystemInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	latestSystemInfo = info
	mu.Unlock()

	c.Status(http.StatusOK)
}

func GetSystemStatus(c *gin.Context) {
	mu.RLock()
	defer mu.RUnlock()
	c.JSON(http.StatusOK, latestSystemInfo)
}
