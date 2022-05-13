package service

import (
	"context"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/repo"
)

type roleService struct {
	roleRepo repo.RoleRepo
}

func NewRoleService(repo repo.RoleRepo) RoleService {
	return &roleService{
		roleRepo: repo,
	}
}

// GetRoles 获取角色列表
func (s *roleService) GetRoles(ctx context.Context) ([]models.Role, error) {
	return s.roleRepo.GetRoles(ctx)
}

// AddRole 添加角色
func (s *roleService) AddRole(ctx context.Context, role models.Role) error {
	role.CreatedTime = time.Now()
	role.UpdatedTime = time.Now()
	return s.roleRepo.AddRole(ctx, role)
}

// DeleteRole 删除角色
func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	return s.roleRepo.DeleteRole(ctx, id)
}

// UpdateRole 更新角色
func (s *roleService) UpdateRole(ctx context.Context, role models.Role) error {
	return s.roleRepo.UpdateRole(ctx, role)
}

// GetRoleByID 获取角色信息
func (s *roleService) GetRoleByID(ctx context.Context, id int64) (models.Role, error) {
	return s.roleRepo.GetRoleByID(ctx, id)
}

// GetRoleAPIs
func (s *roleService) GetRoleAPIs(ctx context.Context, id int64) ([]models.Casbin, error) {
	return []models.Casbin{}, nil
}

// SetRoleAPIs
func (s *roleService) SetRoleAPIs(ctx context.Context, id int64, apis []models.Casbin) error {
	return nil
}

// GetRoleResources 获取角色资源
func (s *roleService) GetRoleResources(ctx context.Context, id int64) ([]models.AuthMenu, error) {
	return []models.AuthMenu{}, nil
}

// SetRoleResources
func (s *roleService) SetRoleResources(ctx context.Context, id int64, resources []models.AuthMenu) error {
	return nil
}

// GetRoleMenus 获取角色菜单
func (s *roleService) GetRoleMenus(ctx context.Context, id int64) ([]models.Menu, error) {
	return []models.Menu{}, nil
}
