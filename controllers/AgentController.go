package controllers

import (
	"HG-Dashboard/models"
	"HG-Dashboard/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func generateKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result.WriteByte(charset[randomIndex])
	}

	return result.String()
}

func AddAgent(c *gin.Context) {
	var addAgentInfo models.AddAgentInfo
	if err := c.BindJSON(&addAgentInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	agentkey := generateKey(32)
	addAgentInfo.AgentKey = agentkey

	if err := addAgentDatabaese(addAgentInfo.AgentName, addAgentInfo.AgentIp, addAgentInfo.AgentKey); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Add agent successful"})
}

func addAgentDatabaese(agentname, agentip, agentkey string) error {
	utils.DB.Exec("INSERT INTO hg_agent_info (agent_name,agent_ip,agent_key) VALUES (?,?,?)", agentname, agentip, agentkey)
	return nil
}

// GetAllAgents 获取所有代理信息
func GetAllAgents(c *gin.Context) {
	// 执行查询
	rows, err := utils.DB.Query("SELECT agent_id, agent_name, agent_ip, agent_key FROM hg_agent_info")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// 读取结果
	var agents []models.AddAgentInfo
	for rows.Next() {
		var agent models.AddAgentInfo
		if err := rows.Scan(&agent.AgentId, &agent.AgentName, &agent.AgentIp, &agent.AgentKey); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database result"})
			return
		}
		agents = append(agents, agent)
	}

	// 检查查询过程中是否有错误
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during rows iteration"})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"agents": agents})
}
