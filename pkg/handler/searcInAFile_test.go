package handler

import "testing"

func TestSearchInAFile(t *testing.T) {
	type args struct {
		filename string
		arg      string
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{
			name: "",
			args: args{
				filename: "../../test.txt",
				arg:      "hello",
			},

			wantErr: false,
		}, {
			name: "",
			args: args{
				filename: "../../test2.txt",
				arg:      "hello",
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SearchInAFile(tt.args.filename, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInAFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
