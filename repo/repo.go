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

// UserRepo represent the user repository contract
type UserRepo interface {
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
	// GetUserByName return user by name
	GetUserByName(ctx context.Context, name string) (models.User, error)
	DeleteUserByName(ctx context.Context, username string) error
}

// HostRepo represent the host repository contract
type HostRepo interface {
	// GetHostsByUsername return hosts by owner
	GetHostsByUsername(ctx context.Context, owner string) ([]models.Host, error)
	// AddHost add host
	AddHost(ctx context.Context, host models.Host) error
	// UpdateHost update host
	UpdateHost(ctx context.Context, host models.Host) error
	// DeleteHost delete host
	DeleteHost(ctx context.Context, id int64) error
}

type TaskRepo interface {
	// GetTasksByUsername return tasks by username
	GetTasksByUsername(ctx context.Context, username string) ([]models.Task, error)
	// AddTask add task
	AddTask(ctx context.Context, task models.Task) error
	// GetTaskAndHosts  return task and hosts
	GetTaskAndHosts(ctx context.Context, taskId int64, hostIds []int64) (models.Task, []models.Host, error)
	// UpdateTaskStatus update task status
	UpdateTaskStatus(ctx context.Context, taskId int64, status int) error
	// UpdateTaskEntryId update task entry id
	UpdateTaskEntryId(ctx context.Context, taskId int64, entryIds string) error
	// GetTaskByID return task by id
	GetTaskByID(ctx context.Context, taskId int64) (models.Task, error)
	// UpdateTask update task
	UpdateTask(ctx context.Context, task models.Task) error
}

type ExecHistoryInfoRepo interface {
	// AddExecHistory 添加执行历史
	AddExecHistory(ctx context.Context, execHistory models.ExecHistoryInfo) error
	// BatchAddExecHistory 批量添加执行历史
	BatchAddExecHistory(ctx context.Context, execHistories []models.ExecHistoryInfo) error
	// GetExecHistoryByTaskID 根据TaskID获取执行历史
	GetExecHistoryByTaskID(ctx context.Context, taskID int64) (
		[]models.ExecHistoryInfo, error)
	// GetCountGroupByExecID 获取某个Task下的执行次数
	GetCountGroupByExecID(ctx context.Context, taskID int64) (int, error)
	// GetExecHistoryByTaskIDAndExecID
	GetExecHistoryByTaskIDAndExecID(ctx context.Context, taskID int64, execID int64) (
		[]models.ExecHistoryInfo, error)
}

type ExecHistoryRepo interface {
	// AddExecHistory 添加执行历史
	AddExecHistory(ctx context.Context, execHistory models.ExecHistory) error
	// GetExecHistoryByTaskID 根据TaskID获取执行历史
	GetExecHistoryByTaskID(ctx context.Context, taskID int64) ([]models.ExecHistory, error)
	// GetExecHistoryCountByTaskID 获取某个Task下的执行次数
	GetExecHistoryCountByTaskID(ctx context.Context, taskID int64) (int64, error)
}

type ImageRepo interface {
	// GetImageByOwner 获取某个用户的镜像
	GetImageByOwner(ctx context.Context, username string) (
		[]models.Image, error)
	// AddImage 添加镜像
	AddImage(ctx context.Context, image models.Image) error
	// GetImageByImageID 获取某些镜像
	GetImageByImageID(ctx context.Context, imageID int64) (
		[]models.Image, error)
	// GetImageByID 获取某个镜像
	GetImageByID(ctx context.Context, id int) (
		models.Image, error)
	// DeleteImage 删除某个镜像
	DeleteImage(ctx context.Context, name, version string) error
}

type RoleRepo interface {
	// GetRoleByID 获取某个角色
	GetRoleByID(ctx context.Context, id int64) (models.Role, error)
	// AddRole 添加角色
	AddRole(ctx context.Context, role models.Role) error
	// GetRoles 获取所有角色
	GetRoles(ctx context.Context) ([]models.Role, error)
	// DeleteRole 删除角色
	DeleteRole(ctx context.Context, id int64) error
	// UpdateRole 更新角色
	UpdateRole(ctx context.Context, role models.Role) error
}

type ApiRepo interface {
	// GetApis 获取所有接口
	GetApis(ctx context.Context) ([]models.Api, error)
	// GetApisByID 获取某个接口
	GetApisByID(ctx context.Context, ids []int) ([]models.Api, error)
}

type MenuRepo interface {
}

type AuthMenuRepo interface {
}

type CasbinRepo interface {
	// GetByRoleID 获取某个角色的所有权限
	GetByRoleID(ctx context.Context, roleID int64) ([]models.Casbin, error)
}
