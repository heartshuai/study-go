package main

import (
	"fmt"
	"sync"
)

/**
子goroutine 如何通知到主的goroutine自己结束了，主的goroutine如何知道子goroutine结束的
*/

func main() {
	var wg sync.WaitGroup
	//我要监控多少个goroutine结束
	//wg.Add(100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)

	}
	//等到所有的goroutine结束
	wg.Wait()
	//wait group 主要用于goroutine的执行等到 add方法要和done方法成对出现
}
