package handler

import (
	"end/model"
	"end/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type StudentHandler struct {
	StudentService    service.StudentServiceImpl
	contextTimeout time.Duration
}


func (h *StudentHandler) GetStudent(c *gin.Context) any{

	id := c.Param("id")
	return ApiResp2(h.StudentService.GetStudentById(c,id))
}

func (h *StudentHandler) GetStudentList(c *gin.Context) any{

	pageSizeStr := c.Query("pageSize")
	pageNoStr := c.Query("pageNo")

	pageSize, _ := strconv.Atoi(pageSizeStr)
	pageNo, _ := strconv.Atoi(pageNoStr)

	var student model.Student
	err := c.ShouldBindQuery(&student)
	if err != nil {
		return ApiResp(UNKNOWN,err)
	}

	return ApiResp2(h.StudentService.GetStudentList(c, model.PageDomain[model.Student]{
		PageSize: pageSize,
		PageNo: pageNo,
		Condition: &student,
	}))
}

func (h *StudentHandler) CreateStudent(c *gin.Context) any{

	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		return ApiResp(UNKNOWN,err)
	}
	return ApiResp2(h.StudentService.InsertStudent(c,student))
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) any{

	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err!=nil{
		return ApiResp(UNKNOWN,err)
	}
	return ApiResp2(h.StudentService.UpdateStudent(c,student))
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) any{

	ids := c.Param("ids")
	return ApiResp2(h.StudentService.DeleteStudent(c,ids))
}