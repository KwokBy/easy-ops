package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlHostRepo struct {
	DB *gorm.DB
}

func NewMysqlHostRepo(DB *gorm.DB) HostRepo {
	return &mysqlHostRepo{DB}
}

func (h *mysqlHostRepo) TableName() string {
	return "t_host"
}

// GetHostsByUsername 根据用户名获取主机列表
func (h *mysqlHostRepo) GetHostsByUsername(ctx context.Context, owner string) (
	[]models.Host, error) {
	var hosts []models.Host
	if err := h.DB.Where("owner = ?", owner).Find(&hosts).Error; err != nil {
		return nil, err
	}
	return hosts, nil
}

// AddHost 添加主机
func (h *mysqlHostRepo) AddHost(ctx context.Context, host models.Host) error {
	if err := h.DB.Create(&host).Error; err != nil {
		return err
	}
	return nil
}

// DeleteHost 删除主机
func (h *mysqlHostRepo) DeleteHost(ctx context.Context, id int64) error {
	if err := h.DB.Delete(&models.Host{}, id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateHost 更新主机信息
func (h *mysqlHostRepo) UpdateHost(ctx context.Context, host models.Host) error {
	if err := h.DB.Save(&host).Error; err != nil {
		return err
	}
	return nil
}
