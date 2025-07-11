package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/**

	不同类型的数值零值不一样
	bool false
	numbers 0
	string ""
	pointer nil
	slice nill
	map nil
	channel interface function nil
	struct 默认值不是nil 默认值是具体字段的默认值
	*/
	//
	//p1 := Person{
	//	name: "张三",
	//	age:  18,
	//}
	//p2 := Person{
	//	"张四",
	//	19,
	//}
	//fmt.Println(p1 == p2)
	//
	////slice 的默认值
	////var ps []Person
	//var ps2 = make([]Person, 0)
	//if ps2 != nil {
	//	fmt.Println("slice is not nil")
	//}
	var m map[string]string
	var m2 = make(map[string]string, 0)

	for key, value := range m {
		fmt.Println(key, value)
	}
	//m["name"] = "bobby"
	m2["name"] = "bobby"
	for key, value := range m2 {
		fmt.Println(key, value)
	}

	//if m2 != nil {
	//	fmt.Println("map is not nil1")
	//}
	//if m != nil {
	//	fmt.Println("map is not nil2")
	//}
}
