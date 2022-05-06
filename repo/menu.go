package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlMenuRepo struct {
	DB *gorm.DB
}

func NewMysqlMenuRepo(DB *gorm.DB) MenuRepo {
	return &mysqlMenuRepo{DB}
}

// GetMenus 获取所有菜单
func (r *mysqlMenuRepo) GetMenus(ctx context.Context) ([]models.Menu, error) {
	var menus []models.Menu
	if err := r.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// GetMenuByIDs 根据IDs获取菜单
func (r *mysqlMenuRepo) GetMenuByIDs(ctx context.Context, ids []int) ([]models.Menu, error) {
	var menus []models.Menu
	if err := r.DB.Where("id in (?)", ids).Find(&menus).Error; err != nil {
		return []models.Menu{}, err
	}
	return menus, nil
}
