package admin

import (
	"awesomeProject/model/admin"
	"strings"
	"time"
)

type MenuList struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	Sort      int    `json:"sort"`
	Pid       int    `json:"pid"`
	Lvl       int    `json:"lvl"`
	Leftpin   int    `json:"leftpin"`
	Lefthtml  string `json:"lefthtml"`
}

// 构建菜单返回
func BuildMenuList(items []admin.AuthRule) []MenuList {
	menuList := []MenuList{}
	for _, item := range items {
		menu := MenuList{
			ID:        item.ID,
			Title:     item.Title,
			Name:      item.Name,
			Status:    item.Status,
			CreatedAt: time.Unix(item.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			Sort:      item.Sort,
			Pid:       item.Pid,
		}

		menuList = append(menuList, menu)
	}

	return buildMenuList(menuList, "— — ", 0, 0, 0)
}

// 构建样式
func buildMenuList(items []MenuList, left string, pid int, lvl int, leftpin int) []MenuList {
	list := []MenuList{}

	for _, item := range items {
		if item.Pid == pid {
			item.Lvl = lvl + 1
			item.Leftpin = leftpin + 0
			item.Lefthtml = strings.Repeat(left, lvl)
			list = append(list, item)

			list = append(list, buildMenuList(items, left, int(item.ID), lvl+1, leftpin+20)...)
		}
	}
	return list
}
