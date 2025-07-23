package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"learngo/OldPackageTest/grpc_test/proto/helloworld/proto"
	"log"
	"time"
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
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{"timestamp": time.Now().Format(time.RFC3339)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 调用 RPC 方法
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "yxs"})
	if err != nil {
		panic(err)
	}

	println(r.Message)

}
