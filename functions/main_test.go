package main

import (
	"testing"
)

func TestReplaceWord(t *testing.T) {
	testCases := []struct {
		name      string
		sentence  string
		toReplace string
		newWord   string
		expected  string
	}{
		{
			name:      "First",
			sentence:  "The quick brown fox",
			toReplace: "brown",
			newWord:   "white",
			expected:  "The quick white fox",
		},
		{
			name:      "Second",
			sentence:  "My name is Subhayan",
			toReplace: "Subhayan",
			newWord:   "Shaayan",
			expected:  "My name is Shaayan",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceWholeWord(tc.sentence, tc.toReplace, tc.newWord)
			if result != tc.expected {
				t.Errorf("replace entire word failed %s, expected %s result %s", tc.name, tc.expected, result)
			}
		})
	}
}

func TestReplaceSubstring(t *testing.T) {
	testCases := []struct {
		testName  string
		sentence  string
		toReplace string
		newWord   string
		expected  string
		hasError  bool
	}{
		{
			testName:  "First",
			sentence:  "programming",
			toReplace: "gram",
			newWord:   "XX",
			expected:  "proXXming",
			hasError:  false,
		},
		{
			testName:  "Second",
			sentence:  "cat dog",
			toReplace: "cat",
			newWord:   "elephant",
			expected:  "elephant dog",
			hasError:  false,
		},
		{
			testName:  "Third",
			sentence:  "abcabc",
			toReplace: "abc",
			newWord:   "X",
			expected:  "XX",
			hasError:  false,
		},
		{
			testName:  "Fourth",
			sentence:  "hello world",
			toReplace: "hello ",
			newWord:   "",
			expected:  "world",
			hasError:  false,
		},
		{
			testName:  "Fifth",
			sentence:  "aaaa",
			toReplace: "aa",
			newWord:   "X",
			expected:  "XX",
			hasError:  false,
		},
		{
			testName:  "Sixth",
			sentence:  "",
			toReplace: "aa",
			newWord:   "X",
			expected:  "XX",
			hasError:  true,
		},
		{
			testName:  "Seventh",
			sentence:  "great",
			toReplace: "",
			newWord:   "X",
			expected:  "grext",
			hasError:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := replaceSubstrings(tc.sentence, tc.toReplace, tc.newWord)
			if err != nil {
				if !tc.hasError {
					t.Error("Expected an error to be returned")
				}
			} else {
				if result != tc.expected {
					t.Errorf("replace substring failed for testcase %s, expected %s result %s", tc.testName, tc.expected, result)
				}
			}

		})
	}
}
