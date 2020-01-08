package route

import (
	"awesomeProject/api"
	"awesomeProject/common"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine  {
	r := gin.Default()
	// session
	r.Use(middleware.Sessions(os.Getenv("SESSION_SECRET")))
	// 跨域中间件
	r.Use(middleware.Cors())
	// 获取当前用户
	r.Use(middleware.CurrentUser())

	// 路由
	g := r.Group("/api")
	{
		// 不需要登陆
		// 用户登陆
		g.POST("user/dologin", api.Login)

		// 登陆页渲染
		g.GET("user/login", api.RenderLogin)

		auth := g.Group("")
		// 需要登陆
		auth.Use(middleware.AuthRequired())
		{
			// 上传文件
			auth.POST("common/upload", common.Upload)

			// 退出登陆
			auth.POST("user/logout", api.Logout)

			// 首页渲染
			auth.GET("index/index", api.RenderIndex)

			// 用户列表渲染
			auth.GET("user/listpage", api.RenderUserList)

			// 获取用户列表
			auth.GET("user/list", api.GetUserList)

			// 用户注册
			auth.POST("user/register", api.Register)

			// 用户注册渲染
			auth.GET("user/registerpage", api.RenderUserAdd)

			// 用户修改渲染
			auth.GET("user/editpage/:id", api.RenderUserEdit)

			// 用户修改信息
			auth.POST("user/edit", api.UserEdit)

			// 用户删除
			auth.POST("user/del", api.UserDel)
		}
	}

	return r
}