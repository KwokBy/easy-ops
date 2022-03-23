package repo

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_mysqlExecHistoryRepo_GetCountGroupByExecID(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"root",
		"Gl987963951",
		"127.0.0.1",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	type args struct {
		ctx    context.Context
		taskID int64
	}
	tests := []struct {
		name    string
		h       *mysqlExecHistoryInfoRepo
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_mysqlExecHistoryRepo_GetCountGroupByExecID",
			h:    &mysqlExecHistoryInfoRepo{db},
			args: args{
				ctx:    context.Background(),
				taskID: 1,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetCountGroupByExecID(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlExecHistoryRepo.GetCountGroupByExecID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlExecHistoryRepo.GetCountGroupByExecID() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
