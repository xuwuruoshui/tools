package router

import (
	"end/internal"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine, app *internal.App) {
	// 用户相关
	User(r, app)
}
