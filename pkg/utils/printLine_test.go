package utils

import "testing"

func TestPrintLine(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				data: []string{"hello world", "not welcome"},
			},
		}, {
			name: "1",
			args: args{
				data: []string{""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintLine(tt.args.data)
		})
	}
}
