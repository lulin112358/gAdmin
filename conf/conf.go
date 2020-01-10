package conf

import (
	"awesomeProject/catch"
	"awesomeProject/middleware"
	"awesomeProject/migration"
	"awesomeProject/model"
	"awesomeProject/util"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	// 读取配置
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	// 数据迁移
	migration.Migration()

	// Redis
	catch.Redis()

	// 初始化权限认证模块
	middleware.Init()
}
