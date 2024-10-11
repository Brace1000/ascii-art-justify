package ascii_art

import (
	"testing"
)

func Test_getTerminalSize(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTerminalSize(); got != tt.want {
				t.Errorf("getTerminalSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
