package ai


import (
	U "github.com/neurons-platform/gotools/utils"
	"time"
)

var aiUrl = "http://10.188.65.100:80/"


type ChatRequest struct {
	// 问题
	Q string `json:"q"`
	// start
	S int `json:"s"`
	// limit
	L int `json:"l"`
}

type ChatResponse struct {
	A string `json:"a"`
	C int `json:"c"`
}

type TrainData struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type AIResult struct {
	Result string `json:"result"`
}

func chat(input string,total int) string {
	result := make(chan string, 200)
	timeout := time.After(2 * time.Second)

	for i := total; i >= 0; i-- {
		// fmt.Println(i)
		chatRequest := &ChatRequest{Q:input,S:i*10000,L:10000}
		postData := U.StructToJsonString(chatRequest)
		go func(data string) {
			result <- U.HttpPostJson(aiUrl+"/chat", data)
		}(postData)
	}

	maxConfidence := 0
	bestMatch := "不懂你说的是什么"
	for i := 0; i <= total; i++ {
		// fmt.Println(i)
		select {
		case res := <-result:
			r := &ChatResponse{}
			U.JsonStringToStruct(res, r)
			if r.C == 100 {
				bestMatch = r.A
				goto L
			}
			if r.C >= maxConfidence {
				maxConfidence = r.C
				bestMatch = r.A
			}
		case <-timeout:
			goto L
		}
	}
L:
	return bestMatch
}


func train(input string, response string) string {
	trainData := &TrainData{Q:input,A:response}
	data := U.StructToJsonString(trainData)
	return U.HttpPostJson(aiUrl+"/train", data)
}


func AITrain(input string, response string) string {
	return train(input, response)
}

func AIChat(str string) string {
	r := chat(str,0)
	return r
}
