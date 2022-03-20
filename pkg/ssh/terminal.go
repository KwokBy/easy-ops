package ssh

import (
	"fmt"

	"github.com/KwokBy/easy-ops/models"
)

func RunSSHTerminal(h models.Host) error {
	client, err := NewSSHClient(h)
	if err != nil {
		return err
	}
	defer client.Close()
	str, err := RunCommand(client, "cd /home;ls")
	fmt.Println(str)
	return err
}
