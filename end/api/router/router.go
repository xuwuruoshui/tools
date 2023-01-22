package router

import (
	"end/bootstrap"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine, app *bootstrap.App) {
	// 用户相关
	User(r, app)
}
