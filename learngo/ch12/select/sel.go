package main

import (
	"fmt"
	"time"
)

//var done = make(chan struct{}) //channel 要初始化 是多线程安全
//var lock sync.Mutex

// 很多时候我并不会多个goroutine去操作同一个channel
func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}

}
func g2(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
}
func main() {
	/**
	select 类似于 switch case语句，但是select的功能和我们操作linux里面提供的io的select poll epoll
	select 主要作用于多个channel
	现在有个需求 我们现在有两个goroutine 但是呢我在主的goroutine中，当某一个执行完成以后，这个时候我可以立马知道
	*/
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go g1(g1Channel)
	go g2(g2Channel)
	//我要监控多个channel 任何一个channel 返回都知道
	/**
	1.某一个分支就绪了就执行该分支
	2.如果两个都就绪了，随机执行一个
	*/
	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("default")
			return
		}
	}

	//for {
	//	if done {
	//		fmt.Println("done")
	//		time.Sleep(10 * time.Microsecond)
	//		return
	//	}
	//}
}
