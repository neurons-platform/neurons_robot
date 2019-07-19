package cmd

import (
	U "github.com/neurons-platform/gotools/utils"
)


type JsonMethod struct {
	MethodSuper
}

func (m *JsonMethod) GetMethodResult() string {
	url := U.ParserMapToString(m.M, m.A.Url)
	data := U.ParserMapToString(m.M, m.A.Data)
	r := U.HttpPostJson(url, data)
	return r
}



func NewJsonMethod(a *Action, m map[string]string) *JsonMethod {
	method := new(JsonMethod)
	method.A = a
	method.M = m
	method.MethodSuper.GetMethodResult = method.GetMethodResult
	return method
}
