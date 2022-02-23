package handlers

import (
	"net/http"

	"github.com/KwokBy/easy-ops/pkg/zlog"

	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoService service.IDemoService
}

func NewDemoHandler(demoService service.IDemoService) DemoHandler {
	return DemoHandler{
		demoService: demoService,
	}
}

func (d *DemoHandler) GetLongDemo(c *gin.Context) {
	demo, err := d.demoService.GetLongDemo(c)
	if err != nil {
		zlog.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	zlog.Infof("demo: %+v", demo)
	zlog.Debugf("demo: %+v", demo)
	zlog.Warnf("demo: %+v", demo)
	c.JSON(http.StatusOK, demo)
}
