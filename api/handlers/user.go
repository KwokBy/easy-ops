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
