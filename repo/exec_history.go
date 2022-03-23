package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlExecHistoryRepo struct {
	DB *gorm.DB
}

func NewMysqlExecHistoryRepo(DB *gorm.DB) ExecHistoryRepo {
	return &mysqlExecHistoryRepo{DB}
}

// AddExecHistory 添加执行历史
func (h *mysqlExecHistoryRepo) AddExecHistory(ctx context.Context, execHistory models.ExecHistory) error {
	if err := h.DB.Create(&execHistory).Error; err != nil {
		return err
	}
	return nil
}

// GetExecHistoryByTaskID 根据TaskID获取执行历史
func (h *mysqlExecHistoryRepo) GetExecHistoryByTaskID(ctx context.Context, taskID int64) (
	[]models.ExecHistory, error) {
	var execHistories []models.ExecHistory
	if err := h.DB.Where("task_id = ?", taskID).Find(&execHistories).Error; err != nil {
		return nil, err
	}
	return execHistories, nil
}

// GetExecHistoryCountByTaskID 根据TaskID获取执行历史数量
func (h *mysqlExecHistoryRepo) GetExecHistoryCountByTaskID(ctx context.Context, taskID int64) (int64, error) {
	var count int64
	if err := h.DB.Model(&models.ExecHistory{}).Where("task_id = ?", taskID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
} 
