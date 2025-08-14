package main

import (
	"testing"
)

func TestLongestWord(t *testing.T) {
	testCases := []struct {
		name     string
		sentence string
		expected string
	}{
		{
			name:     "First",
			sentence: "The quick brown fox",
			expected: "quick",
		},
		{
			name:     "Second",
			sentence: "My name is Subhayan",
			expected: "Subhayan",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getLongestWord(tc.sentence)
			if result != tc.expected {
				t.Errorf("checkAnagram failed for testcase %s, expected %s", tc.name, tc.expected)
			}
		})
	}
}
