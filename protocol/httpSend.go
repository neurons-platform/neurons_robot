package protocol

import (
	CF "neurons_robot/conf"
	U "github.com/neurons-platform/gotools/utils"
)

// 机器人发送消息 robot -> gw
type HTTPSend struct {
	// normal: 普通消息 , group: 群组消息, batch: 批量消息
	Version  string `json:"version"`
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Gid      string `json:"gid"`
	Content  string `json:"content"`
	URL      string `json:"url"`
	SendTime int64  `json:"sendTime"`
	Token    string `json:"token"`
}

func NewDefaultHTTPSend() *HTTPSend {
	msg := &HTTPSend{}
	msg.Version = "0.1"
	msg.From = CF.Robot.Name
	msg.SendTime = U.MilliTimeStamp()
	return msg
}

func (this *HTTPSend) CheckToken() bool {
	r := false
	if this.Token == this.GetToken() {
		r = true
	}
	return r

}

func (this *HTTPSend) GetToken() string {
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

func (this *HTTPSend) LoadFromJsonString(str string) {
	U.JsonStringToStruct(str, this)
}

func (this *HTTPSend) ToProtocolString() string {
	js := U.StructToJsonStringNotEscapHTML(this)
	return js
}
