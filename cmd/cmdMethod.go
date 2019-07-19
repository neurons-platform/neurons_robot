package cmd

import (
	U "github.com/neurons-platform/gotools/utils"
)

type CmdMethod struct {
	MethodSuper
}

func (m *CmdMethod) GetMethodResult() string {
	r := ""
	cmdStr := U.ParserMapToString(m.M, m.A.Cmd)
	cmd,err := GetCmd(cmdStr)
	if err == nil {
		r = cmd.Do()
	}
	return r
}

func NewCmdMethod(a *Action, m map[string]string) *CmdMethod {
	method := new(CmdMethod)
	method.A = a
	method.M = m
	method.MethodSuper.GetMethodResult = method.GetMethodResult
	return method
}
