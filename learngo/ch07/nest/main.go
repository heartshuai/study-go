package main

import "fmt"

type Person struct {
	name string
	age  int
	//address struct {
	//	city   string
	//	street string
	//}
}

func (p *Person) print() {
	//有可能该函数中想要改变p的值 就是python对象很大 数据较大的时候 考虑使用指针方式
	p.age = 180
	fmt.Printf("name:%s , age:%d", p.name, p.age)
}

func (p Person) print2() {
	p.age = 180
	fmt.Printf("name:%s , age:%d", p.name, p.age)
}

type Student struct {
	//第一种嵌套方式
	//P     Person
	//第二种嵌套方式
	Person
	score float32
	name  string
}

func main() {
	//s := Student{
	//	P: Person{
	//		Name: "张三",
	//		Age:  18,
	//	},
	//	Score: 90.0,
	//}
	//fmt.Println(s.P.Age)
	//	s := Student{}
	//	s.P.Name = "张三"
	//	fmt.Println(s.P.Name)
	//s := Student{
	//	Person{
	//		"张三",
	//		18,
	//	},
	//	90.0,
	//	"张三",
	//}
	//fmt.Println(s.name)

	//s := Student{
	//		Person{
	//			Name: "张三",
	//			Age:  18,
	//		},
	//		Score: 90.0,
	//	}
	//	fmt.Println(s.Age)
	p := &Person{
		name: "张三",
		age:  18,
	}
	p.print2()

}
