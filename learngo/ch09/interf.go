package main

import "fmt"

/**
//class Animal{
//	birthday timestamp
//	legs int
//	mouth int
	walk (){
	}
//}
//
//duck 必须显示的指定animal 隐含的信息 duck具备了leg字段 也具备了birthday字段
//class Duck implements Animal{
//	birthday timestamp
//	legs int
//	mouth int
//}

鸭子类型强调的是事物的外部行为 而不是内部的结构 在其他语言中
*/
//接口的定义
type Duck interface {
	//方法的声明
	Gaga() string
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() string {
	return "psk"
}
func (pd *pskDuck) Walk() {
	fmt.Println("pskDuck is walking")
}
func (pd *pskDuck) Swimming() {
	fmt.Println("pskDuck is swimming")
}

func main() {
	//go语言的接口 鸭子类型 php  python
	//go语言中处处都是interface

	/**
	当看到一只鸟走起来像鸭子，游泳起来像鸭子，叫起来像鸭子，那么这只鸟就是鸭子
	动词 方法 具备某些方法
	*/

	var d Duck = &pskDuck{}
	d.Walk()

}
