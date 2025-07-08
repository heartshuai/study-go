package main

import "fmt"

func main() {
	//switch 1 {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2")
	//default:
	//	fmt.Println("default")
	//}

	//中文的星期几输出英文
	day := "星期一"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	case "星期四":
		fmt.Println("Thursday")
	case "星期五":
		fmt.Println("Friday")
	case "星期六":
		fmt.Println("Saturday")
	default:
		fmt.Println("Unknown")
	}

	score := 60
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
