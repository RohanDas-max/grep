package utils

import (
	"reflect"
	"testing"
)

func TestSearchCaseInsensitive(t *testing.T) {
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
			name: "pass stdin",
			args: args{
				data:  []string{"Hello world", "not welcome"},
				a:     "hello",
				path:  "",
				stdin: true,
			},
			want:  []string{},
			want1: "Hello world",
		},
		{
			name: "pass",
			args: args{
				data:  []string{"Hello world", "not welcome"},
				a:     "hello",
				path:  "",
				stdin: false,
			},
			want:  []string{"Hello world"},
			want1: "",
		},
		{
			name: "pass2",
			args: args{
				data:  []string{"Hello world", "not welcome"},
				a:     "Hello",
				path:  "",
				stdin: false,
			},
			want:  []string{"Hello world"},
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SearchCaseInsensitive(tt.args.data, tt.args.a, tt.args.path, tt.args.stdin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchCaseInsensitive() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SearchCaseInsensitive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
