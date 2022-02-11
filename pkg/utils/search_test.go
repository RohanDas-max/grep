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
		c     map[string]bool
		cf    map[string]int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 string
		want2 int
	}{
		{
			name: "case sensitive",
			args: args{
				data:  []string{"Hello world", "hEllo World"},
				a:     "hello",
				path:  "",
				stdin: false,
				c: map[string]bool{
					"i": false,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:  []string{},
			want1: "",
			want2: 0,
		},
		{
			name: "case Insensitive",
			args: args{
				data:  []string{"Hello world", "hEllo World"},
				a:     "hello",
				path:  "",
				stdin: false,
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:  []string{},
			want1: "",
			want2: 2,
		},
		{
			name: "case Insensitive2",
			args: args{
				data:  []string{"Hello world", "hEllo World"},
				a:     "hello",
				path:  "",
				stdin: false,
				c: map[string]bool{
					"i": true,
					"c": false,
				},
				cf: map[string]int{},
			},
			want:  []string{"Hello world", "hEllo World"},
			want1: "",
			want2: 0,
		},
		{
			name: "case Insensitive stdin",
			args: args{
				data:  []string{"Hello world", "hEllo World"},
				a:     "hello",
				path:  "",
				stdin: true,
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:  []string{},
			want1: "",
			want2: 1,
		},
		{
			name: "case Insensitive stdin",
			args: args{
				data:  []string{"Hello world", "hEllo World"},
				a:     "hello",
				path:  "",
				stdin: true,
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{},
			},
			want:  []string{},
			want1: "",
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Search(tt.args.data, tt.args.a, tt.args.path, tt.args.stdin, tt.args.c, tt.args.cf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Search() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
