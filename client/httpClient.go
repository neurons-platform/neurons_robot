package client

import (
	"fmt"
	U "github.com/neurons-platform/gotools/utils"
	CI "neurons_robot/client_interface"
	CF "neurons_robot/conf"
	H "neurons_robot/handler"
	P "neurons_robot/protocol"
	"time"
)

type HttpClient struct {
	Sender   CI.Sender
	Receiver CI.Receiver
}

func (h *HttpClient) Init() {
	h.Sender = &HTTPSender{Addr: CF.Robot.SendURL}

	r := &HttpReceiver{
		Addr: "",
		Mq:   []string{},
	}
	go r.Init()
	h.Receiver = r
}

func (h *HttpClient) Recv() {
	for {
		msg := h.Receiver.Recv()
		if len(msg) > 0 {
			fmt.Println("recv: ++++++++++++++++++++++++" + msg)
			rMsg := P.Fac.CreateProtocol("httpReceive").(*P.HTTPReceive)
			rMsg.LoadFromJsonString(msg)
			token := rMsg.GetToken()
			fmt.Print("token: -----------------------")
			fmt.Println(token)
			if rMsg.CheckToken() {
				mp := U.JsonStringToMap(msg)
				fmt.Print("recv: -----------------------")
				fmt.Println(mp)
				switch mp["type"] {
				case "normal":
					go H.Fac.CreateHandler("httpReceiveMsg").Handle(h.Sender, mp, msg)
				case "group":
					go H.Fac.CreateHandler("httpReceiveMsg").Handle(h.Sender, mp, msg)
				case "notice":
					go H.Fac.CreateHandler("httpReceiveNotice").Handle(h.Sender, mp, msg)
				default:
					fmt.Println(msg)
				}
			}

		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
