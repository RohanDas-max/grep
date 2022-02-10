package utils

import (
	"reflect"
	"testing"
)

func TestSearchCaseSensitive(t *testing.T) {

	type args struct {
		data  []string
		a     string
		path  string
		stdin bool
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 string
	}{
		{
			name: "Empty",
			args: args{
				data:  []string{"Hello World", "not welcome"},
				a:     "hello",
				path:  "",
				stdin: false,
			},
			want:  []string{},
			want1: "",
		},
		{
			name: "NotEmpty",
			args: args{
				data:  []string{"Hello World", "not welcome"},
				a:     "Hello",
				path:  "",
				stdin: false,
			},
			want:  []string{"Hello World"},
			want1: "",
		},
		{
			name: "stdin",
			args: args{
				data:  []string{"Hello World", "not welcome"},
				a:     "Hello",
				path:  "",
				stdin: true,
			},
			want:  []string{},
			want1: "Hello World",
		},
		{
			name: "stdin Case sensitive",
			args: args{
				data:  []string{"Hello World", "not welcome"},
				a:     "hello",
				path:  "",
				stdin: true,
			},
			want:  []string{},
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SearchCaseSensitive(tt.args.data, tt.args.a, tt.args.path, tt.args.stdin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchCaseSensitive() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SearchCaseSensitive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
