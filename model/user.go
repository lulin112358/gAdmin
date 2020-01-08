package model

import "github.com/jinzhu/gorm"

// 用户表定义
type User struct {
	gorm.Model
	UserName string
	Password string
	Avatar   string
	Status   int
}

// 获取用户信息
func GetUser(uid interface{}) (User, error) {
	var user User
	result := DB.First(&user, uid)
	return user, result.Error
}
