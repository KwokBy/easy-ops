package models

import (
	"time"

	"github.com/KwokBy/easy-ops/pkg/str"
)

type TaskDTO struct {
	ID          int64     `json:"id" form:"id"`
	Username    string    `json:"username" form:"username"`
	Name        string    `json:"name" form:"name"`
	Content     string    `json:"content" form:"content"`
	FileType    string    `json:"file_type" form:"file_type"`
	HostIDs     []int64   `json:"host_ids" form:"host_ids"`
	Spec        string    `json:"spec" form:"spec"`
	ExecIds     []int64   `json:"exec_ids" form:"exec_ids"`
	Status      int       `json:"status" form:"status"`
	Desc        string    `json:"desc" form:"desc"`
	UpdatedTime time.Time `json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `json:"created_time" form:"created_time"`
}

func (t *TaskDTO) ToPOJO() (Task, error) {
	hostIDs, err := str.Int64s2String(t.HostIDs)
	if err != nil {
		return Task{}, err
	}
	execIDs, err := str.Int64s2String(t.ExecIds)
	if err != nil {
		return Task{}, err
	}
	return Task{
		ID:          t.ID,
		Username:    t.Username,
		Name:        t.Name,
		Content:     t.Content,
		FileType:    t.FileType,
		HostIDs:     hostIDs,
		Spec:        t.Spec,
		ExecIds:     execIDs,
		Status:      t.Status,
		Desc:        t.Desc,
		UpdatedTime: t.UpdatedTime,
		CreatedTime: t.CreatedTime,
	}, err
}


