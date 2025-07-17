package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/OldPackageTest/grpc_test/proto/helloworld/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
	defer g.Stop()
}
