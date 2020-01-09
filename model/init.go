package model

import (
	"awesomeProject/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

// 连接数据库
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)

	// Error
	if err != nil {
		util.Log().Panic("数据库连接失败", err)
	}

	// 设置空闲
	db.DB().SetMaxIdleConns(50)

	// 设置最大连接数
	db.DB().SetMaxOpenConns(100)

	// 设置连接超时时间
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}
