package handler

import (
	"end/model"
	"end/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserHandler struct {
	UserUsecase    service.UserServiceImpl
	contextTimeout time.Duration
}


func (h *UserHandler) GetUser(c *gin.Context) any{

	id := c.Param("id")
	return ApiResp2(h.UserUsecase.GetUserById(c,id))
}

func (h *UserHandler) GetUserList(c *gin.Context) any{

	pageSizeStr := c.Query("pageSize")
	pageNoStr := c.Query("pageNo")

	pageSize, _ := strconv.Atoi(pageSizeStr)
	pageNo, _ := strconv.Atoi(pageNoStr)

	var user model.User
	err := c.ShouldBindQuery(&user)
	if err != nil {
		return ApiResp(UNKNOWN,err)
	}

	return ApiResp2(h.UserUsecase.GetUserList(c, model.PageDomain[model.User]{
		PageSize: pageSize,
		PageNo: pageNo,
		Condition: &user,
	}))
}

func (h *UserHandler) CreateUser(c *gin.Context) any{

	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return ApiResp(UNKNOWN,err)
	}
	return ApiResp2(h.UserUsecase.InsertUser(c,user))
}

func (h *UserHandler) UpdateUser(c *gin.Context) any{

	var user model.User
	err := c.ShouldBindJSON(&user)
	if err!=nil{
		return ApiResp(UNKNOWN,err)
	}
	return ApiResp2(h.UserUsecase.UpdateUser(c,user))
}

func (h *UserHandler) DeleteUser(c *gin.Context) any{

	ids := c.Param("ids")
	return ApiResp2(h.UserUsecase.DeleteUser(c,ids))
}