package router

import (
	"end/handler"
	"end/internal"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, app *internal.App) {

	userHandler := handler.NewUserHandler(app)
	// 设置
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.GET("/list", userHandler.GetUser)
		userGroup.POST("", userHandler.GetUser)
		userGroup.PUT("", userHandler.GetUser)
		userGroup.DELETE("", userHandler.GetUser)
	}

}
