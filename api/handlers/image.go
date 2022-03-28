package handlers

import (
	"github.com/KwokBy/easy-ops/pkg/docker"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/validate"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type ImageHandler struct{}

func NewImageHandler() ImageHandler {
	return ImageHandler{}
}

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
	hr, err := docker.ExecContainer(cli, "test", []string{"/bin/bash"})
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
