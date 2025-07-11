package main

import "fmt"

// 我现在想要放多个person的信息到一个集合中
var persons [][]interface{}

// 个人信息 name age address height
// 类型集合
type Person struct {
	name    string
	age     int
	address struct{}
	height  float64
}

func main() {
	//很麻烦
	//persons := append(persons, []interface{}{"body1", "1", "body3", 1.80}) //断言

	//如何初始化结构体
	//person := Person{
	//	name:    "body1",
	//	age:     1,
	//	address: "body3",
	//	height:  1.80,
	//}
	//var persons []Person
	//persons = append(persons, person)
	//fmt.Println(persons)
	//
	//person2 := []Person{
	//	{
	//		name:    "body1",
	//		age:     1,
	//		address: "body3",
	//		height:  1.80,
	//	},
	//}
	//fmt.Println(person2)
	var p Person
	p.age = 20
	fmt.Println(p)

	//匿名结构体
	address := struct {
		province string
		city     string
	}{
		province: "广东",
		city:     "深圳",
	}

	fmt.Println(address)

}
