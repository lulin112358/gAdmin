package conf

import (
	"awesomeProject/catch"
	"awesomeProject/model"
	"awesomeProject/util"
	"github.com/joho/godotenv"
	"os"
)

func Init()  {
	// 读取配置
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// Redis
	catch.Redis()
}