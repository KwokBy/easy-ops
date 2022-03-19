package ssh

import (
	"github.com/KwokBy/easy-ops/models"
)

func RunSSHTerminal(h models.Host) error {
	client, err := NewSSHClient(h)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = RunCommand(client, "cd /home;ls")
	return err
}

