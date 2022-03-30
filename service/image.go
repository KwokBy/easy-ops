package service

import (
	"context"
	"fmt"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/docker"
	myDocker "github.com/KwokBy/easy-ops/pkg/docker"
	"github.com/KwokBy/easy-ops/pkg/str"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/repo"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type imageService struct {
	docker    *docker.Docker
	imageRepo repo.ImageRepo
}

func NewImageService(repo repo.ImageRepo) ImageService {
	cli, _ := client.NewClientWithOpts(client.FromEnv)
	return &imageService{
		docker:    myDocker.NewDocker(cli),
		imageRepo: repo,
	}
}

// DebugImage 调试镜像
func (i *imageService) DebugImage(ctx context.Context, id int) (types.HijackedResponse, error) {
	image, err := i.imageRepo.GetImageByID(ctx, id)
	if err != nil {
		return types.HijackedResponse{}, err
	}
	zlog.Infof("debug image: %s", image.Name)
	tag := fmt.Sprintf("troublekwok/%s:%s", image.Name, image.Version)
	zlog.Infof("tag: %s", tag)
	// 创建镜像
	if err := i.docker.CreateImageByDockerFile(image.Dockerfile, tag, "troublekwok"); err != nil {
		zlog.Errorf("create image error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	// 创建执行容器
	containerName := fmt.Sprintf("%s-%s", image.Name, time.Now().Format("20060102150405"))
	if err := i.docker.CreateContainer(tag, containerName); err != nil {
		zlog.Errorf("create container error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	// containerName := "nervous_bhaskara"
	// 获取容器连接
	hr, err := i.docker.ExecContainer(containerName)
	if err != nil {
		zlog.Errorf("exec container error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	return hr, err
}

// SaveImage 保存镜像
func (i *imageService) SaveImage(ctx context.Context, image models.Image) error {
	// 存到本地数据库
	image.CreatedTime = time.Now()
	image.UpdatedTime = time.Now()
	image.Version = str.VersionIncrease(image.Version)
	if image.ImageID == "" {
		image.ImageID = fmt.Sprintf("%s-%s", image.Name, time.Now().Format("20060102150405"))
		image.Owner = "troublekwok"
	}
	if err := i.imageRepo.AddImage(ctx, image); err != nil {
		zlog.Errorf("add image error: %s", err.Error())
		return err
	}
	imageName := fmt.Sprintf("troublekwok/%s:%s", image.Name, image.Version)
	// 创建镜像
	if err := i.docker.CreateImageByDockerFile(image.Dockerfile, imageName, imageName); err != nil {
		zlog.Errorf("create image error: %s", err.Error())
		return err
	}
	// 上传到远程
	if err := i.docker.PushImageToRegistry(imageName, "troublekwok", "Gl@987963951"); err != nil {
		zlog.Errorf("save image error: %s", err.Error())
		return err
	}
	return nil
}

// GetImages 获取镜像列表
func (i *imageService) GetImages(ctx context.Context, username string) ([]models.Image, error) {
	return i.imageRepo.GetImageByOwner(ctx, username)
}
