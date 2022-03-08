package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlTaskRepo struct {
	DB *gorm.DB
}

func NewMysqlTaskRepo(DB *gorm.DB) ITaskRepo {
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
