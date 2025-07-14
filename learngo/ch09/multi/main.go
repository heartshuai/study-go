package main

import "fmt"

type MyWriter interface {
	Write(string) error
}
type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter
	//MyCloser //interface 也是一个类型 想放入一个写文件的实现，我想放一个写数据库的实现

}
type fileWriter struct {
	filepath string
}

type databaseWriter struct {
	host string
	port int
	db   string
}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write file string")
	return nil
}
func (fw *databaseWriter) Write(string) error {
	fmt.Println("write database string")
	return nil
}

//	func (wc *writerCloser) Write(string) error {
//		fmt.Println("write string")
//		return nil
//	}
func (wc *writerCloser) Close() error {
	fmt.Println("close")
	return nil
}
func main() {
	var fw MyWriter = &writerCloser{
		&fileWriter{},
	}
	//var mc MyCloser = &writerCloser{}
	fw.Write("hello")
	//mc.Close()
}
