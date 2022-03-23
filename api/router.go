package api

import (
	"net/http"

	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/KwokBy/easy-ops/pkg/jwt"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Demo        handlers.DemoHandler
	User        handlers.UserHandler
	WsSsh       handlers.WsSshHandler
	Host        handlers.HostHandler
	Task        handlers.TaskHandler
	ExecHistoryInfo handlers.ExecHistoryInfoHandler
}

func (r *Router) With(engine *gin.Engine) {
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

	demo := engine.Group("/api/v1/demo")
	{
		demo.GET("/", r.Demo.GetLongDemo)
		demo.GET("/ws", r.Demo.Wshandler)
	}
	user := engine.Group("/api/v1/user")
	{
		user.POST("/login", r.User.Login)
	}
	host := engine.Group("/api/v1/host")
	{
		host.POST("/get", r.Host.GetHosts)
		host.POST("/add", r.Host.AddHost)
		host.POST("/delete", r.Host.DeleteHost)
		host.POST("/update", r.Host.UpdateHost)
		host.POST("/verify", r.Host.VerifyHost)
	}
	wsSsh := engine.Group("/api/v1/ws")
	{
		wsSsh.GET("/ssh", r.WsSsh.WSSSH)
	}
	task := engine.Group("/api/v1/task")
	{
		task.POST("/get", r.Task.GetTasks)
		task.POST("/add", r.Task.AddTask)
		task.POST("/exec", r.Task.ExecuteTask)
		task.POST("/stop", r.Task.StopTask)
		task.POST("/addAndRun", r.Task.AddTaskAndExecute)
		task.POST("/test", r.Task.ExecuteTest)
	}
	execHistory := engine.Group("/api/v1/execHistory")
	{
		execHistory.POST("/get", r.ExecHistoryInfo.GetExecHistory)
	}

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
