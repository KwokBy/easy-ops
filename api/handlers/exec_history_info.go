package handlers

import (
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type ExecHistoryInfoHandler struct {
	execHistoryInfoService service.ExecHistoryInfoService
}

func NewExecHistoryInfoHandler(execHistoryInfoService service.ExecHistoryInfoService) ExecHistoryInfoHandler {
	return ExecHistoryInfoHandler{
		execHistoryInfoService: execHistoryInfoService,
	}
}

type GetExecHistoryReq struct {
	TaskID int64 `json:"task_id"`
	ExecID int64 `json:"exec_id"`
}

// GetExecHistory 获取执行历史
func (h *ExecHistoryInfoHandler) GetExecHistoryInfo(c *gin.Context) {
	var req GetExecHistoryReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("get exec history error: %s", err.Error())
		response.FailWithData(err, "get exec history error", c)
		return
	}
	execHistory, err := h.execHistoryInfoService.GetExecHistoryInfos(c, req.TaskID, req.ExecID)
	if err != nil {
		zlog.Errorf("get exec history error: %s", err.Error())
		response.FailWithData(err, "get exec history error", c)
		return
	}
	response.OKWithData(execHistory, "get exec history success", c)
}
