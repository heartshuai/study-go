package main

import "fmt"

func main() {
	//go语言提供了哪些集合类型的数据结构 数组 切片(slice) map list
	//数组 定义 var name [count]int
	//var courses1 [3]string //courses 类型 数组 只有三个元素的数组
	//// []string 和[3] string  是两种不同的类型
	////var courses2 [4]string
	//courses1[0] = "go"
	//courses1[1] = "gprc"
	//courses1[2] = "gin"
	//
	//fmt.Println(courses1)
	//fmt.Printf("%T", courses1)
	//fmt.Printf("%T", courses2)

	//for _, value := range courses1 {
	//	fmt.Println(value)
	//}
	//
	//var courses2 = [3]string{"go", "gprc", "gin"}
	//fmt.Println(courses2)
	//
	//course4 := [3]string{"go", "gprc", "gin"}
	//for _, value := range course4 {
	//	fmt.Println(value)
	//}

	//course2 := [3]string{2: "gin"}
	//for _, value := range course2 {
	//	fmt.Println(value)
	//}
	//
	//course3 := [...]string{"go", "gprc", "gin"}
	//for _, value := range course3 {
	//	fmt.Println(value)
	//}
	//
	//course4 := [...]string{"go", "gprc", "gin"}
	//if course3 == course4 {
	//	fmt.Println("相等")
	//} else {
	//	fmt.Println("不相等")
	//}

	//多维数组
	var coursesInfo [3][4]string
	coursesInfo[0] = [4]string{"go", "1h", "bobby", "go体系课"}
	coursesInfo[1] = [4]string{"gprc", "2h", "bobby2", "gprc入门"}
	coursesInfo[2] = [4]string{"gin", "3h", "bobby3", "gin入门"}

	//for _, value := range coursesInfo {
	//	for _, value2 := range value {
	//		fmt.Println(value2)
	//	}
	//
	//}
	for i := 0; i < len(coursesInfo); i++ {
		for j := 0; j < len(coursesInfo[i]); j++ {
			fmt.Print(coursesInfo[i][j] + " ")
		}
		fmt.Println()
	}
	//coursesInfo[0][0]="go"
}
