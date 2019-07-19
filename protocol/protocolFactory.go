package protocol

import (
	"fmt"
)

var Fac = NewProtocolFactory()

type Protocol interface {
	ToProtocolString() string
}

type ProtocolFactory struct {
}

func NewProtocolFactory() *ProtocolFactory {
	return &ProtocolFactory{}
}

func (this *ProtocolFactory) CreateProtocol(protocolName string) Protocol {
	switch protocolName {
	case "httpReceive":
		return NewDefaultHTTPReceive()
	case "httpSend":
		return NewDefaultHTTPSend()
	default:
		fmt.Println("无效协议")
		return nil
	}
}
