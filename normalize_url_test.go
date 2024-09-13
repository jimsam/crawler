package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev",
		},
		{
			name:     "remove scheme",
			inputURL: "http://www.boot.dev/path",
			expected: "www.boot.dev",
		},
		{
			name:     "remove scheme",
			inputURL: "http://www.boot.dev/path?param1=test",
			expected: "www.boot.dev",
		},
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path?param1=test&param2=test",
			expected: "blog.boot.dev",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
