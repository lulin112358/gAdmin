package admin

import "github.com/jinzhu/gorm"

type AuthGroupAccess struct {
	gorm.Model
	Uid     int
	GroupId int
}
