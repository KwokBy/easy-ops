package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(DB *gorm.DB) UserRepo {
	return &mysqlUserRepo{DB}
}

// GetUsersByNameAndPwd 根据用户名和密码获取用户
func (u *mysqlUserRepo) GetUsersByNameAndPwd(ctx context.Context, name, pwd string) (
	models.User, error) {
	var user models.User
	if err := u.DB.Where("username = ? and password_hash = ?", name, pwd).Find(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUser 更新用户信息
func (u *mysqlUserRepo) UpdateUser(ctx context.Context, user models.User) error {
	if err := u.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// AddUser 添加用户
func (u *mysqlUserRepo) AddUser(ctx context.Context, user models.User) error {
	if err := u.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (u *mysqlUserRepo) DeleteUser(ctx context.Context, id int64) error {
	if err := u.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserByName 删除用户
func (u *mysqlUserRepo) DeleteUserByName(ctx context.Context, name string) error {
	if err := u.DB.Delete(&models.User{}, "username = ?", name).Error; err != nil {
		return err
	}
	return nil
}

// GetUsers 获取用户列表
func (u *mysqlUserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByName 获取用户
func (u *mysqlUserRepo) GetUserByName(ctx context.Context, name string) (
	models.User, error) {
	var user models.User
	if err := u.DB.Where("username = ?", name).Find(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
