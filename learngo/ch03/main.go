package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	name := "imooc 体系课学习"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	courseName := `go"体系课"`
	fmt.Println(string(courseName))

	username := "bobby"
	age := 18
	address := "北京"
	mobile := "1892302020"

	//fmt.Println("用户名:" + username , ",年龄：" + strconv.Itoa(age) + ",地址:" +address)
	fmt.Printf("用户名:%s 年龄:%d 地址:%s 电话:%s", username, age, address, mobile)
	userMsg := fmt.Sprintf("用户名:%s 年龄:%d 地址:%s 电话:%s", username, age, address, mobile)
	fmt.Println(userMsg)

	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString("年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("地址:")
	builder.WriteString(address)
	builder.WriteString("电话:")
	builder.WriteString(mobile)
	fmt.Println(builder.String())

}
