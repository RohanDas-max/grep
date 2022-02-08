package utils

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "negative",
			args: args{
				filename: "test.txt",
			},
			want:    []string{""},
			wantErr: true,
		}, {
			name: "positive",
			args: args{
				filename: "../../test1.txt",
			},
			want:    []string{"hello world", "not welcome"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
