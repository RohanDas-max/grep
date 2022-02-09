package utils

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		data []string
		a    string
		path string
		res  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "pass",
			args: args{
				data: []string{"hello world", "not welcome"},
				a:    "hello",
				path: "",
				res:  []string{},
			},
			want: []string{"hello world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.data, tt.args.a, tt.args.path, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
