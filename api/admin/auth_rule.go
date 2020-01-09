package admin

import (
	"awesomeProject/serializer"
	"awesomeProject/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 菜单列表渲染
func RenderMenuList(c *gin.Context) {
	list := admin.MenuList()
	c.HTML(http.StatusOK, "menu_list.html", map[string]interface{}{
		"list": list.Data,
	})
}

// 添加菜单
func MenuAdd(c *gin.Context) {
	var service admin.MenuAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MenuAdd()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		})
	}
}

// 删除菜单
func MenuDel(c *gin.Context) {
	var service admin.MenuDelService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MenuDel()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		})
	}
}

// 修改菜单状态
func MenuStatus(c *gin.Context) {
	var service admin.MenuStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MenuStatus()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		})
	}
}

// 修改菜单排序
func MenuSort(c *gin.Context) {

}
