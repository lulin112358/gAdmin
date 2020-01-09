package migration

import (
	"awesomeProject/model"
	"awesomeProject/model/admin"
)

// 数据迁移
func Migration() {
	model.DB.AutoMigrate(&admin.User{}, &admin.AuthRule{}, &admin.AuthGroup{}, &admin.AuthGroupAccess{})
}
