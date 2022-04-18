package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_mysqlHostRepo_AddHost(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"easy_ops",
		"Gl@987963951",
		"42.192.11.9",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	type args struct {
		ctx  context.Context
		host models.Host
	}
	tests := []struct {
		name    string
		h       *mysqlHostRepo
		args    args
		wantErr bool
	}{
		{
			name: "Test_mysqlHostRepo_AddHost",
			h:    &mysqlHostRepo{db},
			args: args{
				ctx: context.Background(),
				host: models.Host{
					HostName:    "106.55.161.12",
					Host:        "106.55.161.12",
					Port:        22,
					Owner:    "doubleguo",
					Password:    "Gl@987963951",
					SSHType:     "password",
					Name:        "root",
					Desc:        "kwok cvm",
					CreatedTime: time.Now(),
					UpdatedTime: time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.AddHost(tt.args.ctx, tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("mysqlHostRepo.AddHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
