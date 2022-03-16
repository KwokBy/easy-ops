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
	// GetHostsByUsername 根据用户名获取主机列表
	GetHostsByUsername(ctx context.Context, username string) ([]models.Host, error)
	// AddHost 添加主机
	AddHost(ctx context.Context, host models.Host) error
	// DeleteHost 删除主机
	DeleteHost(ctx context.Context, id int64) error
	// UpdateHost 更新主机信息
	UpdateHost(ctx context.Context, host models.Host) error
}

type MirrorService interface {
}

type TaskService interface {
}
