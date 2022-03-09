package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

// NewSSHClient 创建ssh客户端
func NewSSHClient(h models.Host) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		// ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		Timeout:         time.Second,
		User:            h.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// 判断登录类型
	if h.SSHType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.SSHKeyPath)}
	}
	addr := fmt.Sprintf("%s:%d", h.Host, h.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		zlog.Errorf("[NewSSHClient] Dial error: %s", err.Error())
		return nil, err
	}
	return client, nil
}

// RunCommand 执行命令
func RunCommand(client *ssh.Client, cmd string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		zlog.Errorf("[RunCommand] NewSession error: %s", err.Error())
		return "", err
	}
	defer session.Close()
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		zlog.Errorf("RunCommand error: %s", err)
		return "", err
	}
	return string(output), nil
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
