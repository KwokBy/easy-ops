package models

import (
	"testing"
	"time"
)

func TestTask_ToDTO(t *testing.T) {
	tests := []struct {
		name    string
		tr      *Task
		wantErr bool
	}{
		{
			name: "test",
			tr: &Task{
				ID:          1,
				Username:    "test",
				Name:        "test",
				Content:     "test",
				FileType:    "test",
				HostIDs:     "[1,2]",
				Spec:        "test",
				ExecIds:     "[1,2]",
				Status:      1,
				Desc:        "test",
				UpdatedTime: time.Now(),
				CreatedTime: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.ToDTO()
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.ToDTO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Task.ToDTO() = %v", got)
		})
	}
}
