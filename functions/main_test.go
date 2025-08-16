package main

import (
	"testing"
)

func TestReplaceSubstring(t *testing.T) {
	testCases := []struct {
		testName string
		sentence string
		expected string
	}{
		{
			testName: "First",
			sentence: "leetcode",
			expected: "l",
		},
		{
			testName: "Second",
			sentence: "aabbcc",
			expected: "",
		},
		{
			testName: "Third",
			sentence: "a",
			expected: "a",
		},
		{
			testName: "Fourth",
			sentence: "åäöåä",
			expected: "ö",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := findFirstNonRepeatingCharacter(tc.sentence)
			if result != tc.expected {
				t.Errorf("test failed for testcase %s, expected %s result %s", tc.testName, tc.expected, result)
			}

		})
	}
}
