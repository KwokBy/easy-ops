package service

import (
	"context"
	"fmt"
	"time"

	"github.com/KwokBy/easy-ops/configs"
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

// GenerateToken 生成token
func (u *userService) GenerateToken(ctx context.Context, oldToken models.Token) (models.Token, error) {
	user, err := u.userRepo.GetUserByName(ctx, oldToken.Username)
	zlog.Infof("[GenerateToken] GetUserByName: %+v", user)
	if err != nil {
		zlog.Errorf("[GenerateToken] GetUserByName error: %s", err.Error())
		return oldToken, fmt.Errorf("用户名不存在")
	}
	if oldToken.Token != "" {
		if user.AccessToken != oldToken.Token {
			zlog.Errorf("[GenerateToken] GetUserByName error: token not match")
			return oldToken, fmt.Errorf("token不匹配")
		}
	}
	config := configs.New()
	token, expireTime, err := jwt.New(jwt.Data{UserID: user.ID, RoleID: user.RoleID}, config.JWT.Secret)
	if err != nil {
		zlog.Errorf("[GenerateToken] 获取token失败: %s", err.Error())
		return models.Token{}, fmt.Errorf("获取token失败: %s", err.Error())
	}

	// 更新数据库
	user.AccessToken = token
	user.UpdatedTime = time.Now()
	u.userRepo.UpdateUser(ctx, user)

	return models.Token{
		Token:      token,
		Username:   user.Username,
		ExpireTime: expireTime,
	}, nil
}

// PasswordReset 密码重置
func (u *userService) PasswordReset(ctx context.Context, username string) error {
	user, err := u.userRepo.GetUserByName(ctx, username)
	if err != nil {
		zlog.Errorf("[PasswordReset] GetUserByName error: %s", err.Error())
		return fmt.Errorf("用户名不存在")
	}
	user.PasswordHash = "123456"
	if err := u.userRepo.UpdateUser(ctx, user); err != nil {
		zlog.Errorf("[PasswordReset] UpdateUser error: %s", err.Error())
		return fmt.Errorf("重置密码失败")
	}
	return nil
}

// RoleSet 角色设置
func (u *userService) RoleSet(ctx context.Context, username string, roles []int64) error {
	user, err := u.userRepo.GetUserByName(ctx, username)
	if err != nil {
		zlog.Errorf("[RoleSet] GetUserByName error: %s", err.Error())
		return fmt.Errorf("用户名不存在")
	}
	user.RoleID = roles[0]
	if err := u.userRepo.UpdateUser(ctx, user); err != nil {
		zlog.Errorf("[RoleSet] UpdateUser error: %s", err.Error())
		return fmt.Errorf("设置角色失败")
	}
	return nil
}

// DeleteUser 删除用户
func (u *userService) DeleteUser(ctx context.Context, username string) error {
	if err := u.userRepo.DeleteUserByName(ctx, username); err != nil {
		zlog.Errorf("[DeleteUser] DeleteUserByName error: %s", err.Error())
		return fmt.Errorf("删除用户失败")
	}
	return nil
}
