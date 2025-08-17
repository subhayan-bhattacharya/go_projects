package main

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	testCases := []struct {
		testName  string
		sentence  string
		substring string
		expected  int
	}{
		{
			testName:  "First",
			sentence:  "abababa",
			substring: "aba",
			expected:  3,
		},
		{
			testName:  "Second",
			sentence:  "banana",
			substring: "ana",
			expected:  2,
		},
		{
			testName:  "Third",
			sentence:  "mississippi",
			substring: "issi",
			expected:  2,
		},
		{
			testName:  "Fourth",
			sentence:  "aaaaa",
			substring: "aa",
			expected:  4,
		},
		{
			testName:  "Fourth",
			sentence:  "hello world",
			substring: "l",
			expected:  3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := countSubstringsUsingBuiltin(tc.sentence, tc.substring)
			if result != tc.expected {
				t.Errorf("test failed for testcase %s, expected %d result %d", tc.testName, tc.expected, result)
			}

		})
	}
}
