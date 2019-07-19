package cmd

import (
	U "github.com/neurons-platform/gotools/utils"
)


type PostMethod struct {
	MethodSuper
}


func (m *PostMethod) GetMethodResult() string {
	url := U.ParserMapToString(m.M, m.A.Url)
	data := U.ParserMapToString(m.M, m.A.Data)
	r := U.HttpPost(url, data)
	return r

}

func NewPostMethod(a *Action, m map[string]string) *PostMethod {
	method := new(PostMethod)
	method.A = a
	method.M = m
	method.MethodSuper.GetMethodResult = method.GetMethodResult
	return method
}
