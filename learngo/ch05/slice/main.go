package main

import "fmt"

func main() {
	//go 折中 slice切片 - 数组 弱化
	//var courses []string
	//fmt.Printf("%T \r\n", courses)
	////这个方法很特别 第一次用很头疼 原理：slice切片底层是数组
	//courses = append(courses, "go")
	//courses = append(courses, "go")
	//fmt.Println(courses)
	//
	////切片的初始化 3种 1:从数组直接创建 2:使用string{} 3：make
	//allCourses := [5]string{"go", "gprc", "gin", "mysql", "elasticsearch"}
	//courseSlice := allCourses[:] //左闭右开
	//fmt.Println(courseSlice)

	//courseSlice = []string{"go", "gprc", "gin", "mysql", "elasticsearch"}
	//make 函数
	//courseSlice2 := make([]string, 5)
	//courseSlice2[0] = "go"
	//fmt.Println(courseSlice2)
	//
	//courseSlice := []string{"go", "gprc"}
	//courseSlice2 := []string{"12", "456"}
	//courseSlice = append(courseSlice, courseSlice2[1:]...)
	//courseSlice = append(courseSlice, "go", "mysql", "es", "bs")

	//如何删除切片
	courseSlice := []string{"go", "gprc", "gin", "mysql", "elasticsearch"}
	//myslice := append(courseSlice[:2], courseSlice[3:]...)
	//fmt.Println(myslice)

	//courseSliceCopy := courseSlice
	//courseSliceCopy2 := courseSlice[:]
	//fmt.Println(courseSliceCopy)
	//fmt.Println(courseSliceCopy2)

	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy2)
	var courseSliceCopy = make([]string, len(courseSlice))
	copy(courseSliceCopy, courseSlice)
	fmt.Println(courseSliceCopy)

	fmt.Println("=======================")
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy)

}
