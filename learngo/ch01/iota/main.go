package main

import "fmt"

func a() (int, bool) {
	return 0, false
}

var name = "bobby"

func main() {
	//匿名变量 变量的作用域

	var _ int
	_, ok := a()
	if ok {
		//打印
	}
	var mainName = "main"
	fmt.Println(mainName)
	//iota 特殊常量 可以认为是一个可以被编译器修改的常量
	const (
		ERR1 = iota
		ERR2
		ERR25 = "ha"
		ERR26
		ERR3

		ERR4 = iota
	)

	const (
		ERRNEW1 = iota
	)
	const a = iota
	const b = iota

	fmt.Println(ERR1, ERR2, ERR26, ERR3, ERR4)

	/**
	如果中断了iota那么必须显示的恢复 后续会自动递增
	自增类型默认是int类型 iota能简化const类型的定义
	*/

	//{
	//	localName := "local"
	//	fmt.Println(localName)
	//}
	var mname string
	if name == "bobby" {
		mname := "imooc"
		fmt.Println(mname)
	}
	fmt.Println(mname)

}
