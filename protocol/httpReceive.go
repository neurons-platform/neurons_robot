package protocol

import (
	CF "neurons_robot/conf"
	U "github.com/neurons-platform/gotools/utils"
)

// 机器人接收消息 robot <- gw
type HTTPReceive struct {
	// normal: 普通消息,group: 群组消息,notice: 通知消息
	Version string `json:"version"`
	Type    string `json:"type"`
	// 消息来源
	From string `json:"from"`
	To   string `json:"to"`
	Gid  string `json:"gid"`
	// 消息内容
	Content string `json:"content"`
	URL     string `json:"url"`
	// 发送时间戳
	SendTime int64  `json:"sendTime"`
	Token    string `json:"token"`
}

func NewDefaultHTTPReceive() *HTTPReceive {
	msg := &HTTPReceive{}
	return msg
}

func (this *HTTPReceive) MsgToHTTPReceive(msg string) *HTTPReceive {
	this.Version = "0.1"
	this.Type = "normal"
	this.To = CF.Robot.Name
	this.Content = msg
	this.From = "jingminglang"
	this.SendTime = U.MilliTimeStamp()
	this.Token = this.GetToken()
	return this
}

func (this *HTTPReceive) CheckToken() bool {
	r := false
	if this.Token == this.GetToken() {
		r = true
	}
	return r
}

func (this *HTTPReceive) GetToken() string {
	r := ""
	r = U.Md5(
		this.Type +
			this.From +
			this.To +
			this.Content +
			this.URL +
			U.Int64toStr(this.SendTime) +
			CF.Robot.Token)
	return r
}

func (this *HTTPReceive) LoadFromJsonString(str string) {
	U.JsonStringToStruct(str, this)
}

func (this *HTTPReceive) ToProtocolString() string {
	js := U.StructToJsonStringNotEscapHTML(this)
	return js
}
