package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func main() {
	//1.实例化一个server

	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}

}
