package admin

import (
	"awesomeProject/serializer"
	"awesomeProject/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 角色列表渲染
func RenderRoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "role_list.html", gin.H{})
}

// 获取角色列表
func RolesList(c *gin.Context) {
	var service admin.RoleListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RolesList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		})
	}
}

// 添加角色渲染
func RenderRoleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "role_add.html", gin.H{})
}

// 添加角色
func RoleAdd(c *gin.Context) {
	var service admin.RoleAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RoleAdd()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		})
	}
}
