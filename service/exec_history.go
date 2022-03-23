package service

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/repo"
)

type execHistoryService struct {
	execHistoryRepo repo.ExecHistoryRepo
}

func NewExecHistoryService(execHistoryRepo repo.ExecHistoryRepo) ExecHistoryService {
	return &execHistoryService{
		execHistoryRepo: execHistoryRepo,
	}
}

// GetExecHistoriesByTaskID 根据任务ID获取执行历史列表
func (s *execHistoryService) GetExecHistoriesByTaskID(ctx context.Context, taskID int64) ([]models.ExecHistory, error) {
	return s.execHistoryRepo.GetExecHistoryByTaskID(ctx, taskID)
}
