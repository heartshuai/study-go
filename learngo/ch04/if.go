package main

import "fmt"

/**
if 布尔表达式 {
	逻辑
}

*/

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("可以投票")
	} else if age >= 16 {
		fmt.Println("可以参与问卷调查")
	} else {
		fmt.Println("不可以投票")
	}
}
