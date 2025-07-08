package main

import "fmt"

func main() {
	// map 是一个key(索引) 和value (值) 无序的集合 主要是查询方便
	var courseMap = map[string]string{
		"go":    "golang",
		"gprc":  "gprc的咕咕咕",
		"gin":   "gin的咕咕咕",
		"mysql": "mysql的咕咕咕",
	}
	//
	//courseMap["mysql"] = "mysql的咕咕咕"
	//fmt.Println(courseMap)
	//var courseMap = map[string]string{}
	//
	////var courseMap = make(map[string]string)
	//courseMap["go"] = "golang"
	//fmt.Println(courseMap)
	////map 必须初始化才能使用 1.map[string]string{} 2. make(map[string]string ,3)
	//
	//var m []string
	//m = append(m, "go")
	//for k, v := range courseMap {
	//	fmt.Println(k, v)
	//}
	//for key, value := range courseMap {
	//	fmt.Println(key, value)
	//}
	//map是无序的 而且不保证每次打印都是相同的顺序 想要有序不能使用map

	//if d, ok := courseMap["java"]; !ok {
	//	fmt.Println("not find")
	//} else {
	//	fmt.Println(d)
	//}
	delete(courseMap, "go")
	//很重要的提示，map不是线程安全的
	delete(courseMap, "java")
	fmt.Println(courseMap)

}
