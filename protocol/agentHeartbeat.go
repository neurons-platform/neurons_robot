package protocol

import (
	U "github.com/neurons-platform/gotools/utils"
)

// agent 心跳协议
type AgentHeartbeat struct {
	Type    string   `json:"protocol_type"`
	Vertion string   `json:"version"`
	IP       string   `json:"ip"`
	System   string   `json:"system"`
	Env      string   `json:"env"`
	Cluster    string `json:"cluster"`
}

func NewDefaultAgentHeartbeat() *AgentHeartbeat {
	agentHeartbeat := &AgentHeartbeat{}
	return agentHeartbeat
}


func (this *AgentHeartbeat) JsonStringToProtocolStruct(str string) {
	U.JsonStringToStruct(str, this)
}

func (this *AgentHeartbeat) ToProtocolString() string {
	js := U.StructToJsonStringNotEscapHTML(this)
	return js
}

func (this *AgentHeartbeat) LoadFromJsonString(str string) {
	U.JsonStringToStruct(str, this)
}
