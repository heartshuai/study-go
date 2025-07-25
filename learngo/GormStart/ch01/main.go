package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "gorm_test:bdTmwwTcENMziceE@tcp(152.136.246.107:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	/**
	设置全局的logger 这个logger在我们执行每个sql语句的时候会打印每一行sql
	sql才是最重要的，本着这个原则尽量看到每个api背后的sql语句是什么样的
	定义一个表结构 将表结构直接生成对应的表 -migrations
	*/
	// 迁移 schema
	//db.AutoMigrate(&Product{}) //此处应该有sql语句

	// Create
	db.Create(&Product{Code: sql.NullString{
		String: "D42",
		Valid:  true,
	}, Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{
		String: "",
		Valid:  true,
	}}) // 仅更新非零值字段
	//如果我们去更新一个product 只设置了price :200
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product 逻辑删除
	//db.Delete(&product, 1)

}
