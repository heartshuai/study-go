package main

import "fmt"

func main() {
	//常量 定义的时候就指定的值 不能修改 常量尽量全部大写
	const PI float32 = 3.1415926 //显式定义

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		x int = 16
		y
		s = "abc"
		z
	)
	fmt.Println(x, y, s, z)
	/**
	常量类型只可以定义bool、数值（整数、浮点数和整数） 和字符串
	不曾使用的常量，没有强制使用的要求
	显示指定类型的时候，必须确保常量左右数据一致
	*/

}
