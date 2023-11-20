package models

type AddAgentInfo struct {
	AgentName string `json:"agentName"`
	AgentIp   string `json:"agentIp"`
	AgentKey  string `json:"agentKey"`
}
