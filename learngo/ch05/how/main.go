package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

type slice struct {
	array unsafe.Pointer //用来存储实际数据的数组指针 指向一块连续的内存
	len   int            //切片中元素的数量
	cap   int            //array数组的长度
}

func main() {
	//ms := slice{
	//	array: unsafe.Pointer(&ms),
	//	len:   3,
	//	cap:   3,
	//}

	//go的slice在函数传递的时候是 值传递 还是引用传递 值传递 效果又呈现出了引用的效果
	//courseSlice := []string{"go", "gprc", "gin", "mysql", "elasticsearch"}
	//printSlice(courseSlice)
	//fmt.Println(courseSlice)

	//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := data[1:6]
	//s2 := data[2:7]
	////s2 = append(s2, 1, 2, 3, 4, 5, 6, 67, 8, 9, 9, 9, 9, 9, 9, 9, 1)
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(len(s2), cap(s2))
	//s2[0] = 22
	//fmt.Println(s1, s2)
	var data []int
	for i := 0; i < 2000; i++ {
		data = append(data, i)
		fmt.Println(len(data), cap(data))
	}
}
