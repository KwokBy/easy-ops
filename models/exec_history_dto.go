package models

type ExecHistoryDTO struct {
	AvgCount      float64       `json:"avg_count"`
	FailCount     int           `json:"fail_count"`
	SuccessCount  int           `json:"success_count"`
	ExecHistories []ExecHistoryInfo `json:"exec_history"`
}

