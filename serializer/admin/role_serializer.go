package admin

import (
	"awesomeProject/model/admin"
	"time"
)

// 角色列表
type RoleList struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    int    `json:"status"`
}

// 构建角色列表
func BuildRoleList(items []admin.AuthGroup) []RoleList {
	list := []RoleList{}
	for _, item := range items {
		role := RoleList{
			ID:        item.ID,
			Title:     item.Title,
			CreatedAt: time.Unix(item.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Unix(item.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			Status:    item.Status,
		}

		list = append(list, role)
	}

	return list
}
