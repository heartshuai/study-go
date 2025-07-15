package main

import (
	"fmt"
	"learngo/OldPackageTest/new_helloworld/client_proxy"
)

func main() {
	/**
	1.建立连接
	2.只想写业务逻辑 不想关注每个函数的名称
	*/
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	//var reply *string = new(string)
	var reply string
	//客户端部分
	err := client.Hello("bobby123", &reply)
	//client.Hello() 不用自己去封装 hello方法
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
	/**
	1.这些概念在grpc中都有对应
	2.发自灵魂的拷问:server_proxy和client_proxy能否自动生成
	3.都能满足 这个就是protobuf+grpc的一个实现
	*/
}
