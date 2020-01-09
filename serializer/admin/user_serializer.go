package admin

import (
	"awesomeProject/model/admin"
	"time"
)

// 用户列表
type UserList struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	Status    int    `json:"status"`
}

// 用户序列化
func UserSerializer(items []admin.User) []UserList {
	list := []UserList{}
	for _, item := range items {
		user := UserList{
			ID:        item.ID,
			UserName:  item.UserName,
			Avatar:    item.Avatar,
			CreatedAt: time.Unix(item.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			Status:    item.Status,
		}
		list = append(list, user)
	}
	return list
}
