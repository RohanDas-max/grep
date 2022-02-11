package utils

import "testing"

func Test_count(t *testing.T) {
	type args struct {
		reS []string
		s   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{
				reS: []string{""},
				s:   "hello",
			},
			want: 1,
		},
		{
			name: "pass",
			args: args{
				reS: []string{"Hello world", "Hello world"},
				s:   "",
			},
			want: 2,
		},
		{
			name: "pass",
			args: args{
				reS: []string{""},
				s:   "Hello",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.reS, tt.args.s); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
