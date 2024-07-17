package helper

import (
	"testing"
)

func TestGenerateNextCode(t *testing.T) {
	tests := []struct {
		name     string
		lastCode string
		expected string
	}{
		{
			name:     "Empty last code",
			lastCode: "",
			expected: "0001",
		},
		{
			name:     "Valid last code",
			lastCode: "0001",
			expected: "0002",
		},
		{
			name:     "Valid last code with leading zeros",
			lastCode: "0099",
			expected: "0100",
		},
		{
			name:     "Valid last code with no leading zeros",
			lastCode: "1234",
			expected: "1235",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateNextCode(tt.lastCode)
			if result != tt.expected {
				t.Errorf("GenerateNextCode(%s) = %s; expected %s", tt.lastCode, result, tt.expected)
			}
		})
	}
}
