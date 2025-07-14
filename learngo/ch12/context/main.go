package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 共享变量

func cpuInfo(ctx context.Context) {
	//这里能拿到一个请求的id
	fmt.Println(ctx.Value("name"))
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu is ")
		}

	}

}
func main() {
	/**
	渐进式的方式
	有一个goroutine监控cpu的信息
	*/
	//var stop = make(chan struct{})
	wg.Add(1)
	//context 提供了三中函数 withCancel withValue withTimeout
	//ctx, cancel := context.WithCancel(context.Background())
	//如果你的goroutine 函数中 如果希望被控制，超时，传值，但是我不希望影响我原来的接口信息的时候，函数参数中第一个参数就是context
	//1.ctx1, cancel1 := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1)
	//2.timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//3.withdeadline 主动设置截止时间

	//4.withValue
	ctx = context.WithValue(ctx, "name", "zhangsan")
	go cpuInfo(ctx)
	//time.Sleep(6 * time.Second)
	//cancel1()
	wg.Wait()

	fmt.Println("监控完成")

}
