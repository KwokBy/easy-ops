package service

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"github.com/docker/docker/api/types"
)

type IDemoService interface {
	GetLongDemo(ctx context.Context) (string, error)
}

type UserService interface {
	// Login 登录
	Login(ctx context.Context, username, password string) (models.User, error)
	// Register 注册
	Register(ctx context.Context, username, password string) (models.User, error)
	// GenerateToken 生成token
	GenerateToken(ctx context.Context, oldToken models.Token) (models.Token, error)
	// PasswordReset 密码重置
	PasswordReset(ctx context.Context, username string) error
	// RoleSet 角色设置
	RoleSet(ctx context.Context, username string, roles []string) error
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, username string) error
}

type HostService interface {
	// GetHostsByUsername 根据用户名获取主机列表
	GetHostsByUsername(ctx context.Context, owner string) ([]models.Host, error)
	// AddHost 添加主机
	AddHost(ctx context.Context, host models.Host) error
	// DeleteHost 删除主机
	DeleteHost(ctx context.Context, id int64) error
	// UpdateHost 更新主机信息
	UpdateHost(ctx context.Context, host models.Host) error
	// VerifyHost 校验主机信息
	VerifyHost(ctx context.Context, host models.Host) error
}

type TaskService interface {
	// GetTasksByUsername 根据用户名获取任务列表
	GetTasksByUsername(ctx context.Context, userName string) ([]models.TaskDTO, error)
	// AddTask 添加任务
	AddTask(ctx context.Context, task models.TaskDTO) error
	// AddTaskAndRun 添加任务并执行
	AddTaskAndRun(ctx context.Context, task models.TaskDTO) error
	// ExecuteTask 执行任务
	ExecuteTask(ctx context.Context, task models.TaskDTO) error
	// ExecuteTest 执行测试任务
	ExecuteTest(ctx context.Context, task models.TaskDTO) error
	// StopTask 停止任务
	StopTask(ctx context.Context, id int64) error
	// DeleteTask
	DeleteTask(ctx context.Context, id int64) error
}

type ExecHistoryInfoService interface {
	// GetExecHistoryInfos
	GetExecHistoryInfos(ctx context.Context, taskID, execID int64) (models.ExecHistoryDTO, error)
}

type ExecHistoryService interface {
	// GetExecHistoriesByTaskID 获取任务执行历史
	GetExecHistoriesByTaskID(ctx context.Context, taskID int64) ([]models.ExecHistory, error)
}

type ImageService interface {
	// DebugImage 测试镜像
	DebugImage(ctx context.Context, id int) (types.HijackedResponse, error)
	// SaveImage 保存镜像
	SaveImage(ctx context.Context, image models.Image) error
	// GetImageByOwner 根据用户名获取镜像列表
	GetImages(ctx context.Context, username string) ([]models.Image, error)
	// DeleteImage 删除镜像
	DeleteImage(ctx context.Context, name, version string) error
}

type RoleService interface {
	// GetRoleByID 根据ID获取角色
	GetRoleByID(ctx context.Context, id int64) (models.Role, error)
	// AddRole 添加角色
	AddRole(ctx context.Context, role models.Role) error
	// DeleteRole 删除角色
	DeleteRole(ctx context.Context, id int64) error
	// UpdateRole 更新角色
	UpdateRole(ctx context.Context, role models.Role) error
	// GetRoles 获取角色列表
	GetRoles(ctx context.Context) ([]models.Role, error)
	// 获取角色API权限
	GetRoleAPIs(ctx context.Context, id int64) ([]models.Casbin, error)
	// 设置角色API权限
	SetRoleAPIs(ctx context.Context, id int64, apis []models.Casbin) error
	// 获取角色资源权限
	GetRoleResources(ctx context.Context, id int64) ([]models.AuthMenu, error)
	// 设置角色资源权限
	SetRoleResources(ctx context.Context, id int64, resources []models.AuthMenu) error
}
