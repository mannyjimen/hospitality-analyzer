package reviewfilter

import (
	"encoding/json"
	"testing"
)

func TestUnfairReview(t *testing.T) {
	keywords = map[string]struct{}{
		"invisible":     {},
		"uncomfortable": {},
	}

	tests := []struct {
		input    json.RawMessage
		expected bool
	}{
		{
			input:    json.RawMessage(`This business does not satisfy my needs, I felt invisible!`),
			expected: true},
		{
			input:    json.RawMessage(`I love this place, its my goto spot for boba!`),
			expected: false},
		{
			input:    json.RawMessage(`The waiter made us feel uncomfortable, really disappointed...`),
			expected: true},
		{
			input:    json.RawMessage(``),
			expected: false},
	}

	for _, tt := range tests {
		result := isUnfairReview(tt.input)
		if result != tt.expected {
			t.Errorf("Expected %t, got %t\nReview: %s", tt.expected, result, tt.input)
		}
	}
}
