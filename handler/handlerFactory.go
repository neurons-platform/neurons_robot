package handler

import (
	"fmt"
	CI "neurons_robot/client_interface"
)

var Fac = NewHandlerFactory()

type Handler interface {
	Handle(CI.Sender, map[string]interface{}, string)
}

type HandlerFactory struct {
}

func NewHandlerFactory() *HandlerFactory {
	return &HandlerFactory{}
}

func (this *HandlerFactory) CreateHandler(handlerName string) Handler {
	switch handlerName {
	default:
		fmt.Println("无效协议")
		return nil
	}
}
