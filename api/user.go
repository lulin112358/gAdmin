package api

import (
	"awesomeProject/serializer"
	service2 "awesomeProject/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func Register(c *gin.Context) {
	var service service2.UserRegister
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Msg:  err.Error(),
		})
	}
}

// 用户登陆页面
func RenderLogin(c *gin.Context) {
	s := sessions.Default(c)
	uid := s.Get("user_id")
	if uid != nil {
		c.Redirect(302, "/api/index/index")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

// 用户登陆
func Login(c *gin.Context) {
	var service service2.UserLogin
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "登陆失败",
			Error: err.Error(),
		})
	}
}

// 退出登陆
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, serializer.Response{
		Code: 1,
		Msg:  "退出成功",
	})
}

// 首页渲染
func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 用户列表页渲染
func RenderUserList(c *gin.Context) {
	c.HTML(http.StatusOK, "user_list.html", gin.H{})
}

// 获取用户列表数据
func GetUserList(c *gin.Context) {
	var service service2.UserList
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetUserList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Msg:  "参数绑定错误",
		})
	}
}

// 渲染用户添加
func RenderUserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "user_add.html", gin.H{})
}

// 渲染用户修改
func RenderUserEdit(c *gin.Context) {
	id := c.Param("id")
	data := service2.GetUserInfo(id)
	c.HTML(http.StatusOK, "user_edit.html", gin.H{
		"data": data.Data,
	})
}

// 用户修改
func UserEdit(c *gin.Context) {
	var service service2.UserEdit
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Msg:  err.Error(),
		})
	} else {
		res := service.UserEdit()
		c.JSON(http.StatusOK, res)
	}
}

// 删除用户
func UserDel(c *gin.Context) {
	var service service2.UserDel
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserDel()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   err.Error(),
			Error: err.Error(),
		})
	}
}
