package cmd

import (
	"errors"
	"regexp"
	U "github.com/neurons-platform/gotools/utils"
)


type MethodSuper struct {
	A               *Action
	M               map[string]string
	GetMethodResult func() string
}

func (m MethodSuper) DoMethod() (string, error) {

	var r string
	r = m.GetMethodResult()
	match, _ := regexp.Match(m.A.Success, []byte(r))
	if !match {
		r = "执行: " + U.Int2Str(m.A.Step) + " 失败"
		return r, errors.New("执行失败")
	}
	if len(m.A.JsonResultFormat) > 0 {
		rf, ok := U.ParserJsonStrToTemplateStr(r, m.A.JsonResultFormat)
		if ok {
			r = rf
			if m.A.AutoExe {
				cmd, err := GetCmd(r)
				if err == nil {
					r = cmd.Do()
				}
			}
		}
	}
	return r, nil

}







func printArray(arr []string) {
	for i,v := range arr {
		U.LogPrintln(i)
		U.LogPrintln(v)
	}
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}











