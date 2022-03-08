package repo

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_mysqlMirrorRepo_GetMirrorsByAdmin(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"root",
		"Gl987963951",
		"127.0.0.1",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	type args struct {
		ctx   context.Context
		admin string
	}
	tests := []struct {
		name string
		m    *mysqlMirrorRepo
		args args
		// want    []models.Mirror
		wantErr bool
	}{
		{
			name: "Test_mysqlMirrorRepo_GetMirrorsByAdmin",
			m:    &mysqlMirrorRepo{db},
			args: args{
				ctx:   context.Background(),
				admin: "doubleguo",
			},
			// want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetMirrorsByAdmin(tt.args.ctx, tt.args.admin)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlMirrorRepo.GetMirrorsByAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
