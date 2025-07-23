package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learngo/OldPackageTest/grpc_validate_test/proto"
	"log"
	"time"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "1010101",
		"appkey": "I am Key",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "1010101",
	//		"appkey": "I am Key",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	log.Println("耗时：", time.Since(start))
	//	return err
	//}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))

	// 建立客户端连接（不加密）
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("无法连接到 gRPC 服务: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{"timestamp": time.Now().Format(time.RFC3339)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 调用 RPC 方法
	r, err := c.SayHello(ctx, &proto.Person{Id: 1000, Email: "yxs@sina.com", Mobile: "18401692345"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r)

}
