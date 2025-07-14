package main

import (
	"fmt"
	"sync"
	"time"
)

/**
锁本质上是将并行的代码串行化 使用lock肯定会影响性能 避免使用锁
即使是设计锁 也应该尽量的保证并行
我们有两组协程 其中一组负责写数据 另一个组负责读数据 web系统中绝大部分的数据都是读数据 读多写少
虽然有多个goroutine 但是仔细分析我们会发现，读协程之间应该并发 读和写之间应该串行
*/

func main() {
	//var num int
	var rwlock sync.RWMutex

	var wg sync.WaitGroup
	wg.Add(6)
	//写的goroutine
	go func() {
		time.Sleep(time.Second)
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁获取，和读锁获取
		defer rwlock.Unlock()
		fmt.Println("write:lock")
		time.Sleep(time.Second * 5)

	}()
	//读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				rwlock.RLock() //读锁会防止写锁获取，但是读锁可以和读锁获取
				time.Sleep(500 * time.Microsecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}

		}()
	}

	wg.Wait()
}
