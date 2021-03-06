package handlers

import (
	"fmt"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type HostHandler struct {
	hostService service.HostService
}

func NewHostHandler(service service.HostService) HostHandler {
	return HostHandler{hostService: service}
}

// GetGetHostsReq 获取主机列表请求参数
type GetGetHostsReq struct {
	Username string `json:"owner"`
}

// GetHosts 获取用户有权限的主机列表
func (h *HostHandler) GetHosts(c *gin.Context) {
	var req GetGetHostsReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("get owner form uri error: %s", err.Error())
		response.FailWithData(err, "get owner form uri error", c)
		return
	}
	hosts, err := h.hostService.GetHostsByUsername(c, req.Username)
	if err != nil {
		zlog.Errorf("get hosts by owner error: %s", err.Error())
		response.FailWithData(err, "get hosts by owner error", c)
		return
	}
	response.OKWithData(hosts, "get hosts by owner success", c)
}

// AddHost 添加主机
func (h *HostHandler) AddHost(c *gin.Context) {
	var host models.Host
	if err := c.ShouldBind(&host); err != nil {
		zlog.Errorf("add host error: %s", err.Error())
		response.FailWithData(err, "add host error", c)
		return
	}
	if err := h.hostService.AddHost(c, host); err != nil {
		zlog.Errorf("add host error: %s", err.Error())
		response.FailWithData(err, "add host error", c)
		return
	}
	host.SSHType = "ssh-password"
	host.Host = fmt.Sprint(host.Host, ":", host.Port)
	h.hostService.VerifyHost(c, host)
	response.OKWithData(nil, "add host success", c)
}

// DeleteHostReq 删除主机请求参数
type DeleteHostReq struct {
	HostID int64 `json:"host_id"`
}

// DeleteHost 删除主机
func (h *HostHandler) DeleteHost(c *gin.Context) {
	var req DeleteHostReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("delete host error: %s", err.Error())
		response.FailWithData(err, "delete host error", c)
		return
	}
	if err := h.hostService.DeleteHost(c, req.HostID); err != nil {
		zlog.Errorf("delete host error: %s", err.Error())
		response.FailWithData(err, "delete host error", c)
		return
	}
	response.OKWithData(nil, "delete host success", c)
}

// UpdateHost 更新主机信息
func (h *HostHandler) UpdateHost(c *gin.Context) {
	var host models.Host
	if err := c.ShouldBind(&host); err != nil {
		zlog.Errorf("update host error: %s", err.Error())
		response.FailWithData(err, "update host error", c)
		return
	}
	if err := h.hostService.UpdateHost(c, host); err != nil {
		zlog.Errorf("update host error: %s", err.Error())
		response.FailWithData(err, "update host error", c)
		return
	}
	response.OKWithData(nil, "update host success", c)
}

// VerifyHost 验证主机
func (h *HostHandler) VerifyHost(c *gin.Context) {
	var host models.Host
	if err := c.ShouldBind(&host); err != nil {
		zlog.Errorf("verify host error: %s", err.Error())
		response.FailWithData(err, "verify host error", c)
		return
	}
	if err := h.hostService.VerifyHost(c, host); err != nil {
		zlog.Errorf("verify host error: %s", err.Error())
		response.FailWithData(err, "verify host error", c)
		return
	}
	response.OKWithData(nil, "verify host success", c)
}
