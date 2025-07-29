package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/:id/*action", goodsDetail)
		goodsGroup.POST("/add", createGoods)
	}
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
