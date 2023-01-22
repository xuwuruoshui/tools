package handler

import (
	"end/domain"
	"end/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UserUsecase usecase.UserUseCase
	contextTimeout time.Duration
}


func (h *UserHandler) GetUser(c *gin.Context) {

	id := c.Param("id")
	u ,err := h.UserUsecase.GetById(c,id)
	if err!=nil{
		c.JSON(http.StatusNotFound, u)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) GetUserList(c *gin.Context) {

	pageSizeStr := c.Query("pageSize")
	pageNoStr := c.Query("pageNo")

	pageSize, _ := strconv.Atoi(pageSizeStr)
	pageNo, _ := strconv.Atoi(pageNoStr)

	var user domain.User
	err := c.ShouldBindQuery(&user)
	if err != nil {
		panic(err)
	}

	u ,_ := h.UserUsecase.GetList(c,domain.PageDomain[domain.User]{
		PageSize: pageSize,
		PageNo: pageNo,
		Condition: domain.User{},
	})
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	u ,err := h.UserUsecase.Create(c,user)
	if err!=nil{
		c.JSON(http.StatusNotFound, u)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {

	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	res ,err := h.UserUsecase.Update(c,user)
	if err!=nil{
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {

	ids := c.Param("ids")
	res ,err := h.UserUsecase.Delete(c,ids)
	if err!=nil{
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}