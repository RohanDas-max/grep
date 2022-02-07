package handler

import "testing"

func TestSearchInAFile(t *testing.T) {
	type args struct {
		filename string
		arg      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				filename: "../../test.txt",
				arg:      "hello",
			},
			want:    "hello world asdasdasd adhashdjashd asjdh ajshda jsdjash dhsa jhas jasj Hello world",
			wantErr: false,
		}, {
			name: "pass",
			args: args{
				filename: "test1.txt",
				arg:      "hello",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchInAFile(tt.args.filename, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchInAFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SearchInAFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
