package controller

import "testing"

func TestController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass1",
			args: args{
				args: []string{"hello", "../../dir"},
			},
			wantErr: false,
		},
		{
			name: "pass2",
			args: args{
				args: []string{"hello", "../../test.txt"},
			},
			wantErr: false,
		},
		{
			name: "pass3",
			args: args{
				args: []string{"hello"},
			},
			wantErr: false,
		},
		{
			name: "fail1",
			args: args{
				args: []string{"hello", "dir"},
			},
			wantErr: true,
		},
		{
			name: "fail2",
			args: args{
				args: []string{"hello", "any", "any"},
			},
			wantErr: true,
		},
		{
			name: "fail3",
			args: args{
				args: []string{""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
