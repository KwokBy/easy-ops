package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
)

type Docker struct {
	client *client.Client
}

func NewDocker(c *client.Client) *Docker {
	return &Docker{
		client: c,
	}
}

// CreateImageByDockerFile 创建镜像
func (d *Docker) CreateImageByDockerFile(dockerFile string, imageName, project string) error {
	dockerTarFile, err := createDockerTarFile(dockerFile)
	if err != nil {
		return err
	}
	// 创建镜像
	output, err := d.client.ImageBuild(context.Background(), dockerTarFile, types.ImageBuildOptions{
		Tags:       []string{imageName},
		Dockerfile: "dockerfile",
		Labels: map[string]string{
			project: "project",
		},
	})
	if err != nil {
		return err
	}
	defer output.Body.Close()

	// 读取镜像的输出
	body, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return err
	}
	// 判断构建是否成功
	if strings.Contains(string(body), "error") {
		return errors.New("build image to docker error")
	}
	return nil
}

// CreateAndRunContainer 创建并运行容器
func (d *Docker) CreateContainer(imageName, containerName string) error {
	// 创建容器
	_, err := d.client.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		// Cmd:          []string{"/bin/bash"},
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}, nil, nil, nil, containerName)
	if err != nil {
		return err
	}
	// 启动容器
	err = d.client.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// ExecContainer 运行容器并返回连接
func (d *Docker) ExecContainer(containerName string) (types.HijackedResponse, error) {
	// 创建一个新的 exec 配置来运行一个 exec 进程。
	cli, err := d.client.ContainerExecCreate(context.Background(), containerName, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"/bin/bash"},
	})
	if err != nil {
		panic(err)
	}
	// 将与container的连接附加到上面的exec进程
	hr, err := d.client.ContainerExecAttach(context.Background(), cli.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	})
	if err != nil {
		panic(err)
	}

	return hr, nil
}

// PushImageToRegistry 将镜像推送到镜像仓库
func (d *Docker) PushImageToRegistry(imageName, username, password string) error {
	authConfig := types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: "https://index.docker.io/v1/",
	}
	encodeAuth, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	// 将镜像推送到镜像仓库
	output, err := d.client.ImagePush(context.Background(), imageName, types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(encodeAuth),
	})
	if err != nil {
		return err
	}
	defer output.Close()
	// 读取镜像的输出
	body, err := ioutil.ReadAll(output)
	if err != nil {
		return err
	}
	// 判断构建是否成功
	if strings.Contains(string(body), "error") {
		return errors.New("push image to registry error")
	}
	return nil
}

// createDockerTarFile 创建dockerfile
func createDockerTarFile(dockerfile string) (io.Reader, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: "dockerfile",
		Mode: 0600,
		Size: int64(len([]byte(dockerfile))),
	}

	tw.WriteHeader(hdr)
	_, err := tw.Write([]byte(dockerfile))
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

type wsMsg struct {
	Type int    `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

const (
	wsMsgCmd    = 1
	wsMsgResize = 2
)

// WsReaderCopy 将前端的输入转发到终端
func WsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		_, data, err := reader.ReadMessage()
		if err != nil {
			return
		}
		var msg wsMsg
		if err = json.Unmarshal(data, &msg); err != nil {
			return
		}
		switch msg.Type {
		case wsMsgCmd:
			decodeBytes, _ := base64.StdEncoding.DecodeString(msg.Cmd)
			writer.Write(decodeBytes)
		}
	}
}

// WsWriterCopy 将终端的输出转发到前端
func WsWriterCopy(writer *websocket.Conn, reader io.Reader) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.TextMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

// DeleteImage 删除镜像
func (d *Docker) DeleteImage(imageName string) error {
	// 删除镜像
	_, err := d.client.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}
	return nil
}
