package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	protobuftest "learngo/OldPackageTest/helloworld/proto/helloworld/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := protobuftest.HelloRequest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"math", "english"},
	}
	jsonStruct := Hello{Name: "bobby", Age: 18, Courses: []string{"math", "english"}}
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Println(len(jsonRsp))
	rsp, _ := proto.Marshal(&req) //具体的编码是如何做到的，protobuf的原理
	newReq := protobuftest.HelloRequest{}
	proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
