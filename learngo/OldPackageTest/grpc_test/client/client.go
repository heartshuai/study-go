package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learngo/OldPackageTest/grpc_test/proto/helloworld/proto"
	"log"
)

func main() {
	// 建立客户端连接（不加密）
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 调用 RPC 方法
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "yanxingshuai"})
	if err != nil {
		panic(err)
	}

	println(r.Message)

}
