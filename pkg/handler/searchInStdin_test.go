package handler

import "testing"

func TestSearchInStdin(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				arg: "hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SearchInStdin(tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("SearchInStdin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
