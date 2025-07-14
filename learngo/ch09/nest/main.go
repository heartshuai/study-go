package main

import "fmt"

type MyReader interface {
	Write(string)
}
type MyWriter interface {
	Read() string
}

type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWrite()
}

type SreadWriter struct {
}

func (s *SreadWriter) Write(s2 string) {
	//TODO implement me
	fmt.Println("write me")
}

func (s *SreadWriter) Read() string {
	//TODO implement me
	fmt.Println("read me")
	return ""
}

func (s *SreadWriter) ReadWrite() {
	//TODO implement me
	fmt.Println("read write")
}

func main() {
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Read()
}
