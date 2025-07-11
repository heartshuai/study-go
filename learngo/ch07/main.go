package main

import (
	"fmt"
	"strconv"
)

type MyInt int

func (mi MyInt) aaa() string {
	return strconv.Itoa(int(mi))
}
func main() {
	//type关键字
	/**
	定义结构体
	定义接口
	定义类型别名
	类型定义
	类型判断
	*/

	//别名实际上是为了更好地理解代码，而不是为了提高性能。
	//type myInt1 = int //类型别名
	//var a myInt1 = 12
	//var j int = 8
	//fmt.Println(a + j) //在编译的时候 类型别名会被直接替换为int
	//fmt.Printf("%T", a)
	//type MyInt int
	//var a MyInt = 12
	//fmt.Println(a.aaa())
	//var j int = 8 //寄希望你是int类型 又希望可以增加方法
	//fmt.Println(int(a) + j)
	//fmt.Printf("%T", a)

	var a interface{} = "abc"
	//switch a.(type) {
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//}
	//m := a.(string)
	fmt.Println(a)
}
