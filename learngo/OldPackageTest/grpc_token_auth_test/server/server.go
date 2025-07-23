package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"learngo/OldPackageTest/grpc_test/proto/helloworld/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}
	for k, v := range md {
		fmt.Println(k, v)
	}
	return &proto.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("interceptor")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println(ok, md)
			return resp, status.Error(codes.Unauthenticated, "get metadata error")
		}
		var (
			appid  string
			appkey string
		)
		if var1, ok := md["appid"]; ok {
			appid = var1[0]
		}
		if var2, ok := md["appkey"]; ok {
			appkey = var2[0]
		}
		if appid != "1010101" || appkey != "I am Key" {
			return resp, status.Error(codes.Unauthenticated, "appid or appkey error")
		}

		res, err := handler(ctx, req)
		fmt.Println("请求已完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
	defer g.Stop()
}
