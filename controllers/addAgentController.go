package controllers

import (
	"HG-Dashboard/models"
	"HG-Dashboard/utils"
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
