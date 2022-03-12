package api

import (
	"net/http"

	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/KwokBy/easy-ops/pkg/jwt"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Demo handlers.DemoHandler
	User handlers.UserHandler
}

func (r *Router) With(engine *gin.Engine) {
	demo := engine.Group("/api/v1/demo")
	{
		demo.GET("/", r.Demo.GetLongDemo)
		demo.GET("/ws", r.Demo.Wshandler)
	}
	user := engine.Group("/api/v1/user")
	{
		user.POST("/login", r.User.Login)
	}
	engine.LoadHTMLFiles("index.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	engine.GET("/getAsyncRoutes", func(c *gin.Context) {
		response.OKWithData([]PermissionRouter{
			{
				Path:     "/permission",
				Name:     "permission",
				Redirect: "/permission/page/index",
				Meta: Meta{
					Title: "menus.permission",
					Icon:  "lollipop",
					I18n:  true,
					Rank:  3,
				},
				Children: []Child{
					{
						Path: "/permission/page/index",
						Name: "permissionPage",
						Meta: Meta{
							Title: "menus.permissionPage",
							I18n:  true,
							Rank:  3,
						},
					},
					{
						Path: "/permission/button/index",
						Name: "permissionButton",
						Meta: Meta{
							Title:     "menus.permissionButton",
							I18n:      true,
							Authority: []string{"v-admin"},
							Rank:      3,
						},
					},
				},
			},
		},
			"获取成功", c)
	})
}

type PermissionRouter struct {
	Path     string  `json:"path"`
	Name     string  `json:"name"`
	Redirect string  `json:"redirect"`
	Meta     Meta    `json:"meta"`
	Children []Child `json:"children"`
}
type Meta struct {
	Title     string   `json:"title"`
	Icon      string   `json:"icon"`
	I18n      bool     `json:"i18n"`
	Rank      int      `json:"rank"`
	Authority []string `json:"authority"`
}
type Child struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Meta Meta   `json:"meta"`
}

//   // 添加不同按钮权限到/permission/button页面中
//   function setDifAuthority(authority, routes) {
// 	routes.children[1].meta.authority = [authority];
// 	return routes;
//   }
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
