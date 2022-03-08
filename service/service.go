package service

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
)

type IDemoService interface {
	GetLongDemo(ctx context.Context) (string, error)
}

type UserService interface {
	// Login 登录
	Login(ctx context.Context, username, password string) (models.User, error)
	// Register 注册
	Register(ctx context.Context, username, password string) (models.User, error)
}

type HostService interface {
}

type MirrorService interface {
}

type TaskService interface {
}
