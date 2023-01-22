package router

import (
	"end/api/handler"
	"end/bootstrap"
	"end/repository"
	"end/usecase"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, app *bootstrap.App) {
	userHandler := &handler.UserHandler{
		UserUsecase: usecase.NewUserusecase(repository.NewUserRepository(app),app.Config.Timeout),
	}
	// 设置
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.GET("/list", userHandler.GetUserList)
		userGroup.POST("", userHandler.CreateUser)
		userGroup.PUT("", userHandler.UpdateUser)
		userGroup.DELETE("/:ids", userHandler.DeleteUser)
	}

}
