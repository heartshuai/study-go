package main

func deferReturn() (Ret int) {
	defer func() {
		Ret++
	}()
	return 10
}

func main() {
	// 连接数据库 打开文件 开始锁 无论如何 最后都要记得去关闭数据库 关闭文件 解锁
	//var mu sync.Mutex
	//mu.Lock()
	//defer mu.Unlock() //defer 后面的代码是放在函数 return之前执行
	////打开文件 释放文件

	/**
	defer 先进后出
	*/
	//defer fmt.Println("1")
	//	defer fmt.Println("2")
	//
	//
	//
	//	fmt.Println("main")
	//	return
	//ret := deferReturn()
	//fmt.Println(ret)

	//error panic recover
	//go语言错误处理的理念，一个函数可能出错 try catch 包住这个函数，如果出错，就返回错误
	//开发函数的人需要有一个返回值告诉调用者是否成功 go设计者要求我们必须处理这个err 代码中大量的会出现 if err !=nil return err
	//go设计者认为必须要处理这个error防御性编程
}
