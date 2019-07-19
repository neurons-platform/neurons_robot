package cmd

import (
	U "github.com/neurons-platform/gotools/utils"
)

type GetMethod struct {
	MethodSuper
}

func (m *GetMethod) GetMethodResult() string {
	url := U.ParserMapToString(m.M, m.A.Url)
	r := U.HttpGet(url)
	return r
}


func NewGetMethod(a *Action, m map[string]string) *GetMethod {
	method := new(GetMethod)
	method.A = a
	method.M = m
	method.MethodSuper.GetMethodResult = method.GetMethodResult
	return method
}
