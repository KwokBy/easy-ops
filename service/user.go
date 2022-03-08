package service

import (
	"context"
	"fmt"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/jwt"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/repo"
)

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Login 登录
func (u *userService) Login(ctx context.Context, username, password string) (models.User, error) {
	user, err := u.userRepo.GetUsersByNameAndPwd(ctx, username, password)
	if err != nil {
		zlog.Errorf("[Login] GetUsersByNameAndPwd error: %s", err.Error())
		return user, fmt.Errorf("用户名或密码错误")
	}
	if user.Username == "" {
		zlog.Errorf("[Login] GetUsersByNameAndPwd error: user not found")
		return user, fmt.Errorf("用户名或密码错误/用户不存在")
	}
	token, err := jwt.New(jwt.Data{UserID: 1})
	if err != nil {
		zlog.Errorf("[Login] 获取token失败: %s", err.Error())
		return user, fmt.Errorf("获取token失败: %s", err.Error())
	}
	user.AccessToken = token
	return user, nil
}

// Register 注册
func (u *userService) Register(ctx context.Context, username, password string) (models.User, error) {
	user, err := u.userRepo.GetUserByName(ctx, username)
	if err != nil {
		zlog.Errorf("[Register] GetUsersByName error: %s", err.Error())
		return user, fmt.Errorf("用户名已存在")
	}
	if user.Username != "" {
		zlog.Errorf("[Register] GetUsersByName error: user not found")
		return user, fmt.Errorf("用户名已存在")
	}
	user.Username = username
	user.PasswordHash = password
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	err = u.userRepo.AddUser(ctx, user)
	if err != nil {
		zlog.Errorf("[Register] CreateUser error: %s", err.Error())
		return user, fmt.Errorf("注册失败")
	}
	return user, nil
}
