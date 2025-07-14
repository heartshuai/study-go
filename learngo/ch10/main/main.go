package main

import (
	"fmt"
	u "learngo/ch10/user"
	//. "learngo/ch10/user" //导入的别名 尽量少用
	_ "learngo/ch09/user"
)

func main() {
	c := u.Course{
		Name: "Go语言",
	}
	name, err := u.GetCourse(c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)
	}
}
