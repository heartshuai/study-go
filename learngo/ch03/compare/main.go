package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串的比较
	a := "hello"
	b := "bello"
	fmt.Println(a != b)

	//字符串的大小比较

	fmt.Println(a > b)

	//字符串是否包含
	name := "imooc体系课-go"
	fmt.Println(strings.Contains(name, "go"))

	//出现次数
	fmt.Println(strings.Count(name, "o"))

	//分割
	fmt.Println(strings.Split(name, "-"))

	//字符串是否包含前缀，是否包含后缀
	fmt.Println(strings.HasPrefix(name, "imooc"))
	fmt.Println(strings.HasSuffix(name, "go"))

	fmt.Println(strings.Index(name, "go"))

	//查询子串出现的位置
	fmt.Println(strings.IndexRune(name, []rune("go")[0]))

	fmt.Println(strings.LastIndex(name, "go"))

	fmt.Println(strings.Replace(name, "o", "go语言", -1))

	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))

	//去掉左右特殊字符
	fmt.Println(strings.Trim(name, "go"))

}
