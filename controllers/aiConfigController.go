package controllers

import (
	"HG-Dashboard/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

func GetAIConfig(path string) (*models.Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg models.Config
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ProcessAIConfig(configPath string) ([]string, error) {
	jsonConfig, err := GetAIConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	yamlConfigPath := jsonConfig.AI.Config

	// 读取 YAML 配置文件
	yamlBytes, err := os.ReadFile(yamlConfigPath)
	if err != nil {
		return nil, err
	}

	var yamlConfig models.YAMLConfig
	err = yaml.Unmarshal(yamlBytes, &yamlConfig)
	if err != nil {
		return nil, err
	}

	return yamlConfig.USERTOKENS, nil
}

func GetAIConfigHandler(c *gin.Context) {
	userTokens, err := ProcessAIConfig("./config/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usertokens": userTokens})
}
