package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learngo/OldPackageTest/grpc_validate_test/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	return &proto.Person{Id: 32}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	//p := new(proto.Person)
	//p.Id = 1000
	//err := p.Validate()
	//if err != nil {
	//	panic(err)
	//}
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r any, ok error) {
		//proto. Person 可以满足person的验证 所有的接口 还有其他的接口怎么办
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
	defer g.Stop()

}
