package cmd

import (
	// "bytes"
	"errors"
	"fmt"
	"regexp"
	// "strings"
	U "github.com/neurons-platform/gotools/utils"
	PS "github.com/neurons-platform/neurons_script/lib/parser"
	CF "neurons_robot/conf"
	// "text/template"
)

type Cmd struct {
	Name    string `json:"name"`
	CmdName string `json:"cmd_name"`
	// 如果中间有一个action执行失败是否继续执行下面的action
	FailContinue bool `json:"fail_continue"`
	// eval可以写一个类似shell的简单程序语言来进行变量的判断和替换
	Eval string `json:"eval,omitempty"`
	// 该命令是否被废弃
	Disable bool `json:"disable"`
	// 执行命令作为消息发给用户
	Msg string `json:"msg"`
	// 命令包含的变量和对应的值
	Params map[string]string
	// 有执行这个命令权限的erp
	Users []string `json:"users"`
	// 命令的执行动作
	Actions []Action `json:"actions"`
	// 可以作为参数传给函数
	Content string `json:"content"`
}

type Action struct {
	Step int    `json:"step"`
	Url  string `json:"url,omitempty"`
	// 可以作为FUNC的参数 会做变量的替换
	Arg     string `json:"arg,omitempty"`
	Method  string `json:"method"`
	Data    string `json:"data,omitempty"`
	Success string `json:"success,omitempty"`
	// 自动把action的返回结果作为命令执行
	AutoExe bool `json:"autoExe,omitempty"`
	// 如果返回结果是json格式的，可以通过这个配置提取json中的内容作为返回结果
	JsonResultFormat string `json:"jsonResultFormat,omitempty"`
	// 可以做为参数传入给本地代码
	Content string `json:"content,omitempty"`
	// 要执行的本地函数
	Function string `json:"function,omitempty"`
	// 执行命令 当method为CMD/SHELL时才有效
	Cmd string `json:"cmd"`
}

func (c *Cmd) Do() string {
	var rst string

	if len(c.Params) <= 0 {
		rst = "命令格式错误"
		return rst
	}

	if _, ok := c.Params["exit_cmd"]; ok {
		rst = "输入非法参数"
	} else {
		for _, a := range c.Actions {
			//TODO: 这个地方写的不好，容易引发bug
			// 记录用户的原始输入信息或指定的参数
			c.Params["_cmd_content"] = c.Content
			// c.Params["_cmd_content"] = "test"
			r, err := a.Act(c.Params)
			rst = rst + r + "\n"
			// rst = rst + "Step " + U.Int2Str(a.Step) + ": " + r + "\n"

			if err != nil {
				if c.FailContinue != true {
					break
				}
			}
		}
	}

	return rst
}

func (a *Action) Act(m map[string]string) (string, error) {
	var r string
	if a.Method == "GET" {
		return NewGetMethod(a, m).DoMethod()
	}
	if a.Method == "POST" {
		return NewPostMethod(a, m).DoMethod()
	}
	if a.Method == "JSON" {
		return NewJsonMethod(a, m).DoMethod()
	}
	if a.Method == "CMD" {
		return NewCmdMethod(a, m).DoMethod()
	}
	return r, nil
}

func GetCmdName(content string) string {
	r, _ := regexp.Compile(`=\s+(?P<Name>[^\s]*)`)
	m, _ := U.ParserStringToMap(content, r)
	return m["Name"]
}

func GetCmd(content string) (Cmd, error) {
	var cmd Cmd
	cmdName := GetCmdName(content)
	body := CF.GetCmdConf(cmdName)
	// U.LogPrintln(body)
	// body := U.HttpGet(ucc_addr)
	if body == "" {
		return cmd, errors.New("cmd 解析失败")
	}
	U.JsonStringToStruct(body, &cmd)
	// log.Println("get_cmds cmds:", cmds)
	// fmt.Println(cmd.Name)
	cmd.CmdName = cmdName
	// reg := U.FormatToReg(cmd.Name)
	fmt.Println("===========================cmd content================================")
	U.LogPrintln(content)
	fmt.Println("===========================cmd content================================")

	fmt.Println("===========================cmd Name================================")
	U.LogPrintln(cmdName)
	fmt.Println("===========================cmd Name================================")
	reg := U.FormatToReg2(cmd.Name)
	fmt.Println("===========================cmd reg================================")
	U.LogPrintln(reg)
	fmt.Println("===========================cmd reg================================")
	// fmt.Println(reg)
	m, t := U.ParserStringToMap(content, reg)
	fmt.Println("===========================cmd================================")
	U.LogPrintln(cmd)
	fmt.Println("===========================cmd================================")
	// fmt.Println(t)
	// fmt.Println(m)
	if t {
		PS.Eval(cmd.Eval, m)
		// fmt.Println(m)
		if len(m) > 0 {
			cmd.Params = m
			return cmd, nil
		}
	}
	if cmd.Name == "" {
		return cmd, errors.New("cmd 为空")
	}
	return cmd, nil
}
