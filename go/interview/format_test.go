package interview

import "testing"

func TestFormatTestString(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FormatTestString()
		})
	}
}
