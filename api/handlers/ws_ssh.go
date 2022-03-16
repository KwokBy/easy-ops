package handlers

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/KwokBy/easy-ops/pkg/ssh"
	"github.com/KwokBy/easy-ops/pkg/validate"
	"github.com/gorilla/websocket"

	"github.com/KwokBy/easy-ops/models"
	"github.com/gin-gonic/gin"
)

type WsSshHandler struct {
}

func NewWsSshHandler() WsSshHandler {
	return WsSshHandler{}
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (w *WsSshHandler) WSSSH(c *gin.Context) {
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if validate.IsWSError(wsConn, err) {
		return
	}
	defer wsConn.Close()

	// TODO: 主机信息从数据库获取
	machine := models.Host{
		HostName: "106.55.161.12",
		Host:     "106.55.161.12",
		Port:     22,
		Name:     "root",
		Password: "Gl@987963951",
		SSHType:  "password",
	}
	// 从url中获取terminal大小
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))
	if validate.IsWSError(wsConn, err) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if validate.IsWSError(wsConn, err) {
		return
	}

	// 创建ssh客户端
	client, err := ssh.NewSSHClient(machine)
	if validate.IsWSError(wsConn, err) {
		return
	}
	defer client.Close()

	//  创建ssh session
	conn, err := ssh.NewSSHConn(cols, rows, client)
	if validate.IsWSError(wsConn, err) {
		return
	}
	defer conn.Close()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)
	// most messages are ssh output, not webSocket input
	go conn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go conn.SendComboOutput(wsConn, quitChan)
	go conn.SessionWait(quitChan)

	<-quitChan
}
