package handler

import (
	"reflect"
	"testing"
)

func TestSearchInDir(t *testing.T) {
	type args struct {
		f string
		a string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				f: "../../dir",
				a: "hello",
			},
			want:    []string{"../../dir/dir2/test.txt:this is a test file in dir2 hello world", "../../dir/test.txt:hello world", "../../dir/test1.txt:hello world 1"},
			wantErr: false,
		}, {
			name: "fail",
			args: args{
				f: "dir",
				a: "hello",
			},
			want:    []string{""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchInDir(tt.args.f, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
