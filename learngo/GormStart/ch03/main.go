package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
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
	// 迁移 schema
	//db.AutoMigrate(&User{}) //此处应该有sql语句
	user := User{
		Name: "yxs2",
	}
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	//fmt.Println(user.Error)
	//db.Model(&User{ID: 1}).Update("name", "")
	//updates不会更新零值 update会更新
	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})
	/**
	解决仅更新非零值的方法有两种
	1、将string 字段设置为 *string
	2.使用sql的NullString来解决
	如果你只需要判断“这个字段有没有要更新”，用 *string 更简单。

	如果你需要明确区分“空字符串” 和 “NULL”，那就用 sql.NullString。

	如果是表单提交或 JSON 接收，*string 更好整合；sql.NullString 需要额外的序列化支持。
	*/

}
