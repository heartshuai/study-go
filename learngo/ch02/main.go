package main

import "fmt"

func main() {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//
	//var e int //动态类型 用的时候比较麻烦
	//
	//a= int8(b)
	//
	//var f1 float32 //3.4e38
	//var f2 float64 //1.8e308
	//
	//f1 = 3.14
	//f2 = 3.14
	//
	//f1 = 3

	//var c byte  // 主要适用于存放字符 char类型
	var c2 rune //也是字符

	c2 = '测'
	//c = 'a' + 1
	//c1 := 97
	fmt.Printf("c=%c\n", c2)

	var name string
	name = "I am boboy"
	fmt.Println(name)

}
