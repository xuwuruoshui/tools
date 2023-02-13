package handler

import (
	"end/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// gin Router wrapper 统一返回

type RouterWrapper struct {
	Router *gin.Engine
}

func NewRouterPlus(engine *gin.Engine) *RouterWrapper {
	return &RouterWrapper{
		Router: engine,
	}
}

type RouterHandler func(ctx *gin.Context) any

// func (r *RouterWrapper) GET(path string, handler.txt RouterHandler,handlers ...gin.HandlerFunc) {
// 	preHandler := PreHandler(handler.txt)
// 	postHanlder := PermAuth(handler.txt)
// 	handlers = append(handlers,preHandler, postHanlder)
// 	r.Router.GET(path, handlers...)
// }

func (r *RouterWrapper) GETJWT(path string, handler RouterHandler, perms ...string) {
	r.Router.GET(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterWrapper) GET(path string, handler RouterHandler, perms ...string) {
	r.Router.GET(path, DefaultResp(handler))
}

func (r *RouterWrapper) POSTJWT(path string, handler RouterHandler, perms ...string) {
	r.Router.POST(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterWrapper) POST(path string, handler RouterHandler, perms ...string) {
	r.Router.POST(path, DefaultResp(handler))
}

func (r *RouterWrapper) PUTJWT(path string, handler RouterHandler, perms ...string) {
	r.Router.PUT(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterWrapper) PUT(path string, handler RouterHandler, perms ...string) {
	r.Router.PUT(path, DefaultResp(handler))
}

func (r *RouterWrapper) DELETEJWT(path string, handler RouterHandler, perms ...string) {
	r.Router.DELETE(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterWrapper) DELETE(path string, handler RouterHandler, perms ...string) {
	r.Router.DELETE(path, DefaultResp(handler))
}

func (r *RouterWrapper) Group(path string, handlers ...gin.HandlerFunc) *RouterGroupWrapper {
	return &RouterGroupWrapper{
		routerGroup: r.Router.Group(path, handlers...),
	}
}

type RouterGroupWrapper struct {
	routerGroup *gin.RouterGroup
}

func (r *RouterGroupWrapper) GETJWT(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.GET(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterGroupWrapper) GET(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.GET(path, DefaultResp(handler))
}

func (r *RouterGroupWrapper) POSTJWT(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.POST(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterGroupWrapper) POST(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.POST(path, DefaultResp(handler))
}

func (r *RouterGroupWrapper) PUTJWT(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.PUT(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterGroupWrapper) PUT(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.PUT(path, DefaultResp(handler))
}

func (r *RouterGroupWrapper) DELETEJWT(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.DELETE(path, JwtAuth(), PermAuth(perms...), DefaultResp(handler))
}
func (r *RouterGroupWrapper) DELETE(path string, handler RouterHandler, perms ...string) {
	r.routerGroup.DELETE(path, DefaultResp(handler))
}

func (r *RouterGroupWrapper) Group(path string, handlers ...gin.HandlerFunc) *RouterGroupWrapper {
	return &RouterGroupWrapper{
		routerGroup: r.routerGroup.Group(path, handlers...),
	}
}

func DefaultResp(handler RouterHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m := handler(ctx).(*ApiRespData)
		// 不是websocket直接返回json
		if ctx.GetHeader("Upgrade")!="websocket"{
			ctx.JSON(http.StatusOK, m)
		}
		ctx.Next()
	}
}

// 返回处理
func PermAuth(perms ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(perms) != 0 {

			ctx.JSON(http.StatusForbidden, ApiResp(NOPERM, nil))
			ctx.Abort()
			return
		}
		ctx.Next()
		return
	}
}

func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, ApiResp(NOLOGIN, nil))
			c.Abort()
			return
		}

		tokens := strings.Split(token, " ")
		if tokens[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, ApiResp(NOLOGIN, nil))
		}

		j := utils.NewJWT()
		parseToken, err := j.ParseToken(tokens[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, ApiResp(NOLOGIN, nil))
			c.Abort()
			return
		}
		c.Set(utils.Claims, parseToken)
		c.Next()
		return
	}
}
