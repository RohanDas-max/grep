package controller

import "testing"

func TestController(t *testing.T) {
	type args struct {
		args []string
		i    bool
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
				i:    false,
			},
			wantErr: false,
		},
		{
			name: "pass2",
			args: args{
				args: []string{"hello", "../../test.txt"},
				i:    false,
			},
			wantErr: false,
		},
		{
			name: "pass3",
			args: args{
				args: []string{"hello"},
				i:    false,
			},
			wantErr: false,
		},
		{
			name: "fail1",
			args: args{
				args: []string{"hello", "dir"},
				i:    false,
			},
			wantErr: true,
		},
		{
			name: "fail2",
			args: args{
				args: []string{"hello", "any", "any"},
				i:    false,
			},
			wantErr: true,
		},
		{
			name: "fail3",
			args: args{
				args: []string{""},
				i:    false,
			},
			wantErr: true,
		},
		{
			name: "case insensitive1",
			args: args{
				args: []string{"Hello", "../../dir"},
				i:    true,
			},
			wantErr: false,
		},
		{
			name: "case insensitive2",
			args: args{
				args: []string{"HeLLo", "../../test1.txt"},
				i:    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.args.args, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
