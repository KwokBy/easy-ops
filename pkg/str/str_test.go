package str

import "testing"

func TestVersionIncrease(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1.0.0",
			args: args{
				str: "1.0.9",
			},
			want: "1.1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionIncrease(tt.args.str); got != tt.want {
				t.Errorf("VersionIncrease() = %v, want %v", got, tt.want)
			}
		})
	}
}
