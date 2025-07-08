package main

import (
	"fmt"
	"time"
)

// 函数参数传递的时候，值传递，引用传递，go语言中全部都是值传递 指针
func add(item ...int) (sum int, err error) {
	for _, v := range item {
		sum += v
	}
	return sum, nil
}

func add2(a int, b int) {
	fmt.Printf("sum is :%d\r\n", a+b)
	//return 1
}

// 省略号

func runForever() {
	for {
		time.Sleep(time.Second)
		fmt.Println("run forever")
	}
}

func autoIncrement() func() int {
	local := 0 //一个函数中访问另一个函数的局部变量 不行的 闭包
	//local += 1
	return func() int {
		local += 1
		return local
	}
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func cal(myfunc func(items ...int) int) int {
	return myfunc()

}

func main() {
	//go 函数支持普通函数 匿名函数 闭包
	/**
	go中函数是“一等公民”
		1. 函数本身可以当做变量
		2.匿名函数 闭包
		3. 函数可以满足接口
	*/
	//funcVar := add
	//
	//a := 1
	//b := 2
	//sum, _ := funcVar(a, b, 3, 3, 4, 5, 6)
	//fmt.Println(sum)

	//runForever()
	sum := cal(func(items ...int) int {
		sum := 0
		for _, v := range items {
			sum += v
		}
		return sum
	})
	fmt.Println(sum)
	callback(1, add2)

	//闭包 有个需求，我希望有个函数每次调用一次返回的结果值都是增加后的值
	func() {
		fmt.Println("hello world")
	}()
	nextFunc := autoIncrement()
	for i := 0; i < 10; i++ {
		fmt.Println(nextFunc())
	}
}
func A() {
	a := 1
	fmt.Println(a)
}
