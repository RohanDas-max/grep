package handler

import (
	"reflect"
	"testing"
)

func TestSearchInDir(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				f: "test.txt",
			},
			want:    []string{""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchInDir(tt.args.f)
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
