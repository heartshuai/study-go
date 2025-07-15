package server_proxy

import (
	"learngo/OldPackageTest/new_helloworld/handler"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

// 如何做到解耦-我们关心的是函数 鸭子类型
func RegisterHelloService(srv HelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
