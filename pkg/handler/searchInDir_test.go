package handler

import (
	"reflect"
	"testing"
)

func TestSearchInDir(t *testing.T) {
	type args struct {
		f  string
		a  string
		c  map[string]bool
		cf map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   int
		wantErr bool
	}{
		{
			name: "pass1",
			args: args{
				f:  "../../dir",
				a:  "hello",
				c:  map[string]bool{"i": true, "c": true},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   5,
			wantErr: false,
		}, {
			name: "pass2",
			args: args{
				f:  "../../dir",
				a:  "hello",
				c:  map[string]bool{"i": false, "c": true},
				cf: map[string]int{},
			},
			want: []string{
				"../../dir/dir2/test.txt:this is a test file in dir2 hello world",
				"../../dir/test.txt:hello world",
				"../../dir/test.txt:hello world also from test.txt",
				"../../dir/test1.txt:hello world 1",
			},
			want1:   4,
			wantErr: false,
		},
		{
			name: "pass2",
			args: args{
				f:  "../../dir",
				a:  "hello",
				c:  map[string]bool{"i": true, "c": false},
				cf: map[string]int{},
			},
			want: []string{
				"../../dir/dir2/test.txt:this is a test file in dir2 hello world",
				"../../dir/test.txt:hello world",
				"../../dir/test.txt:Hello world2",
				"../../dir/test.txt:hello world also from test.txt",
				"../../dir/test1.txt:hello world 1",
			},
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := SearchInDir(tt.args.f, tt.args.a, tt.args.c, tt.args.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInDir() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SearchInDir() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
