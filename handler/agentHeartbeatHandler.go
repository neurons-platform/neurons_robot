package handler

import (
	// "fmt"
	// "net"
	CI "neurons_robot/client_interface"
	P "neurons_robot/protocol"
	// U "github.com/neurons-platform/gotools/utils"
)


type AgentHeartbeatHandler struct {
}


func (hd *AgentHeartbeatHandler) SetSystem() Action {
	return ActionFunc(func(h HandlerStruct) HandlerStruct {
		msg := P.Fac.CreateProtocol("agentHeartbeat").(*P.AgentHeartbeat)
		msg.LoadFromJsonString(h.ProtocolStr)
		m := map[string]string{"op": "sadd","key":"all_system","value":msg.System}
		h.Sender.Send(m)
		return h
	})
}


func (hd *AgentHeartbeatHandler) SetIP() Action {
	return ActionFunc(func(h HandlerStruct) HandlerStruct {
		msg := P.Fac.CreateProtocol("agentHeartbeat").(*P.AgentHeartbeat)
		msg.LoadFromJsonString(h.ProtocolStr)
		m := map[string]string{"op": "sadd","key":"system_name_"+msg.System,"value":msg.IP}
		h.Sender.Send(m)
		return h
	})
}


func (this *AgentHeartbeatHandler) Handle(sender CI.Sender, mp map[string]interface{}, protocolStr string) {
	DoActions(
		HandlerStruct{
		Sender: sender,
		Mp: mp,
		ProtocolStr: protocolStr,
		Quit: false},
		this.SetSystem(),
		this.SetIP(),
	)
}

