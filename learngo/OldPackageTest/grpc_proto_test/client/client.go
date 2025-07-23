package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learngo/OldPackageTest/grpc_proto_test/proto_bak"
	"log"
	"time"
)

func main() {
	// 建立客户端连接（不加密）
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto_bak.NewGreeterClient(conn)
	// 调用 RPC 方法
	r, err := c.SayHello(context.Background(), &proto_bak.HelloRequest{Name: "yxs", Url: "https://baidu.com", G: proto_bak.Gender_MALE, Mp: map[string]string{
		"name": "yxs",
		"url":  "https://baidu.com",
	},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}

	println(r.Message)

}
