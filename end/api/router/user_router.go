package router

import (
	"end/api/handler"
	"end/bootstrap"
	"end/repository"
	"end/service"
)

func User(r *handler.RouterWrapper, app *bootstrap.App) {
	userHandler := &handler.UserHandler{
		UserUsecase: service.NewUserService(repository.NewUserRepository(app),app.Config.Timeout),
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
