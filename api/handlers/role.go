package handlers

import (
	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(service service.RoleService) RoleHandler {
	return RoleHandler{
		roleService: service,
	}
}

// GetRoles 获取角色列表
func (h *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.roleService.GetRoles(c)
	if err != nil {
		response.FailWithData(err, "get roles error", c)
		return
	}
	response.OKWithData(roles, "get roles success", c)
}

// AddRole 添加角色
func (h *RoleHandler) AddRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBind(&role); err != nil {
		zlog.Errorf("add role error: %s", err.Error())
		response.FailWithData(err, "add role error", c)
		return
	}
	if err := h.roleService.AddRole(c, role); err != nil {
		zlog.Errorf("add role error: %s", err.Error())
		response.FailWithData(err, "add role error", c)
		return
	}
	response.OKWithData(nil, "add role success", c)
}

type DeleteRoleReq struct {
	ID int64 `json:"id"`
}

// DeleteRole 删除角色
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	var req DeleteRoleReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("delete role error: %s", err.Error())
		response.FailWithData(err, "delete role error", c)
		return
	}
	if err := h.roleService.DeleteRole(c, req.ID); err != nil {
		zlog.Errorf("delete role error: %s", err.Error())
		response.FailWithData(err, "delete role error", c)
		return
	}
	response.OKWithData(nil, "delete role success", c)
}

// // SetRolePermissions 设置角色权限
// func (h *RoleHandler) SetRolePermissions(c *gin.Context) {
// 	var req models.RolePermissionsReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		zlog.Errorf("set role permissions error: %s", err.Error())
// 		response.FailWithData(err, "set role permissions error", c)
// 		return
// 	}
// 	if err := h.roleService.SetRolePermissions(c, req); err != nil {
// 		zlog.Errorf("set role permissions error: %s", err.Error())
// 		response.FailWithData(err, "set role permissions error", c)
// 		return
// 	}
// 	response.OKWithData(nil, "set role permissions success", c)
// }
