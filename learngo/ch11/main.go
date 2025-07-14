package ch11

//单元测试 go test
//go test命令是按照一个一定约定和组织的测试驱动程序，在包目录中，所有以 _test.go 结尾的文件中，都会被识别为测试文件。
//我们写的以_test.go源码文件不用担心过多，因为go build命令不会将这些测试文件打包到最后的可执行文件
//test文件有4类 Test开头的 功能测试 Benchmark开头的 性能测试 Example 开头的 模糊测试
