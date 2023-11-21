package models

type AddAgentInfo struct {
	AgentId   int    `json:"agentId"`
	AgentName string `json:"agentName"`
	AgentIp   string `json:"agentIp"`
	AgentKey  string `json:"agentKey"`
}
