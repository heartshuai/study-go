package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var mylist list.List

	mylist := list.New()
	mylist.PushBack("go")
	mylist.PushBack("mysql")
	mylist.PushBack("gprc")
	mylist.PushFront("ggg")
	//fmt.Println(mylist)

	//遍历打印值 正序
	e := mylist.Front()
	for ; e != nil; e = e.Next() {
		if e.Next().Value.(string) == "gprc" {
			break
		}

	}
	//mylist.InsertBefore("gin", e)
	mylist.Remove(e)
	//fmt.Println(mylist)
	for e := mylist.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	//集合类型 4种 1.数组 不同长度的数组类型不一样 2.切片 -动态数组 用起来很方便 而且性能高 尽量使用 3.map 4.list 用的少 链表

}
