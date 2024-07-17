package helper

import (
	"testing"
)

func TestGenerateNextCode(t *testing.T) {
	tests := []struct {
		name     string
		lastCode string
		expected string
		err      bool
	}{
		{
			name:     "Empty last code",
			lastCode: "",
			expected: "0001",
			err:      false,
		},
		{
			name:     "Valid last code",
			lastCode: "0001",
			expected: "0002",
			err:      false,
		},
		{
			name:     "Invalid last code",
			lastCode: "Z",
			expected: "",
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateNextCode(tt.lastCode)
			if (result == "" && !tt.err) || (result != "" && tt.err) {
				t.Errorf("GenerateNextCode(%s) returned an unexpected error state; got %s, expected error: %v", tt.lastCode, result, tt.err)
			}
			if result != tt.expected {
				t.Errorf("GenerateNextCode(%s) = %s; expected %s", tt.lastCode, result, tt.expected)
			}
		})
	}
}
