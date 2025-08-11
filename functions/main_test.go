package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var testCases = struct {
		sentences    []string
		isPalindrome []bool
	}{
		sentences: []string{
			"Taco cat",
			"Mom",
			"Murder for a jar of red rum",
			"Go deliver a dare vile dog",
		},
		isPalindrome: []bool{
			true,
			true,
			true,
			true,
		},
	}
	for index, sentence := range testCases.sentences {
		result := isPalindrome(sentence)
		expected := testCases.isPalindrome[index]
		if result != expected {
			t.Errorf("The sentence %s is not a palindrome", sentence)
		}
	}
}
