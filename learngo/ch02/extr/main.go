package main

import "fmt"

func main() {
	// 运算符
	var a, b = 1, 2
	fmt.Println(a + b)
	var astr, bstr = "hello", "bobby"
	fmt.Println(astr + bstr)
	fmt.Println(3 % 2)
	a++
	fmt.Println(a)

	var abool, bbool = true, false
	if abool && bbool {
		fmt.Println("true")
	}
	if abool || bbool {
		fmt.Println("true")
	}

	if !abool {
		fmt.Println("false")
	}
	//位运算符 性能要求高的时候 一般会考虑与运算

	var A, B = 60, 13
	//var c *int
	//c :=&A

	fmt.Println(A & B)

}
