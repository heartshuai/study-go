package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123")
		c.Next()
		end := time.Since(t)
		fmt.Println(end)
		status := c.Writer.Status()
		fmt.Println(status)
	}
}
func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("x-token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
func main() {
	r := gin.Default()
	//使用logger中间件
	r.Use(TokenRequired())
	//写入header
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//authrized := r.Group("/goods")
	//authrized.Use(AuthRequired)
	r.Run(":8080")
}

func AuthRequired(context *gin.Context) {
	fmt.Println("1122")
}
