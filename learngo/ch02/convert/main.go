package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var a int8 = 12
	//var b = uint8(a)
	//var f float32 = 3.14
	//var f64 = float64(a)
	//var c = int32(f)

	//fmt.Println(b, c, f64)

	//type IT int //类型要求很严格
	//var c IT = IT(a)

	//字符串转数字

	var istr = "12"

	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	var myi = 32
	strconv.Itoa(myi)
	fmt.Println(myi)

	parseFloat, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(parseFloat)

	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	parseBool, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("parse bool error")
	}
	fmt.Println(parseBool)

	//基本类型转字符串
	boolStr := strconv.FormatBool(true)
	fmt.Println(boolStr)

	floatStr := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(floatStr)

	fmt.Println(strconv.FormatInt(42, 16))
}
