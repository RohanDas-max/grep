package controller

import "testing"

func TestController(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "pass",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
