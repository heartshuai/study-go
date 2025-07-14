package main

import "fmt"

func mPrint(data ...interface{}) {
	for _, v := range data {
		fmt.Println(v)
	}
}
func mPrint2(data interface{}) {
	fmt.Println(data)
}

type myInfo struct {
}

func (m *myInfo) Error() string {
	return "this is an error"
}
func main() {
	//var data = []interface{}{
	//	"body1",
	//	1,
	//	"body3",
	//}
	//var data = []string{
	//	"body1",
	//	"body2",
	//	"body3",
	//}
	////不可以
	//
	//var datai []interface{}
	//for _, v := range data {
	//	datai = append(datai, v)
	//}
	//
	//mPrint(datai...)
	var err error
	err = &myInfo{}
	err.Error()

}
