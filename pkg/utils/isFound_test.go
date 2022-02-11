package utils

import "testing"

func TestIsFound(t *testing.T) {
	type args struct {
		txt string
		a   string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "pass1",
			args: args{
				txt: "Hello World",
				a:   "hello",
			},
			want:  "",
			want1: false,
		},
		{
			name: "pass2",
			args: args{
				txt: "Hello World",
				a:   "Hello",
			},
			want:  "Hello World",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isFound(tt.args.txt, tt.args.a)
			if got != tt.want {
				t.Errorf("IsFound() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsFound() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
