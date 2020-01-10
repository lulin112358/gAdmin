package route

import (
	"awesomeProject/api/admin"
	"awesomeProject/common"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// admin路由
func AdminRouter(r *gin.Engine) {
	// session
	r.Use(middleware.Sessions(os.Getenv("SESSION_SECRET")))
	// 跨域中间件
	r.Use(middleware.Cors())
	// 获取当前用户
	r.Use(middleware.CurrentUser())

	// 路由
	g := r.Group("/admin")
	{
		// 不需要登陆
		// 用户登陆
		g.POST("user/dologin", admin.Login)

		// 登陆页渲染
		g.GET("user/login", admin.RenderLogin)

		auth := g.Group("")
		// 需要登陆
		auth.Use(middleware.AuthRequired())
		{
			// 上传文件
			auth.POST("common/upload", common.Upload)

			// 退出登陆
			auth.POST("user/logout", admin.Logout)

			// 首页渲染
			auth.GET("index/index", admin.RenderIndex)

			// 用户列表渲染
			auth.GET("user/listpage", admin.RenderUserList)

			// 获取用户列表
			auth.GET("user/list", admin.GetUserList)

			// 用户注册
			auth.POST("user/register", admin.Register)

			// 用户注册渲染
			auth.GET("user/registerpage", admin.RenderUserAdd)

			// 用户修改渲染
			auth.GET("user/editpage/:id", admin.RenderUserEdit)

			// 用户修改信息
			auth.POST("user/edit", admin.UserEdit)

			// 用户删除
			auth.POST("user/del", admin.UserDel)

			// 菜单列表渲染
			auth.GET("index/menu_list", admin.RenderMenuList)

			// 添加菜单
			auth.POST("menu/menu_add", admin.MenuAdd)

			// 删除菜单
			auth.POST("menu/menu_del", admin.MenuDel)

			// 修改菜单状态
			auth.POST("menu/menu_status", admin.MenuStatus)

			// 菜单排序
			auth.POST("menu/menu_sort", admin.MenuSort)

			// 修改菜单渲染
			auth.GET("menu/editpage/:id", admin.RenderMenuEdit)

			// 修改菜单
			auth.POST("menu/menu_edit", admin.MenuEdit)

			// 角色列表渲染
			auth.GET("role/listpage", admin.RenderRoleList)

			// 获取角色列表
			auth.GET("role/list", admin.RolesList)
		}
	}
}
