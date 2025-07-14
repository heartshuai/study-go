package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("num=%d\r\n", v)
	}
}

func main() {
	/**
	默认情况下 channel是双向的
	但是 我们经常一个channel 作为参数进行传递 希望对方是单向使用
	*/
	//var ch1 chan int
	//var ch2 chan<- float64 //单向channel 只能写入float64的数据
	//var ch3 <-chan int     //单向 只能读取int

	//c := make(chan int, 3)
	//var send chan<- int = c //send-only
	//var read <-chan int = c //recv-only
	//send <- 1
	//<-read

	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(time.Second * 10)

}
