package reusablemethods

import "testing"

func TestIsMatched(t *testing.T) {
	type args struct {
		text string
		arg  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "matched",
			args: args{
				text: "helloworld",
				arg:  "hello",
			},
			want: "helloworld",
		}, {
			name: "matched",
			args: args{
				text: "helloworld",
				arg:  "blahh",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatched(tt.args.text, tt.args.arg); got != tt.want {
				t.Errorf("IsMatched() = %v, want %v", got, tt.want)
			}
		})
	}
}
