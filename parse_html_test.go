package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	inputs := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html>
                    <body>
                      <a href="/path/one">
                        <span>Boot.dev</span>
                      </a>
                      <a href="https://other.com/path/one">
                        <span>Boot.dev</span>
                      </a>
                    </body>
                  </html>`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs with params",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html>
                    <body>
                      <a href="/path/one?param1=test">
                        <span>Boot.dev</span>
                      </a>
                      <a href="https://other.com/path/one?param1=test&param2=test">
                        <span>Boot.dev</span>
                      </a>
                    </body>
                  </html>`,
			expected: []string{"https://blog.boot.dev/path/one?param1=test", "https://other.com/path/one?param1=test&param2=test"},
		},
		{
			name:     "multiple relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html>
                    <body>
                      <a href="/path/one">
                        <span>Boot.dev</span>
                      </a>
                      <a href="/path/two">
                        <span>Boot.dev</span>
                      </a>
                      <a href="/path/three">
                        <span>Boot.dev</span>
                      </a>
                      <a href="/path/four">
                        <span>Boot.dev</span>
                      </a>
                      <a href="/path/five">
                        <span>Boot.dev</span>
                      </a>
                      <a href="/path/13456">
                        <span>Boot.dev</span>
                      </a>
                    </body>
                  </html>`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two", "https://blog.boot.dev/path/three", "https://blog.boot.dev/path/four", "https://blog.boot.dev/path/five", "https://blog.boot.dev/path/13456"},
		},
	}
	for i, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(input.inputBody, input.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: unexpected error: %v", i, input.name, err)
			}
			if !reflect.DeepEqual(actual, input.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, input.name, input.expected, actual)
			}
		})
	}
}
