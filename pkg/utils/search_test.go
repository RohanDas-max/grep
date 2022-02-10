package utils

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		data  []string
		a     string
		path  string
		stdin bool
	}
	x, _ := ReadFile("../../test1.txt")
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 string
	}{
		{
			name: "pass",
			args: args{
				data:  x,
				a:     "hello",
				path:  "",
				stdin: false,
			},
			want:  []string{"hello world"},
			want1: "",
		},
		{
			name: "pass1",
			args: args{
				data:  []string{"nothello"},
				a:     "hello",
				path:  "",
				stdin: true,
			},
			want:  []string{},
			want1: "nothello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Search(tt.args.data, tt.args.a, tt.args.path, tt.args.stdin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
