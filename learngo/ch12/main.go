package main

import (
	"fmt"
	"time"
)

/**
python java  php 多线程  多进程编程 多线程和多进程存在的问题是 主要是耗费内存
都会被操作系统调度 内存占用很大 线程切换 耗费cpu资源 web2.0 用户级线程 绿程 轻量级线程 协程
内存占用小 切换快 go语言的协程
*/

func asyncPrint() {
	for {
		time.Sleep(time.Second)
		fmt.Println("bobby")
	}
}

// 主协程
func main() {
	//主死随从
	//匿名函数 启动goroutine

	//1.闭包 for循环的时候 每个变量会重用
	//每次for循环的时候 i变量会复用 当进行到第二轮的for循环的时候 这个i就变了
	for i := 0; i <= 100; i++ {

		//tmp := i
		go func(i int) {
			fmt.Println(i)

		}(i)

	}
	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
