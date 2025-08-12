package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "repeated letters",
			input:    "programming",
			expected: "progamin",
		},
		{
			name:     "consecutive duplicates",
			input:    "aabbcc",
			expected: "abc",
		},
		{
			name:     "case sensitive",
			input:    "AaBbCc",
			expected: "AaBbCc",
		},
		{
			name:     "numbers",
			input:    "1122334455",
			expected: "12345",
		},
		{
			name:     "special characters",
			input:    "!!@@##",
			expected: "!@#",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "no duplicates",
			input:    "abcdef",
			expected: "abcdef",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeDuplicates(tc.input)
			if result != tc.expected {
				t.Errorf("removeDuplicates(%q) = %q, expected %q",
					tc.input, result, tc.expected)
			}
		})
	}
}
