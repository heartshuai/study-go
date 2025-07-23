package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"learngo/OldPackageTest/grpc_error_test/proto/helloworld/proto"
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 调用 RPC 方法
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "yxs"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}

	//println(r.Message)

}
