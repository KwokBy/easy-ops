package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
)

// ExecContainer 运行容器并返回连接
func ExecContainer(client *client.Client, imageName string, cmd []string) (types.HijackedResponse, error) {
	// containerName := fmt.Sprintf("%s-%s", imageName, time.Now().Format("20060102150405"))

	// _, err := client.ContainerCreate(context.Background(), &container.Config{
	// 	Image:        imageName,
	// 	Cmd:          []string{"echo", "hello world"},
	// 	Tty:          false,
	// 	AttachStdin:  true,
	// 	AttachStdout: true,
	// 	AttachStderr: true,
	// }, nil, nil, nil, containerName)
	// if err != nil {
	// 	panic(err)
	// }
	containerName := "test001"
	err := client.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	// 创建一个新的 exec 配置来运行一个 exec 进程。
	cli, err := client.ContainerExecCreate(context.Background(), containerName, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		//[]string{"/bin/bash", "-c", "echo hello"},
		Cmd: cmd,
	})
	if err != nil {
		panic(err)
	}
	// 将与container的连接附加到上面的exec进程
	hr, err := client.ContainerExecAttach(context.Background(), cli.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	})
	if err != nil {
		panic(err)
	}
	return hr, nil
}

// WsReaderCopy 将前端的输入转发到终端
func WsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}

// WsWriterCopy 将终端的输出转发到前端
func WsWriterCopy(writer *websocket.Conn, reader io.Reader) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}
