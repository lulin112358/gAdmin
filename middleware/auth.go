package middleware

import (
	"awesomeProject/model/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取当前登陆用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			// 获取用户信息
			user, err := admin.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// 验证登陆
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*admin.User); ok {
				c.Next()
				return
			}
		}
		c.Redirect(302, "/admin/user/login")
		c.Abort()
	}
}
