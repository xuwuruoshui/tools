package router

import (
	"end/api/handler"
	"end/bootstrap"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine, app *bootstrap.App) {
	// 用户相关
	wrapper := handler.NewRouterPlus(r)
	Ws(wrapper,app)
	User(wrapper ,app)
	Student(wrapper,app)
}
