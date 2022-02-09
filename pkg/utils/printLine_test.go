package utils

import "testing"

func TestPrintLine(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				s: []string{"hello", "world"},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				s: []string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintLine(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("PrintLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
