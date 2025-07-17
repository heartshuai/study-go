package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learngo/OldPackageTest/stream_gprc_test/proto"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}

	}
	return nil
}
func (s *server) PutStream(g grpc.ClientStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	for {
		if req, err := g.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(req.Data)
		}
	}
	return nil
}

func (s *server) AllStream(g grpc.BidiStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := g.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = g.Send(&proto.StreamResData{
				Data: fmt.Sprintf("%v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

	return nil
}
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	_ = s.Serve(lis)
	defer s.Stop()

}
