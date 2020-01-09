package admin

import (
	"awesomeProject/model"
	"awesomeProject/model/admin"
	"awesomeProject/serializer"
	admin2 "awesomeProject/serializer/admin"
)

// 菜单添加字段
type MenuAddService struct {
	Title  string `form:"title" json:"title" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	Sort   int    `form:"sort" json:"sort" binding:"required"`
	Status int    `form:"status" json:"status" binding:"required"`
	Pid    int    `form:"pid" json:"pid"`
}

// 菜单删除字段
type MenuDelService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// 修改菜单状态字段
type MenuStatusService struct {
	ID     uint `form:"id" json:"id" binding:"required"`
	Status int  `form:"status" json:"status" binding:"required"`
}

// 修改菜单排序字段
type MenuSortService struct {
	ID   uint `form:"id" json:"id" binding:"required"`
	Sort int  `form:"sort" json:"sort" binding:"required"`
}

// 菜单列表
func MenuList() serializer.Response {
	var rules []admin.AuthRule

	if err := model.DB.Order("sort desc").Find(&rules).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Data: admin2.BuildMenuList(rules),
			Msg:  "success",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Msg:   "error",
			Error: err.Error(),
		}
	}
}

// 添加菜单
func (s *MenuAddService) MenuAdd() serializer.Response {
	menu := admin.AuthRule{
		Name:   s.Name,
		Title:  s.Title,
		Status: s.Status,
		Sort:   s.Sort,
		Pid:    s.Pid,
	}
	if err := model.DB.Create(&menu).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "添加成功",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Msg:   "添加失败",
			Error: err.Error(),
		}
	}
}

// 删除菜单
func (s *MenuDelService) MenuDel() serializer.Response {
	var menu admin.AuthRule
	if err := model.DB.Delete(&menu, s.ID).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "删除成功",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Msg:   "删除失败",
			Error: err.Error(),
		}
	}
}

// 修改菜单状态
func (s *MenuStatusService) MenuStatus() serializer.Response {
	if s.Status == 1 {
		s.Status = 2
	} else {
		s.Status = 1
	}
	if err := model.DB.Model(&admin.AuthRule{}).Where("id = ?", s.ID).Update("status", s.Status).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "修改成功",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Msg:   "修改失败",
			Error: err.Error(),
		}
	}
}
