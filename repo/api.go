package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlApiRepo struct {
	DB *gorm.DB
}

func NewMysqlApiRepo(DB *gorm.DB) ApiRepo {
	return &mysqlApiRepo{DB}
}

// GetApis 获取所有接口
func (r *mysqlApiRepo) GetApis(ctx context.Context) ([]models.Api, error) {
	var apis []models.Api
	if err := r.DB.Find(&apis).Error; err != nil {
		return nil, err
	}
	return apis, nil
}

// GetApisByID 根据ID获取接口
func (r *mysqlApiRepo) GetApisByID(ctx context.Context, ids []int) ([]models.Api, error) {
	var apis []models.Api
	if err := r.DB.Where("id in (?)", ids).Find(&apis).Error; err != nil {
		return []models.Api{}, err
	}
	return apis, nil
}
