package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlAuthMenuRepo struct {
	DB *gorm.DB
}

func NewMysqlAuthMenuRepo(DB *gorm.DB) AuthMenuRepo {
	return &mysqlAuthMenuRepo{DB}
}

// GetMenusByRoleID 获取角色菜单
func (r *mysqlAuthMenuRepo) GetMenusByRoleID(ctx context.Context, roleID int64) ([]int64, error) {
	var menuIDs []int64
	if err := r.DB.Where("role_id = ?", roleID).Pluck("menu_id", &menuIDs).Error; err != nil {
		return []int64{}, err
	}
	return menuIDs, nil
}

// AddMenuToRole 添加菜单到角色
func (r *mysqlAuthMenuRepo) AddMenuToRole(ctx context.Context, authMenus []models.AuthMenu) error {
	r.DB.Delete(&models.AuthMenu{}, "role_id = ?", authMenus[0].RoleID)
	if err := r.DB.Create(&authMenus).Error; err != nil {
		return err
	}
	return nil
}
