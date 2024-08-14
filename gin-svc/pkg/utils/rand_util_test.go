package utils

import "testing"

func TestGenerateRandomNumber(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "TestGenerateRandomNumber",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomNumber(); got != tt.want {
				t.Errorf("GenerateRandomNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
