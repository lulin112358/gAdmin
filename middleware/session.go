package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 初始化session
func Sessions(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	return sessions.Sessions("singo-session", store)
}