package middleware

import (
	"awesomeProject/model"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

var Auth *casbin.Enforcer

// 初始化权限认证
func Init() {
	a, _ := gormadapter.NewAdapterByDB(model.DB)
	Auth, _ = casbin.NewEnforcer("conf/rbac.conf", a)
}
