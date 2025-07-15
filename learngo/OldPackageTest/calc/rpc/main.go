package main

func add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name string
	Com  Company
}
type PrintResult struct {
	Info string
	Err  error
}

//func RpcPrintln(employee Employee) PrintResult {
//rpc中的第二个点 传输协议 数据编码协议
//http协议1.0 2.0协议
//http协议底层使用的也是tcp http现在主流的是1.x 这种协议呢有性能问题 一次性 一旦结果返回 连接就断开
//1.直接自己基于tcp/udp协议去封装一层协议 myhttp 没有通用性 http2.0 既有http的特性 也有长连接的特性

/**
客户端
	1.建立连接 tcp/http
	2.将employee对象序列化成json 序列化
	3.发送json数据给远程的计算机 - 调用成功后 实际上你接收到的是一个二进制的数据
	4.等待服务器发送结果
	5.将服务器返回的数据解析成PrintResult对象-反序列化
服务端
	1.监听网络端口 80
	2.读取数据-二进制的json数据
	3.对数据进行反序列化Employee对象
	4.开始业务逻辑
	5.将处理的结果PrintResult序列化json 二进制数据
	6.将数据返回
	序列化和反序列化是可以选择的，不一定要采用json xml protobuf msgpack
*/

//}

func main() {
	//现在我们想把Add函数变成一个远程的函数调用，也就意味着我们需要把Add函数的实现放在一个远程的计算机上。
	/**
	我们原本的电商系统，这里地方有一个逻辑，这个逻辑是扣减库存，但是库存服务是一个独立的系统，reduce，那如何调用一定会牵扯到网络，做成一个web服务(gin/beego/httpserver)
	1.这个函数的调用参数如何传递-json (json是一种数据格式的协议) /xml/protobuf/msgpack -编码协议 json 并不是一个高效的协议
		现在网络调用有两个端 -客户端 应该干嘛 将数据传输到gin
		gin-服务端 服务端负责解析数据
	2.
	*/
	//将这个打印的工作放在另一台服务器上，我需要将本地的内存对象struct 这样不行，可行的方式就是将struct序列化成json
	//fmt.Println(Employee{"张三", Company{"百度", "北京"}})

	//远程的服务需要将二进制对象反解成struct对象
	//搞这么麻烦，直接全部使用json去格式化不香么，这种做法在浏览器和gin服务之间是可以行，但是如果你是一个大型的分布式系统

}
