package HG_Dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type SystemInfo struct {
	CPUUtilization    float64   `json:"cpu_utilization"`
	MemoryUtilization float64   `json:"memory_utilization"`
	DiskUsage         float64   `json:"disk_usage"`
	BandwidthUsage    Bandwidth `json:"bandwidth_usage"`
}

type Bandwidth struct {
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}

var (
	// 使用互斥锁保护数据
	mu sync.RWMutex
	// 存储最新的系统信息
	latestSystemInfo SystemInfo
)

func main() {
	r := gin.Default()

	// 接收被监控端数据的 API
	r.POST("/api/collect", func(c *gin.Context) {
		var info SystemInfo
		if err := c.BindJSON(&info); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		latestSystemInfo = info
		mu.Unlock()

		c.Status(http.StatusOK)
	})

	// 提供数据查询的 API
	r.GET("/api/status", func(c *gin.Context) {
		mu.RLock()
		defer mu.RUnlock()
		c.JSON(http.StatusOK, latestSystemInfo)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
