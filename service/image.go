package service

import (
	"fmt"
	"io"
	"time"

	"github.com/KwokBy/easy-ops/pkg/docker"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/docker/docker/api/types"
)

type imageService struct {
	docker *docker.Docker
}

func NewImageService(docker *docker.Docker) ImageService {
	return &imageService{
		docker: docker,
	}
}

// DebugImage 调试镜像
func (i *imageService) DebugImage(imageName string, dockerTarFile io.Reader) (types.HijackedResponse, error) {
	imageName = fmt.Sprintf("troublekwok/%s", imageName)
	// 创建镜像
	if err := i.docker.CreateImageByDockerFile(dockerTarFile, imageName, "troublekwok"); err != nil {
		zlog.Errorf("create image error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	// 创建执行容器
	containerName := fmt.Sprintf("%s-%s", imageName, time.Now().Format("20060102150405"))
	if err := i.docker.CreateContainer(containerName, imageName); err != nil {
		zlog.Errorf("create container error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	// 获取容器连接
	hr, err := i.docker.ExecContainer(containerName)
	if err != nil {
		zlog.Errorf("exec container error: %s", err.Error())
		return types.HijackedResponse{}, err
	}
	return hr, err
}

// SaveImage 保存镜像
func (i *imageService) SaveImage(imageName string, dockerfile string) error {
	// TODO 存到本地仓库

	imageName = fmt.Sprintf("troublekwok/%s", imageName)
	// 上传到远程
	if err := i.docker.PushImageToRegistry(imageName, "troublekwok", "Gl@987963951"); err != nil {
		zlog.Errorf("save image error: %s", err.Error())
		return err
	}
	return nil
}

// TODO 从自己数据库查询
// GetImages 获取镜像列表

