package main

import "fmt"

func add(a, b interface{}) interface{} {
	//return a + b, nil
	//ai, ok := a.(int)
	//if !ok {
	//	panic("not int type")
	//}
	//bi, ok := b.(int)
	//if !ok {
	//	panic("not int type")
	//}
	//return ai + bi
	switch a.(type) {
	case int:
		ai := a.(int)
		bi := b.(int)
		return ai + bi
	case int32:
		ai := a.(int32)
		bi := b.(int32)
		return ai + bi
	case int64:
		ai := a.(int64)
		bi := b.(int64)
		return ai + bi
	case float32:
		ai := a.(float32)
		bi := b.(float32)
		return ai + bi
	case float64:
		ai := a.(float64)

		bi := b.(float64)
		return ai + bi
	default:
		panic("not int type")

	}
}
func main() {
	//a, err := add(1, 2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(a)
	fmt.Println(add(1.0, 2.0))

}
