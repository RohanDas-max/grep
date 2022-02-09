package controller

import (
	"reflect"
	"testing"
)

func TestController(t *testing.T) {
	type args struct {
		args  []string
		iFlag bool
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
				args:  []string{"Hello", "../../test1.txt"},
				iFlag: true,
			},
			want:    []string{"hello world"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Controller(tt.args.args, tt.args.iFlag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}
