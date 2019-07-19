package client

import (
	U "github.com/neurons-platform/gotools/utils"
)

type HTTPSender struct {
	Addr string `json:"addr"`
}

func (s *HTTPSender) Send(msg map[string]string) bool {
	r := U.HttpPostJson(s.Addr, msg["msg"])
	if len(r) > 0 {
		return true
	} else {
		return false
	}
}
