package ssh

import (
	"testing"

	"github.com/KwokBy/easy-ops/models"
)

func TestClientAndExec(t *testing.T) {
	type args struct {
		host models.Host
		cmd  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestClientAndExec",
			args: args{
				host: models.Host{
					HostName: "106.55.161.12",
					Host:     "106.55.161.12:22",
					Port:     22,
					Owner:    "root",
					User:     "root",
					Password: "Gl@987963951",
					SSHType:  "ssh-password",
				},
				cmd: "cd /home;ls",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClientAndExec(tt.args.host, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientAndExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
