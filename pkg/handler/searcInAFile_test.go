package handler

import (
	"reflect"
	"testing"
)

func TestSearchInAFile(t *testing.T) {
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
				f: "../../test1.txt",
				a: "hello",
			},
			want:    []string{"hello world", ""},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				f: "test1.txt",
				a: "hello",
			},
			want:    []string{""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchInAFile(tt.args.f, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInAFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInAFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
