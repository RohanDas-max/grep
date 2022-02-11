package handler

import "testing"

func TestSearchInStdin(t *testing.T) {
	type args struct {
		arg string
		c   map[string]bool
		cf  map[string]int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fail with any arg",
			args: args{
				arg: "",
				c:   map[string]bool{"i": true, "c": false},
				cf:  map[string]int{},
			},
			wantErr: true,
		},
		{
			name: "pass with i flag",
			args: args{
				arg: "hello",
				c:   map[string]bool{"i": true, "c": false},
				cf:  map[string]int{},
			},
			wantErr: false,
		},
		{
			name: "pass without flags",
			args: args{
				arg: "Hello",
				c:   map[string]bool{"i": false, "c": false},
				cf:  map[string]int{},
			},
			wantErr: false,
		},
		{
			name: "pass with both flags",
			args: args{
				arg: "Hello",
				c:   map[string]bool{"i": true, "c": true},
				cf:  map[string]int{},
			},
			wantErr: false,
		},
		{
			name: "pass with c flag",
			args: args{
				arg: "Hello",
				c:   map[string]bool{"i": false, "c": true},
				cf:  map[string]int{},
			},
			wantErr: false,
		},
		{
			name: "fails with wrong argument",
			args: args{
				arg: "",
				c:   map[string]bool{"i": false, "c": true},
				cf:  map[string]int{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SearchInStdin(tt.args.arg, tt.args.c, tt.args.cf); (err != nil) != tt.wantErr {
				t.Errorf("SearchInStdin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
