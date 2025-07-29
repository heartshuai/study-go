package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/**
	优雅退出，当我们关闭程序的时候应该做的后续处理
	微服务 启动之前 或者启动之后会做一件事，将当前的服务的ip地址和端口注册到注册中心
	当服务停止的时候，并没有告知注册中心
	*/
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	go func() {
		router.Run(":8080")
	}()
	//如果想要接收到信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("关闭serve。。。")
	fmt.Println("优雅的关闭")

}
