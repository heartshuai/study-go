package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int    `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err == nil {
			c.JSON(200, gin.H{
				"id":   person.ID,
				"name": person.Name,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "error",
			})
		}
	})
	router.Run()
}

func createGoods(context *gin.Context) {

}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	action := context.Param("action")
	context.JSON(200, gin.H{
		"message": "goods detail",
		"id":      id,
		"action":  action,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "goods list",
	})
}
