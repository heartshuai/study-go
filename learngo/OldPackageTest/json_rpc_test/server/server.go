package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func main() {
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")

	//2.注册处理handle
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	/**
	一连串的代码大部分都是net的包好像和rpc没有关系
	不行 rpc调用中有几个问题需要解决 1 call id 2.序列化和反序列化 编码和解码
	python下的开发而言 这个显得不好用
	可以跨语言调用 1.go语言的rpc的序列化和反序列化的协议是什么 Gob 2.能否替换成常见的序列化
	*/

}
