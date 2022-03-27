package handlers

import (
	"net/http"

	"github.com/KwokBy/easy-ops/pkg/docker"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/validate"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ImageHandler struct{}

func NewImageHandler() ImageHandler {
	return ImageHandler{}
}

func (h ImageHandler) Debug(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		zlog.Errorf("get docker client error: %s", err.Error())
		response.FailWithData(err, "get docker client error", c)
		return
	}
	hr, err := docker.ExecContainer(cli, "test", []string{"/bin/bash"})
	if err != nil {
		zlog.Errorf("exec container error: %s", err.Error())
		response.FailWithData(err, "exec container error", c)
		return
	}
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024 * 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if validate.IsWSError(wsConn, err) {
		return
	}
	defer wsConn.Close()
	// 关闭I/O流
	defer hr.Close()
	// 退出进程
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()
	hr.Conn.Write([]byte("cd /home\n ls\r"))
	docker.WsReaderCopy(wsConn, hr.Conn)
	go docker.WsWriterCopy(wsConn, hr.Conn)
}
