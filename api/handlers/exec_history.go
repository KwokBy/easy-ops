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

type GetExecHistoryReq struct {
	TaskID int64 `json:"task_id"`
}

// GetExecHistory 获取执行历史
func (h *ExecHistoryHandler) GetExecHistory(c *gin.Context) {
	var req GetExecHistoryReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("get exec history error: %s", err.Error())
		response.FailWithData(err, "get exec history error", c)
		return
	}
	execHistory, err := h.execHistoryService.GetExecHistoriesByTaskID(c, req.TaskID)
	if err != nil {
		zlog.Errorf("get exec history error: %s", err.Error())
		response.FailWithData(err, "get exec history error", c)
		return
	}
	response.OKWithData(execHistory, "get exec history success", c)
}
