package main

import (
	"fmt"
	"time"
)

func main() {

	var msg chan int
	msg = make(chan int, 2) //有缓冲

	go func(msg chan int) { //go有一种happen-before的机制 可以保障
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("done")
	}(msg)

	msg <- 1
	msg <- 2
	msg <- 3
	close(msg) //其他的编程语言有=很大的区别 只有go能关闭 channel
	d := <-msg //已经关闭的channel中可以继续取值，但是不能再放值了
	fmt.Println(d)

	//waitgroup 如果少了done调用 容易出现deadlock 无缓冲的channel也容易出现
	time.Sleep(time.Second * 10)

}
