package service

import (
	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/zlog"
)

// const sudoPrefix, sudoSuffix = "[sudo] password for ", ": "
// const sudoPrefixLen = len(sudoPrefix)

// type SSHTerminal struct {
// 	Session            *ssh.Session
// 	exitMsg            string
// 	stdout             io.Reader
// 	stdin              io.Writer
// 	stderr             io.Reader
// 	Password           string
// 	LoginUser          string
// 	EnableSudoPassword bool
// }

func RunSSHTerminal(h models.Host) error {
	client, err := NewSSHClient(h)
	if err != nil {
		return err
	}
	defer client.Close()
	str, err := RunCommand(client, "cd /home;ls")
	zlog.Info(str)
	return err
	// session, err := client.NewSession()
	// if err != nil {
	// 	return err
	// }
	// defer session.Close()

	// s := SSHTerminal{
	// 	Session:            session,
	// 	Password:           h.Password,
	// 	LoginUser:          h.Username,
	// 	EnableSudoPassword: true,
	// }
	// return s.interactiveSession()
}

// func (s *SSHTerminal) interactiveSession() error {
// 	defer func() {
// 		if s.exitMsg != "" {
// 			fmt.Fprintln(os.Stdout, "[Felix]: the connection was closed on the remote side on ", time.Now().Format(time.RFC822))
// 		} else {
// 			fmt.Fprintln(os.Stdout, s.exitMsg)
// 		}
// 	}()
// 	// terminal size
// 	termHeight, termWidth := 80, 120
// 	termType := "msys"
// 	// 判断输出是否为terminal
// 	if isatty.IsTerminal(os.Stdout.Fd()) {
// 		// 获取输入
// 		fd := int(os.Stdin.Fd())

// 		// 设置terminal为行模式并获取terminal的设置前状态
// 		state, err := terminal.MakeRaw(fd)
// 		if err != nil {
// 			return err
// 		}
// 		defer terminal.Restore(fd, state)

// 		// 获取输出
// 		fdOut := int(os.Stdout.Fd())

// 		// 获取terminal大小
// 		termWidth, termHeight, err = terminal.GetSize(fdOut)
// 		if err != nil {
// 			return err
// 		}

// 	}
// 	// 关联远程主机
// 	err := s.Session.RequestPty(termType, termHeight, termWidth, ssh.TerminalModes{})
// 	if err != nil {
// 		return err
// 	}

// 	// 更新terminal大小
// 	// s.updateTerminalSize()

// 	// 获取输入输出管道
// 	s.stdin, err = s.Session.StdinPipe()
// 	if err != nil {
// 		return err
// 	}
// 	s.stdout, err = s.Session.StdoutPipe()
// 	if err != nil {
// 		return err
// 	}
// 	s.stderr, err = s.Session.StderrPipe()

// 	go io.Copy(os.Stderr, s.stderr)
// 	// 密码模式输入密码
// 	if s.EnableSudoPassword {
// 		go enableSudoPassword(s)
// 	} else {
// 		go io.Copy(os.Stdout, s.stdout)
// 	}
// 	go io.Copy(s.stdin, os.Stdin)

// 	// 开始会话
// 	err = s.Session.Shell()
// 	if err != nil {
// 		return err
// 	}
// 	return s.Session.Wait()
// }

// func enableSudoPassword(t *SSHTerminal) {
// 	var (
// 		line string
// 		r    = bufio.NewReader(t.stdout)
// 	)
// 	for {
// 		b, err := r.ReadByte()
// 		if err != nil {
// 			break
// 		}
// 		line += string(b)
// 		os.Stdout.Write([]byte{b})
// 		if b == byte('\n') {
// 			line = ""
// 			continue
// 		}
// 		if len(line) >= sudoPrefixLen && strings.HasPrefix(line, sudoPrefix) && strings.HasSuffix(line, sudoSuffix) && strings.Contains(line, t.LoginUser) {
// 			_, err = t.stdin.Write([]byte(t.Password + "\n"))
// 			if err != nil {
// 				break
// 			}
// 			color.Green("\r\nFelix has automatically input password for %s", color.BlueString(t.LoginUser))
// 		}
// 	}
// }
