package main

import (
	"HelloService/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println("receive: ", request.Name)
	fmt.Println("send: ", "hello "+request.Url)
	//url和name搞反了 protobuf的编码
	return &proto.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
	defer g.Stop()
}
