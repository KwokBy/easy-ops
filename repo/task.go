package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlTaskRepo struct {
	DB *gorm.DB
}

func NewMysqlTaskRepo(DB *gorm.DB) TaskRepo {
	return &mysqlTaskRepo{DB}
}

// GetTasksByUsername 根据用户名获取任务列表
func (m *mysqlTaskRepo) GetTasksByUsername(ctx context.Context, username string) (
	[]models.Task, error) {
	var tasks []models.Task
	if err := m.DB.Model(&tasks).Where("username = ?", username).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// AddTask 添加任务
func (m *mysqlTaskRepo) AddTask(ctx context.Context, task models.Task) error {
	if err := m.DB.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskAndHosts 根据TaskId与HostId获取任务与主机
func (m *mysqlTaskRepo) GetTaskAndHosts(ctx context.Context, taskId int64, hostIds []int64) (
	models.Task, []models.Host, error) {
	var task models.Task
	var hosts []models.Host
	if err := m.DB.Model(&task).Where("id = ?", taskId).Error; err != nil {
		return task, hosts, err
	}
	if err := m.DB.Model(&models.Host{}).Where("id in (?)", hostIds).
		Find(&hosts).Error; err != nil {
		return task, hosts, err
	}
	return task, hosts, nil
}

// UpdateTaskStatus 更新任务状态
func (m *mysqlTaskRepo) UpdateTaskStatus(ctx context.Context, taskId int64, status int) error {
	if err := m.DB.Model(&models.Task{}).Where("id = ?", taskId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTaskEntryId 更新任务EntryId
func (m *mysqlTaskRepo) UpdateTaskEntryId(ctx context.Context, taskId int64, entryIds string) error {
	if err := m.DB.Model(&models.Task{}).Where("id = ?", taskId).Update("entry_ids", entryIds).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskByID 根据ID获取任务
func (m *mysqlTaskRepo) GetTaskByID(ctx context.Context, id int64) (models.Task, error) {
	var task models.Task
	if err := m.DB.Model(&task).Where("id = ?", id).Error; err != nil {
		return task, err
	}
	return task, nil
}
