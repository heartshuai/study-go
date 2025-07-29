package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Gin",
		})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"title": "Gin",
		})
	})
	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{
			"title": "Gin",
		})
	})
	router.Run(":8080")
}
