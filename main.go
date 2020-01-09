package main

import (
	"awesomeProject/conf"
	"awesomeProject/route"
)

func main() {
	// 读取配置(初始化数据库连接等)
	conf.Init()

	// 路由
	r := route.NewRouter()
	// 加载静态资源
	r.Static("/static", "./static")
	// 加载html模板
	r.LoadHTMLGlob("views/**/*")
	// 运行
	r.Run(":3000")
}
