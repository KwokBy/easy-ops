package handlers

import (
	"strconv"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/docker"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/validate"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService service.ImageService
}

func NewImageHandler(service service.ImageService) ImageHandler {
	return ImageHandler{
		imageService: service,
	}
}

// Debug 调试镜像
func (h ImageHandler) Debug(c *gin.Context) {

	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if validate.IsWSError(wsConn, err) {
		return
	}
	defer wsConn.Close()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		zlog.Errorf("get docker client error: %s", err.Error())
		response.FailWithData(err, "get docker client error", c)
		return
	}
	cli.NegotiateAPIVersion(c)
	defer cli.Close()
	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
	if validate.IsWSError(wsConn, err) {
		return
	}
	hr, err := h.imageService.DebugImage(c, id)
	if err != nil {
		zlog.Errorf("exec container error: %s", err.Error())
		response.FailWithData(err, "exec container error", c)
		return
	}
	// 退出进程
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()
	// 关闭I/O流
	defer hr.Close()
	go docker.WsWriterCopy(wsConn, hr.Conn)

	docker.WsReaderCopy(wsConn, hr.Conn)
}

// Save 保存镜像
func (h ImageHandler) Save(c *gin.Context) {
	var image models.Image
	if err := c.ShouldBind(&image); err != nil {
		zlog.Errorf("debug image error: %s", err.Error())
		response.FailWithData(err, "debug image error", c)
		return
	}
	if err := h.imageService.SaveImage(c, image); err != nil {
		zlog.Errorf("save image error: %s", err.Error())
		response.FailWithData(err, "save image error", c)
		return
	}
	response.OK("save image success", c)
}

type GetImagesReq struct {
	Owner string `json:"owner"`
}

// List 列出镜像
func (h ImageHandler) GetImages(c *gin.Context) {
	var req GetImagesReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("list image error: %s", err.Error())
		response.FailWithData(err, "list image error", c)
		return
	}
	images, err := h.imageService.GetImages(c, req.Owner)
	if err != nil {
		zlog.Errorf("list image error: %s", err.Error())
		response.FailWithData(err, "list image error", c)
		return
	}
	response.OKWithData(images, "list image success", c)
}
