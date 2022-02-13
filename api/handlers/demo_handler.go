package handlers

import (
	"fmt"
	"net/http"

	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoService service.IDemoService
}

func NewDemoHandler(demoService service.IDemoService) *DemoHandler {
	return &DemoHandler{
		demoService: demoService,
	}
}

func (d *DemoHandler) GetLongDemo(c *gin.Context) {
	demo, err := d.demoService.GetLongDemo(c)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	c.JSON(http.StatusOK, demo)
}
