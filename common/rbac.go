package common

import (
	"awesomeProject/middleware"
)

// 根据用户id获取用户角色
func GetRoleByUser(uid string) string {
	var r, _ = middleware.Auth.GetRolesForUser(uid)

	return r[0]
}

// 根据角色获取规则
func GetRuleByRole(role string) []string {
	var list []string
	rules := middleware.Auth.GetPermissionsForUser(role)
	for _, item := range rules {
		list = append(list, item[1])
	}

	return list
}

// 根据用户id获取规则
func GetRuleByUser(uid string) []string {
	return GetRuleByRole(GetRoleByUser(uid))
}
