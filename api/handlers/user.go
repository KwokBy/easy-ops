package handlers

import (
	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("login error: %s", err.Error())
		response.FailWithData(err, "login error", c)
		return
	}
	user, err := u.userService.Login(c, req.Username, req.Password)
	if err != nil {
		zlog.Errorf("login error: %s", err.Error())
		response.FailWithData(err, "login error", c)
		return
	}
	token, err := u.userService.GenerateToken(c, models.Token{
		Username: user.Username,
	})
	if err != nil {
		zlog.Errorf("generate token error: %s", err.Error())
		response.FailWithData(err, "generate token error", c)
		return
	}
	response.OKWithData(token, "login success", c)
}

func (u *UserHandler) RefreshToken(c *gin.Context) {
	var oldToken models.Token
	if err := c.ShouldBind(&oldToken); err != nil {
		zlog.Errorf("refresh token error: %s", err.Error())
		response.FailWithData(err, "refresh token error", c)
		return
	}
	token, err := u.userService.GenerateToken(c, oldToken)
	if err != nil {
		zlog.Errorf("generate token error: %s", err.Error())
		response.FailWithData(err, "generate token error", c)
		return
	}
	response.OKWithData(token, "refresh token success", c)
}

type PasswordResetReq struct {
	Username string `json:"username"`
}

// PasswordReset 密码重置
func (u *UserHandler) PasswordReset(c *gin.Context) {
	var req PasswordResetReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("password reset error: %s", err.Error())
		response.FailWithData(err, "password reset error", c)
		return
	}
	if err := u.userService.PasswordReset(c, req.Username); err != nil {
		zlog.Errorf("password reset error: %s", err.Error())
		response.FailWithData(err, "password reset error", c)
		return
	}
	response.OK("password reset success", c)
}

type RoleSetReq struct {
	Username string  `json:"username"`
	Roles    []int64 `json:"roles"`
}

// RoleSet 角色设置
func (u *UserHandler) RoleSet(c *gin.Context) {
	var req RoleSetReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("role set error: %s", err.Error())
		response.FailWithData(err, "role set error", c)
		return
	}
	if err := u.userService.RoleSet(c, req.Username, req.Roles); err != nil {
		zlog.Errorf("role set error: %s", err.Error())
		response.FailWithData(err, "role set error", c)
		return
	}
	response.OK("role set success", c)
}

type DeleteUserReq struct {
	Username string `json:"username"`
}

// Delete 删除用户
func (u *UserHandler) Delete(c *gin.Context) {
	var req DeleteUserReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("delete user error: %s", err.Error())
		response.FailWithData(err, "delete user error", c)
		return
	}
	if err := u.userService.DeleteUser(c, req.Username); err != nil {
		zlog.Errorf("delete user error: %s", err.Error())
		response.FailWithData(err, "delete user error", c)
		return
	}
	response.OK("delete user success", c)
}
