// Package repo 数据层操作
package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
)

// IDemoRepo represent the demo repository contract
type IDemoRepo interface {
	// GetDemos return all demos
	GetDemos(ctx context.Context) ([]models.Demo, error)
}

// IUserRepo represent the user repository contract
type IUserRepo interface {
	// GetUsersByNameAndPwd return user by name and pwd
	GetUsersByNameAndPwd(ctx context.Context, name, pwd string) (models.User, error)
	// UpdateUser update user
	UpdateUser(ctx context.Context, user models.User) error
	// AddUser add user
	AddUser(ctx context.Context, user models.User) error
	// DeleteUser delete user
	DeleteUser(ctx context.Context, id int64) error
	// GetUsers return all users
	GetUsers(ctx context.Context) ([]models.User, error)
}

// IHostRepo represent the host repository contract
type IHostRepo interface {
	// GetHostsByUsername return hosts by username
	GetHostsByUsername(ctx context.Context, username string) ([]models.Host, error)
	// AddHost add host
	AddHost(ctx context.Context, host models.Host) error
	// UpdateHost update host
	UpdateHost(ctx context.Context, host models.Host) error
	// DeleteHost delete host
	DeleteHost(ctx context.Context, id int64) error
}

// IMirrorRepo represent the mirror repository contract
type IMirrorRepo interface {
	
}
