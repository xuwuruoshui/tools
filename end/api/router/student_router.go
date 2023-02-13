package router

import (
	"end/api/handler"
	"end/bootstrap"
	"end/repository"
	"end/service"
)

func Student(r *handler.RouterWrapper, app *bootstrap.App) {
	studentHandler := &handler.StudentHandler{
		StudentService: service.NewStudentService(repository.NewStudentRepository(app),app.Config.Timeout),
	}
	// 设置
	studentGroup := r.Group("/student")
	{
		studentGroup.GET("/:id", studentHandler.GetStudent)
		studentGroup.GET("/list", studentHandler.GetStudentList)
		studentGroup.POST("", studentHandler.CreateStudent)
		studentGroup.PUT("", studentHandler.UpdateStudent)
		studentGroup.DELETE("/:ids", studentHandler.DeleteStudent)
	}

}