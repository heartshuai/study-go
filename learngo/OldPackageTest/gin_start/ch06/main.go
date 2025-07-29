package main

import (
	"github.com/gin-gonic/gin"
	"learngo/OldPackageTest/gin_start/ch06/proto"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJson)
	router.GET("returnProto", returnProto)
	router.Run()
}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := &proto.Teacher{
		Name:   "Lena",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}
func moreJson(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	c.JSON(http.StatusOK, msg)
}
