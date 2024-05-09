package dao

import (
	"time"
	"web_gin/config"
	"web_gin/pkg/logger"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		//panic("failed to connect database")
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		//panic("failed to connect database")
		logger.Error(map[string]interface{}{"database  connect error": err.Error()})
	}
	fmt.Println("连接数据库成功")
	Db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	Db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	Db.DB().SetConnMaxLifetime(time.Hour)
	// userd := "user"
	// result := Db.Find(&userd)

	// // 执行查询
	// fmt.Println("找到数据:", result)
	// fmt.Println("找到数据:", result.RowsAffected)

}
