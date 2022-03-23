package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlExecHistoryInfoRepo struct {
	DB *gorm.DB
}

func NewMysqlExecHistoryInfoRepo(DB *gorm.DB) ExecHistoryInfoRepo {
	return &mysqlExecHistoryInfoRepo{DB}
}

// AddExecHistory 添加执行历史
func (h *mysqlExecHistoryInfoRepo) AddExecHistory(ctx context.Context, execHistory models.ExecHistoryInfo) error {
	if err := h.DB.Create(&execHistory).Error; err != nil {
		return err
	}
	return nil
}

// BatchAddExecHistory 批量添加执行历史
func (h *mysqlExecHistoryInfoRepo) BatchAddExecHistory(ctx context.Context, execHistories []models.ExecHistoryInfo) error {
	if err := h.DB.Create(&execHistories).Error; err != nil {
		return err
	}
	return nil
}

// GetExecHistoryByTaskID 根据TaskID获取执行历史
func (h *mysqlExecHistoryInfoRepo) GetExecHistoryByTaskID(ctx context.Context, taskID int64) (
	[]models.ExecHistoryInfo, error) {
	var execHistories []models.ExecHistoryInfo
	if err := h.DB.Where("task_id = ?", taskID).Find(&execHistories).Error; err != nil {
		return nil, err
	}
	return execHistories, nil
}

type TempCount struct {
	ExecID int64 `gorm:"column:exec_id"`
	Total  int64 `gorm:"column:total"`
	TaskID int64 `gorm:"column:task_id"`
}

// GetCountGroupByExecID 获取某个Task下的执行次数
func (h *mysqlExecHistoryInfoRepo) GetCountGroupByExecID(ctx context.Context, taskID int64) (int, error) {
	var tempCounts []TempCount
	if err := h.DB.Model(&models.ExecHistoryInfo{}).Select("task_id, exec_id, count(*) as  total").
		Where("task_id = ?", taskID).Group("exec_id,task_id").Scan(&tempCounts).Error; err != nil {
		return 0, err
	}
	return len(tempCounts), nil
}
