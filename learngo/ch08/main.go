package main

import "fmt"

type Person struct {
	name string
}

// 接收者
func (p *Person) SayHello() {
	fmt.Println("hello")
}

// 通过指针交换两个值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t

}

func changeName(p *Person) {
	p.name = "ooc"
}
func main() {
	//指针 提一个需求，我希望结构体传值的时候 我在函数中修改的值可以反应到变量中
	//p := Person{
	//	name: "iooc",
	//}
	//changeName(&p)
	//fmt.Println(p.name)
	//	var pi *Person = &p
	//	fmt.Printf("%p\r\n", pi)
	//
	//	var po *Person
	//	po = &p
	//	fmt.Printf("%p", po)
	//(*po).name
	//go 语言限制了指针的运算， 在c语言中你拿到一个指针以后 可以进行+1 在go语言中不行 不能参加运算 go语言是一个阉割版
	//在unsafe包里面，所以说一版我们不会使用unsafe包，但是当你要用到的时候是可以使用的

	//var a int = 10
	//b := &a

	//var b *int
	//fmt.Println(b)
	//var ps *Person
	//fmt.Println(ps.name)

	//ps :=&Person{} //第一种初始化方式
	var emptyPerson Person
	pi := &emptyPerson //第二种初始化方式
	//初始化两个关键字 map channel slice初始化推荐使用make方法
	//指针初始化推荐使用new函数，指针要初始化否则会出现nil pointer
	//map必须初始化

	var pp *Person = new(Person) //第三种初始化方式
	fmt.Println(pp.name)

	fmt.Println(pi.name)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
