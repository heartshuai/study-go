package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Language struct {
	//gorm.Model
	Name    string
	AddTime sql.NullTime //每个记录创建的时候自动加入新增的时间
}

//func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
//	l.AddTime = time.Now()
//	return
//}

//func (Language) TableName() string {
//	return "you_language"
//}

/**
在gorm中可以通过给某一个struct添加一个tableName方法来自定义表名
我们自己定义表明
统一的给所有的表明加入前缀
*/

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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop_",
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "Go",
	})
}
