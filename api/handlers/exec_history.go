package handlers

import (
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type ExecHistoryHandler struct {
	execHistoryService service.ExecHistoryService
}

func NewExecHistoryHandler(execHistoryService service.ExecHistoryService) ExecHistoryHandler {
	return ExecHistoryHandler{
		execHistoryService: execHistoryService,
	}
}

// GetExecHistoriesReq 根据任务ID获取执行历史列表
type GetExecHistoriesReq struct {
	TaskID int64 `json:"task_id"`
}

// GetExecHistories 获取执行历史
func (h *ExecHistoryHandler) GetExecHistories(c *gin.Context) {
	var req GetExecHistoriesReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("get exec history error: %s", err.Error())
		response.FailWithData(err, "get exec history error", c)
		return
	}
	execHistories, err := h.execHistoryService.GetExecHistoriesByTaskID(c, req.TaskID)
	if err != nil {
		response.FailWithData(err, "get exec history error", c)
		return
	}
	response.OKWithData(execHistories, "get exec history success", c)
}
