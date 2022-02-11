package controller

import (
	"reflect"
	"testing"
)

func TestController(t *testing.T) {
	type args struct {
		args []string
		c    map[string]bool
		cf   map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   int
		wantErr bool
	}{
		{
			name: "pass with c/i flag",
			args: args{
				args: []string{"Hello", "../../dir"},
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   5,
			wantErr: false,
		},
		{
			name: "pass with c flags",
			args: args{
				args: []string{"hello", "../../dir"},
				c: map[string]bool{
					"i": false,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   4,
			wantErr: false,
		},
		{
			name: "pass with i",
			args: args{
				args: []string{"hello", "../../dir"},
				c: map[string]bool{
					"i": true,
					"c": false,
				},
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
		{
			name: "pass without flag",
			args: args{
				args: []string{"hello", "../../dir"},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{},
			},
			want: []string{
				"../../dir/dir2/test.txt:this is a test file in dir2 hello world",
				"../../dir/test.txt:hello world",
				"../../dir/test.txt:hello world also from test.txt",
				"../../dir/test1.txt:hello world 1",
			},
			want1:   0,
			wantErr: false,
		},
		{
			name: "fails",
			args: args{
				args: []string{"hello", "dir"},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   0,
			wantErr: true,
		},
		{
			name: "pass with arg[0] only",
			args: args{
				args: []string{"hello"},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   0,
			wantErr: false,
		},
		{
			name: "fail with no arg",
			args: args{
				args: []string{},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   0,
			wantErr: true,
		},
		{
			name: "pass with cf flag",
			args: args{
				args: []string{},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{
					"A": 10,
					"B": 10,
					"c": 10,
				},
			},
			want:    []string{},
			want1:   0,
			wantErr: true,
		},
		{
			name: "fail with alot of arg",
			args: args{
				args: []string{"h", "e", "l", "l", "O", "world"},
				c: map[string]bool{
					"i": false,
					"c": false,
				},
				cf: map[string]int{
					"A": 10,
					"B": 10,
					"c": 10,
				},
			},
			want:    []string{},
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Controller(tt.args.args, tt.args.c, tt.args.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Controller() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
