package main

import (
	"testing"
)

func TestCli(t *testing.T) {
	tests := []struct {
		name     string
		inputURL []string
		expected string
	}{
		{
			name:     "test over argument",
			inputURL: []string{"path", "https://website1.com", "https://overvalue.com", "https://overvalue.com", "1"},
			expected: "too many arguments provided",
		},
		{
			name:     "test no website provided ",
			inputURL: []string{"path"},
			expected: "no website provided",
		},
		{
			name:     "test with no max page",
			inputURL: []string{"path", "https://website1.com"},
			expected: "no max concurrency provided",
		},
		{
			name:     "test with no max page",
			inputURL: []string{"path", "https://website1.com", "1"},
			expected: "no max pages provided",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, err := cmd(tc.inputURL)
			if err.Error() != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, err)
			}
		})
	}
}
