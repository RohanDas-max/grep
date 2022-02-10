package handler

import (
	"reflect"
	"testing"
)

func TestSearchInAFile(t *testing.T) {
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
			name: "",
			args: args{
				f: "../../test1.txt",
				a: "hello",
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   1,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				f: "../../test1.txt",
				a: "hello",
				c: map[string]bool{
					"i": false,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{"hello world"},
			want1:   1,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				f: "../../test.txt",
				a: "hello",
				c: map[string]bool{
					"i": true,
					"c": false,
				},
				cf: map[string]int{},
			},
			want: []string{"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"lorem ipsum hello world",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"hello world lorem ipsum",
				"Hello",
			},
			want1:   0,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				f: "../../test.txt",
				a: "hello",
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   19,
			wantErr: false,
		},
		{
			name: "fail1",
			args: args{
				f: "test.txt",
				a: "hello",
				c: map[string]bool{
					"i": true,
					"c": true,
				},
				cf: map[string]int{},
			},
			want:    []string{},
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := SearchInAFile(tt.args.f, tt.args.a, tt.args.c, tt.args.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInAFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInAFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SearchInAFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
