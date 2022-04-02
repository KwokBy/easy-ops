package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlRoleRepo struct {
	DB *gorm.DB
}

func NewMysqlRoleRepo(DB *gorm.DB) RoleRepo {
	return &mysqlRoleRepo{DB}
}

// GetRoleByID 根据ID获取角色
func (r *mysqlRoleRepo) GetRoleByID(ctx context.Context, id int64) (models.Role, error) {
	var role models.Role
	if err := r.DB.Where("role_id = ?", id).Find(&role).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}

// AddRole 添加角色
func (r *mysqlRoleRepo) AddRole(ctx context.Context, role models.Role) error {
	if err := r.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRole 删除角色
func (r *mysqlRoleRepo) DeleteRole(ctx context.Context, id int64) error {
	if err := r.DB.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetRoles 获取所有角色
func (r *mysqlRoleRepo) GetRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// UpdateRole 更新角色
func (r *mysqlRoleRepo) UpdateRole(ctx context.Context, role models.Role) error {
	if err := r.DB.Save(role).Error; err != nil {
		return err
	}
	return nil
}
