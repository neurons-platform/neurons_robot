package handler

import (
	// "net"
	CI "neurons_robot/client_interface"
	C "neurons_robot/cmd"
)

type HandlerStruct struct {
	// Conn        net.Conn
	Sender      CI.Sender
	Mp          map[string]interface{}
	From        string
	Content     string
	Cmd         C.Cmd
	ProtocolStr string
	Quit        bool
}

func DoActions(s HandlerStruct, actions ...Action) HandlerStruct {
	for i := range actions {
		a := actions[i]
		s = a.Do(s)
		if s.Quit {
			return s
		}
	}
	return s
}

type ActionFunc func(HandlerStruct) HandlerStruct

func (f ActionFunc) Do(h HandlerStruct) HandlerStruct {
	return f(h)
}

type Action interface {
	Do(HandlerStruct) HandlerStruct
}
