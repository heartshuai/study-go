package main

import (
	"fmt"
	"time"
)

func main() {
	//for 循环
	//var i int
	//for i < 10 {
	//fmt.Println(i)
	//i++
	//}
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	//fmt.Println(sum)

	//打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	name := "hello go"
	for _, value := range name {
		//fmt.Println(key, value)
		fmt.Printf("%c \r\n", value)
	}

	round := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(round)
		round++
		if round > 10 {
			break
		}
	}
}
