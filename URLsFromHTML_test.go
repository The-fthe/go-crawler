package main

import (
	"fmt"
	"testing"
)

func TestURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
  <html>
      <body>
        <a href="/path/one">
          <span>Boot.dev</span>
        </a>
        <a href="https://other.com/path/one"> <span>Boot.dev</span>
        </a>
      </body>
  </html>
  `,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
  <html>
      <body>
        <a href="/path/one">
          <span>Boot.dev</span>
        </a>
        <a href="/path/one"> <span>Boot.dev</span>
        </a>
      </body>
  </html>
  `,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/one"},
		},
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
  <html>
      <body>
        <a href="https://other.com/path/one">
          <span>Boot.dev</span>
        </a>
        <a href="https://other.com/path/one"> <span>Boot.dev</span>
        </a>
      </body>
  </html>
  `,
			expected: []string{"https://other.com/path/one", "https://other.com/path/one"},
		},
		{
			name:     "one URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
  <html>
      <body>
        <a href="https://other.com/path/one"> <span>Boot.dev</span>
        </a>
      </body>
  </html>
  `,
			expected: []string{"https://other.com/path/one"},
		},
		{
			name:     "zero URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
  <html>
      <body>
      </body>
  </html>
  `,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actuals, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			for _, link := range actuals {
				fmt.Println("actual link: ", link)
			}
			if len(tc.expected) != len(actuals) {
				t.Errorf("Test %v -FAIL: Different length: expected: %d, get: %d", tc.name, len(tc.expected), len(actuals))
			}
			for i := range tc.expected {
				if actuals[i] != tc.expected[i] {
					t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected[i], actuals[i])
				}
			}
		})
	}
}
