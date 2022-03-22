package service

import (
	"context"
	"sort"

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
func (s *execHistoryService) GetExecHistoriesByTaskID(ctx context.Context, taskID int64) ([][]models.ExecHistory, error) {
	execHistories, err := s.execHistoryRepo.GetExecHistoryByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	return splitSlice(execHistories), nil
}

//按某个字段排序
type sortByExecID []models.ExecHistory

func (s sortByExecID) Len() int           { return len(s) }
func (s sortByExecID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByExecID) Less(i, j int) bool { return s[i].ExecId < s[j].ExecId }

//切片分组
func splitSlice(list []models.ExecHistory) [][]models.ExecHistory {
	sort.Sort(sortByExecID(list))
	returnData := make([][]models.ExecHistory, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && list[i].ExecId == list[j].ExecId; j++ {
		}

		returnData = append(returnData, list[i:j])
		i = j
	}
	return returnData
}
