package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlCasbinRepo struct {
	DB *gorm.DB
}

func NewMysqlCasbinRepo(DB *gorm.DB) CasbinRepo {
	return &mysqlCasbinRepo{DB}
}

// GetByRoleID 获取角色权限
func (r *mysqlCasbinRepo) GetByRoleID(ctx context.Context, roleID int) ([]models.Casbin, error) {
	var casbins []models.Casbin
	if err := r.DB.Where("v0 = ?", roleID).Find(&casbins).Error; err != nil {
		return []models.Casbin{}, err
	}
	return casbins, nil
}
