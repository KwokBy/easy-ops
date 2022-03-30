package repo

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlImageRepo struct {
	DB *gorm.DB
}

func NewMysqlImageRepo(DB *gorm.DB) ImageRepo {
	return &mysqlImageRepo{DB}
}

// GetImageByOwner 根据用户名获取镜像列表
func (m *mysqlImageRepo) GetImageByOwner(ctx context.Context, username string) (
	[]models.Image, error) {
	var images []models.Image
	if err := m.DB.Where("owner = ?", username).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

// AddImage 添加镜像
func (m *mysqlImageRepo) AddImage(ctx context.Context, image models.Image) error {
	if err := m.DB.Create(&image).Error; err != nil {
		return err
	}
	return nil
}

// GetImageByImageID 根据镜像ID获取镜像
func (m *mysqlImageRepo) GetImageByImageID(ctx context.Context, imageID int64) (
	[]models.Image, error) {
	var images []models.Image
	if err := m.DB.Where("image_id = ?", imageID).Find(&images).Error; err != nil {
		return images, err
	}
	return images, nil
}

// GetImageByID 根据ID获取镜像
func (m *mysqlImageRepo) GetImageByID(ctx context.Context, id int) (
	models.Image, error) {
	var image models.Image
	if err := m.DB.Where("id = ?", id).Find(&image).Error; err != nil {
		return image, err
	}
	return image, nil
}
