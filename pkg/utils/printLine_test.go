package utils

import "testing"

func TestPrintLine(t *testing.T) {
	type args struct {
		s      []string
		signal int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				s:      []string{"Hello world", "hello world 2"},
				signal: 0,
			},
		}, {
			name: "",
			args: args{
				s:      []string{"Hello world", "hello world 2"},
				signal: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintLine(tt.args.s, tt.args.signal)
		})
	}
}
