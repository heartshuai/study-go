package server

import (
	"context"
	"learngo/OldPackageTest/grpc_test/proto/helloworld/proto"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedGreeterServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	//g := grpc.NewServer()
	//proto.RegisterGreeterServer(g, &Server{})
}
