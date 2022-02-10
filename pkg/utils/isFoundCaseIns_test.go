package utils

import "testing"

func Test_isFoundCaseIns(t *testing.T) {
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
			name: "1",
			args: args{
				txt: "Hello World",
				a:   "hello",
			},
			want:  "Hello World",
			want1: true,
		}, {
			name: "1",
			args: args{
				txt: "hello World",
				a:   "Hello",
			},
			want:  "hello World",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isFoundCaseIns(tt.args.txt, tt.args.a)
			if got != tt.want {
				t.Errorf("isFoundCaseIns() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isFoundCaseIns() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
