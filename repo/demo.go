package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlDemoRepo struct {
	DB *gorm.DB
}

// NewMysqlDemoRepo returns a new instance of mysqlDemoRepo
func NewMysqlDemoRepo(DB *gorm.DB) IDemoRepo {
	return &mysqlDemoRepo{DB}
}

// GetDemos returns all demos from the database
func (m *mysqlDemoRepo) GetDemos(ctx context.Context) ([]models.Demo, error) {
	var demos []models.Demo
	if err := m.DB.Find(&demos).Error; err != nil {
		return nil, err
	}
	return demos, nil
}
