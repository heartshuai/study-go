package main

import "fmt"

func main() {
	/**
	goto 语句可以让我们跳转到指定位置执行 所以说很少用 goto语言使用场景最多的是 错误的处理
	*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto over
		}
		fmt.Println(i)
	}

	//end:
	//	fmt.Println("end")
over:
	fmt.Println("over")

}
