//package main
//
////
////import (
////	"net/http"
////
////	"github.com/gin-gonic/gin"
////)
////
////// go modules 添加依赖 删除未使用的依赖项
////func main() {
////	r := gin.Default()
////	r.GET("/ping", func(c *gin.Context) {
////		c.JSON(http.StatusOK, gin.H{
////			"message": "pong",
////		})
////	})
////	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
////}
//
//import (
//	"context"
//	"fmt"
//
//	"github.com/redis/go-redis/v9"
//)
//
//var ctx = context.Background()
//
//func ExampleClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}

package main

import (
	//1.go自带的包 2.第三方的包 3.自己内部的包
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//go install 升级go get -u 升级到最新的次要版本或修订版本
	//go get -u=patch 升级到最新的修订版本
	//go get xxx 版本号 带@1.0 go get 会修改go.mod文件
	//一开始的时候 我写了一个A项目，仓库是project-A 但是我的代码仓库的go.mod中设置的是github.com/boby/A
	//B项目由于依赖A项目，import的是github.com/boby/project-A go get命令的时候 由于package和代码仓库的名称不一样
	//replace
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
