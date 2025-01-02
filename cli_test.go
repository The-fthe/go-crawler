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
			inputURL: []string{"path", "https://website1.com", "https://overvalue.com"},
			expected: "too many arguments provided",
		},
		{
			name:     "test no website provided ",
			inputURL: []string{"path"},
			expected: "no website provided",
		},
		{
			name:     "test with website provided",
			inputURL: []string{"path", "https://website1.com"},
			expected: "starting crawl of: https://website1.com",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := cmd(tc.inputURL)
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}
