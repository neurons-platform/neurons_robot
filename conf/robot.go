package conf

type RobotConf struct {
	Name             string
	Aid              string
	Token            string
	Department       string
	SendURL          string
	AgentRedisAddr   string
	AgentRedisPass   string
	AgentMqApp       string
	AgentMqPass      string
	AgentReportTopic string
	AlarmGroup       string
	CmdProfile       string
	LuaProfile       string
}

var Robot = &RobotConf{
	Name:             "robot",
	Aid:              "abc",
	Token:            "abc",
	Department:       "im",
	SendURL:          "http:://v5b7.com/send",
	AgentRedisAddr:   "127.0.0.1:5360",
	AgentRedisPass:   "pass",
	AgentMqApp:       "robot",
	AgentMqPass:      "robot",
	AgentReportTopic: "robot",
	AlarmGroup:       "1234",
	CmdProfile:       "script",
	LuaProfile:       "script",
}
