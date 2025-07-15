package handler

// HelloServiceName 名称冲突的问题
const HelloServiceName = "handler/HelloService"

// 我们关心的是NewHelloService这个名字呢 还是这个结构体中的方法
type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
