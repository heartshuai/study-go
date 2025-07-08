package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	//panic("this is an error") //panic会导致程序的退出，平时开发中不要随便用，一般在我们一个服务启动的过程中
	//比如我的服务需要启动，必须有些依赖启动成功后才能启动，日志文件存在，数据库连接成功这个时候服务才能启动的时候
	//如果我们的服务启动检查中发现了这些任何一个不满足 就调用panic
	//但是你的服务一旦启动了，这个时候你的某行代码不小心panic，那么不好意思你的程序挂了，这个是重大事故
	//但是架不住有些地方我们的代码写的不小心会导致被动触发panic
	//recover 这个函数能捕获到panic

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	//panic("this is an error")
	var names map[string]string
	names["go"] = "go工程师"
	return 0, errors.New("this is an error")
	//defer 需要放在panic之前定义，另外recover只有在defer调用的函数中才会生效
	//recover 处理异常后，逻辑并不会恢复到panic的那个点去
	//多个defer 会形成栈 后定义的defer会先执行
}

func main() {
	//panic 这个函数会导致你的程序退出 不推荐随便使用panic
	//panic("this is a panic")

	_, err := A()
	if err != nil {
		fmt.Println(err)
	}
}
