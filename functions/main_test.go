package main

import (
	"testing"
)

func TestAnagramChecker(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected bool
	}{
		{
			name:     "First",
			word1:    "silent",
			word2:    "listen",
			expected: true,
		},
		{
			name:     "Second",
			word1:    "a gentleman",
			word2:    "elegant man",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkAnagram(tc.word1, tc.word2)
			if result != tc.expected {
				t.Errorf("checkAnagram failed for testcase %s, expected %t", tc.name, tc.expected)
			}
		})
	}
}
