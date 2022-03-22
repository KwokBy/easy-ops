package models

type ExecHistoryDTO struct {
	AvgCount    float64     `json:"avg_count"`
	ExecHistories []ExecHistory `json:"exec_history"`
}
