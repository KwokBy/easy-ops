package ssh

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

// NewSSHClient 创建ssh客户端
func NewSSHClient(h models.Host) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		// ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		Timeout:         time.Second,
		User:            h.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// 判断登录类型
	if h.SSHType == "ssh-password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.SSHKeyPath)}
	}
	client, err := ssh.Dial("tcp", h.Host, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// RunCommand 执行命令
func RunCommand(client *ssh.Client, cmd string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	output, err := session.CombinedOutput(cmd)
	if err != nil {
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

// write data to WebSocket
// the data comes from ssh server.
type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

// implement Write interface to write bytes from ssh server into bytes.Buffer.
func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// connect to ssh server using ssh session.
type SSHConn struct {
	// calling Write() to write data into ssh server
	StdinPipe io.WriteCloser
	// Write() be called to receive data from ssh server
	ComboOutput *wsBufferWriter
	Session     *ssh.Session
}

func NewSSHConn(cols, rows int, sshClient *ssh.Client) (*SSHConn, error) {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	// we set stdin, then we can write data to ssh server via this stdin.
	// but, as for reading data from ssh server, we can set Session.Stdout and Session.Stderr
	// to receive data from ssh server, and write back to somewhere.
	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	comboWriter := new(wsBufferWriter)

	//ssh.stdout and stderr will write output into comboWriter
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}

	// start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SSHConn{
		StdinPipe:   stdinP,
		ComboOutput: comboWriter,
		Session:     sshSession,
	}, nil
}

func (s *SSHConn) Close() {
	if s.Session != nil {
		s.Session.Close()
	}
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

// ReceiveWsMsg 处理ws消息并转发给ssh-Session stdinPipe,同时暂存消息到logBuff
func (s *SSHConn) ReceiveWsMsg(wsConn *websocket.Conn, logBuff *bytes.Buffer, exitCh chan bool) {
	defer setQuit(exitCh)
	for {
		select {
		case <-exitCh:
			return
		default:
			// read data from wsConn
			_, data, err := wsConn.ReadMessage()
			zlog.Infof("[ReceiveWsMsg] read data from wsConn: %s", string(data))
			if err != nil {
				zlog.Errorf("[ReceiveWsMsg] ReadMessage error: %s", err.Error())
				return
			}
			// unmarshal data to sshMsg
			var sshMsg wsMsg
			if err := json.Unmarshal(data, &sshMsg); err != nil {
				zlog.Errorf("[Unmarshal] Unmarshal data error: %s", err.Error())
			}
			switch sshMsg.Type {
			// 处理前端窗口大小变化
			case wsMsgResize:
				if sshMsg.Cols > 0 && sshMsg.Rows > 0 {
					if err := s.Session.WindowChange(sshMsg.Cols, sshMsg.Rows); err != nil {
						zlog.Errorf("[ReceiveWsMsg] WindowChange error: %s", err.Error())
						return
					}
				}
			// 处理前端命令
			case wsMsgCmd:
				//handle xterm.js stdin
				decodeBytes, err := base64.StdEncoding.DecodeString(sshMsg.Cmd)
				if err != nil {
					zlog.Errorf("[ReceiveWsMsg] DecodeString error: %s", err.Error())
				}
				// 命令写入到ssh-session-stdin-pipline
				if _, err := s.StdinPipe.Write(decodeBytes); err != nil {
					zlog.Errorf("[ReceiveWsMsg] StdinPipe.Write error: %s", err.Error())
				}
				zlog.Info(s.ComboOutput)
			}
		}
	}
}

func setQuit(exitCh chan bool) {
	exitCh <- true
}

// SendComboOutput 把ssh.Session的comboWriter中的数据每隔120ms
// 通过调用websocketConn.WriteMessage方法返回给前端
func (s *SSHConn) SendComboOutput(wsConn *websocket.Conn, exitCh chan bool) {
	defer setQuit(exitCh)

	// 定时器
	tick := time.NewTicker(time.Millisecond * 120)
	defer tick.Stop()
	for {
		select {
		case <-exitCh:
			return
		case <-tick.C:
			// 如果comboWriter中有数据，则把数据写入到websocketConn
			if err := flushComboOutput(s.ComboOutput, wsConn); err != nil {
				zlog.Errorf("[SendComboOutput] flushComboOutput error: %s", err.Error())
				return
			}
		}
	}
}

func (s *SSHConn) SessionWait(quitChan chan bool) {
	if err := s.Session.Wait(); err != nil {
		zlog.Errorf("[SessionWait] Wait error: %s", err.Error())
		setQuit(quitChan)
	}
}

func flushComboOutput(w *wsBufferWriter, wsConn *websocket.Conn) error {
	if w.buffer.Len() > 0 {
		// 将comboWriter中的数据写入到websocketConn
		err := wsConn.WriteMessage(websocket.TextMessage, w.buffer.Bytes())
		zlog.Info("[flushComboOutput] write data to websocketConn: ", string(w.buffer.Bytes()))
		if err != nil {
			zlog.Errorf("[flushComboOutput] WriteMessage error: %s", err.Error())
			return err
		}
		w.buffer.Reset()
	}
	return nil
}
