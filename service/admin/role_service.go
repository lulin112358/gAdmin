package admin

import (
	"awesomeProject/model"
	"awesomeProject/model/admin"
	"awesomeProject/serializer"
	admin2 "awesomeProject/serializer/admin"
	"os"
	"strconv"
)

// 角色列表字段
type RoleListService struct {
	Page int    `form:"page" json:"page"`
	Key  string `form:"key" json:"key"`
}

// 获取角色列表
func (s *RoleListService) RolesList() serializer.Response {
	perPage, _ := strconv.Atoi(os.Getenv("PER_PAGE"))
	if s.Page == 0 {
		s.Page = 1
	}

	var roles []admin.AuthGroup
	count := 0

	model.DB.Where("title like ?", "%"+s.Key+"%").Offset((s.Page - 1) * perPage).Limit(perPage).Find(&roles).Count(&count)
	if err := model.DB.Where("title like ?", "%"+s.Key+"%").Offset((s.Page - 1) * perPage).Limit(perPage).Find(&roles).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Data: serializer.BuildList(admin2.BuildRoleList(roles), count),
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
