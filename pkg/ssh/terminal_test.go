package ssh

import (
	"testing"

	"github.com/KwokBy/easy-ops/models"
)

func TestRunSSHTerminal(t *testing.T) {
	type args struct {
		h models.Host
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestRunSSHTerminal",
			args: args{
				h: models.Host{
					HostName: "106.55.161.12",
					Host:     "106.55.161.12",
					Port:     22,
					Owner: "root",
					Password: "Gl@987963951",
					SSHType:  "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunSSHTerminal(tt.args.h); (err != nil) != tt.wantErr {
				t.Errorf("RunSSHTerminal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
