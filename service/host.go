package service

import (
	"context"
	"fmt"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/ssh"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/repo"
)

type hostService struct {
	hostRepo repo.HostRepo
}

func NewHostService(hostRepo repo.HostRepo) HostService {
	return &hostService{
		hostRepo: hostRepo,
	}
}

// GetHostsByUsername 根据用户名获取主机列表
func (h *hostService) GetHostsByUsername(ctx context.Context, owner string) (
	[]models.Host, error) {
	hosts, err := h.hostRepo.GetHostsByUsername(ctx, owner)
	if err != nil {
		zlog.Errorf("get hosts by owner error: %s", err.Error())
		return nil, err
	}
	return hosts, nil
}

const (
	hostVerifySucc = 1
	hostVerifyFail = 0
)

// AddHost 添加主机
func (h *hostService) AddHost(ctx context.Context, host models.Host) error {
	host.UpdatedTime = time.Now()
	host.CreatedTime = time.Now()
	// TODO 使用前端传入的用户名
	host.Owner = "doubleguo"
	// TODO 查询主机信息

	host.Host = fmt.Sprintf("%s:%d", host.HostName, host.Port)
	// 使用公秘钥
	host.SSHType = "ssh-key"
	if host.Password != "" {
		host.SSHType = "ssh-password"
	}
	if err := h.hostRepo.AddHost(ctx, host); err != nil {
		zlog.Errorf("add host error: %s", err.Error())
		return err
	}
	return nil
}

// DeleteHost 删除主机
func (h *hostService) DeleteHost(ctx context.Context, id int64) error {
	if err := h.hostRepo.DeleteHost(ctx, id); err != nil {
		zlog.Errorf("delete host error: %s", err.Error())
		return err
	}
	return nil
}

// UpdateHost 更新主机信息
func (h *hostService) UpdateHost(ctx context.Context, host models.Host) error {
	host.UpdatedTime = time.Now()
	if err := h.hostRepo.UpdateHost(ctx, host); err != nil {
		zlog.Errorf("update host error: %s", err.Error())
		return err
	}
	return nil
}

// 查询主机信息

// VerifyHost 校验主机信息
func (h *hostService) VerifyHost(ctx context.Context, host models.Host) error {
	_, err := ssh.NewSSHClient(host)
	host.UpdatedTime = time.Now()
	host.Status = hostVerifySucc
	if err != nil {
		host.Status = hostVerifyFail
		zlog.Errorf("verify host error: %s", err.Error())
	}
	if err := h.hostRepo.UpdateHost(ctx, host); err != nil {
		zlog.Errorf("update host error: %s", err.Error())
		return err
	}
	return err
}
