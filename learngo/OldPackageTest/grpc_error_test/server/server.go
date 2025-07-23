package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/OldPackageTest/grpc_error_test/proto/helloworld/proto"
	"net"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	//return nil, status.Errorf(codes.NotFound, "记录未找到 %s", request.Name)
	time.Sleep(time.Second * 5)
	return &proto.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	/**
	server端最好是帮我们生成接口，我们只需要去每个接口中实现对应的业务逻辑就行了
	client端需要帮我们生成对应的方法，同时将这个方法都绑定到一个结构体上，生成的时候我们可能需要传参数，ip:port
	java,python __init__() 函数

	*/
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
	defer g.Stop()
}
