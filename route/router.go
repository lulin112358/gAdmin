package route

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// admin路由
	AdminRouter(r)

	// api路由
	ApiRouter(r)

	return r
}
