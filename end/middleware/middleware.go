package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossDomain(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", c.GetHeader("origin"))
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,istoken, x-token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
	c.Header("Access-Control-All-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Expose-Headers", "*")

	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
