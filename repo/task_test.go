package repo

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_mysqlTaskRepo_GetTaskAndHost(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"easy_ops",
		"Gl@987963951",
		"42.192.11.9",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	type args struct {
		ctx     context.Context
		taskId  int64
		hostIds []int64
	}
	tests := []struct {
		name string
		m    *mysqlTaskRepo
		args args
		// want    models.Task
		// want1   []models.Host
		wantErr bool
	}{
		{
			name: "Test_mysqlMirrorRepo_GetMirrorsByAdmin",
			m:    &mysqlTaskRepo{db},
			args: args{
				ctx:     context.Background(),
				hostIds: []int64{1, 2, 3},
				taskId:  1,
			},
			// want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.m.GetTaskAndHosts(tt.args.ctx, tt.args.taskId, tt.args.hostIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTaskRepo.GetTaskAndHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			t.Log(got1)
		})
	}
}

func Test_mysqlTaskRepo_GetTasksByUsername(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"easy_ops",
		"Gl@987963951",
		"42.192.11.9",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		m       *mysqlTaskRepo
		args    args
		wantErr bool
	}{
		{
			name: "Test_mysqlMirrorRepo_GetMirrorsByAdmin",
			m:    &mysqlTaskRepo{db},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			// want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetTaskByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTaskRepo.GetTasksByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
