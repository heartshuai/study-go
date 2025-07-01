package main

import "fmt"

//var name = "body"
//var age = 18
//var ok bool

var (
	name = "body"
	age  = 18
	ok   bool
)

func main() {
	//go 是静态语言 静态语言和动态语言相比变量差异很大
	// 1. 变量必须先定义后使用 2 变量必须有类型 3 类型定下来后不能改变
	//定义变量的方式

	//var name int
	//name = 1
	//var age = 1

	//age := 1
	var age2 int
	//go语言中变量定义了不使用是不行的

	fmt.Println(age)

	//2 多变量定义

	var user1, user2, user3 = "body1", 1, "body3"
	fmt.Println(user1, user2, user3)
	fmt.Println(age2)

	/*
		注意
			变量必须先定义才能使用
			go语言是静态经验，要求变量的类型和赋值的类型一致
			简洁变量不能用于全局
			变量是有零值的
	*/

}
