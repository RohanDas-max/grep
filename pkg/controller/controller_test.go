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
			name: "pass",
			args: args{
				args: []string{"hello", "../../dir"},
			},
			wantErr: false,
		}, {
			name: "fail",
			args: args{
				args: []string{"hello", "dir"},
			},
			wantErr: true,
		}, {
			name: "pass",
			args: args{
				args: []string{"hello", "../../test.txt"},
			},
			wantErr: false,
		}, {
			name: "pass",
			args: args{
				args: []string{"hello"},
			},
			wantErr: false,
		}, {
			name: "pass",
			args: args{
				args: []string{"hello", "any", "any"},
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
