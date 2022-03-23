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

type MirrorService interface {
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
	// GetExecHistoriesByTaskID 获取任务执行历史
	GetExecHistoriesByTaskID(ctx context.Context, taskID int64) ([][]models.ExecHistoryInfo, error)
}
