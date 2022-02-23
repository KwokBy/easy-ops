package api

import (
	"net/http"

	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/KwokBy/easy-ops/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Demo handlers.DemoHandler
	User handlers.UserHandler
}

func (r *Router) With(engine *gin.Engine) {
	demo := engine.Group("/api/v1/demo", JWTAuth())
	{
		demo.GET("/", r.Demo.GetLongDemo)
	}
	user := engine.Group("/api/v1/user")
	{
		user.POST("/login", r.User.Login)
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		ok, err := jwt.IsValid(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token 无效",
			})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token 已过期",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
