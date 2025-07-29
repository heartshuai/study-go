package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)
	router.POST("/post", getPost)
	router.Run()
}

func getPost(c *gin.Context) {

}

func formPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123456")
	c.JSON(http.StatusOK, gin.H{
		"message":  "form post",
		"username": username,
		"password": password,
	})

}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("first_name", "Guest")
	lastName := c.Query("last_name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome " + firstName + " " + lastName,
	})
}
