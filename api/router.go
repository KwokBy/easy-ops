package api

import (
	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	demo handlers.DemoHandler
}

func (r *Router) With(engine *gin.Engine) {
	engine.GET("/demo", r.demo.GetLongDemo)

}

func NewRouter(demo handlers.DemoHandler) *Router {
	router := &Router{
		demo: demo,
	}
	return router
}