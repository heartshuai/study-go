package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kirinlabs/HttpRequest"
)

// rpc远程过程调用 如何做到像本地调用
func Add(a, b int) int {
	//http协议 传输协议
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8080/add?a=%d&b=%d", a, b))
	body, _ := res.Body()
	//fmt.Println(string(body))
	respData := ResponseData{}
	_ = json.Unmarshal(body, &respData)
	return respData.Data
}

type ResponseData struct {
	Data int `json:"data"`
}

func main() {
	fmt.Println(Add(1, 2))
}
