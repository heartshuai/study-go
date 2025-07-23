package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/OldPackageTest/stream_gprc_test/proto"
	"sync"
	"time"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	////服务端流模式
	//res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "hello"})
	//for {
	//	data, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(data.Data)
	//}
	//客户端流模式
	//putS, _ := c.PutStream(context.Background())
	//i := 0
	//for {
	//	i++
	//	_ = putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("hello %d", i)})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}
	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()
	//1.集中学习protobuf,grpc
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: "我是服务器",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
