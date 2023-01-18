package handler

import (
	"context"
	"end/domain"
	"end/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	App         *internal.App
	UserService domain.UserService
}

func NewUserHandler(app *internal.App) *UserHandler {
	return &UserHandler{
		App:         app,
		UserService: &domain.UserImpl{},
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, h.App.Config.Timeout)
	defer cancel()

	// req vo
	u := h.UserService.GetUser(ctx)
	// resp vo
	c.JSON(http.StatusOK, u)
}
