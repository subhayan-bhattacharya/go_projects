package main

import (
	"fmt"
	"strings"
)

func StringReverse(s string) []rune {
	runes := []rune(s)
	var reverse = []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, runes[i])
	}
	return reverse
}

func isPalindrome(sentence string) bool {
	reversed := StringReverse(sentence)
	reversedModified := strings.ReplaceAll(strings.ToLower(string(reversed)), " ", "")
	sentenceModified := strings.ReplaceAll(strings.ToLower(sentence), " ", "")
	return sentenceModified == string(reversedModified)
}

func main() {
	sentence := "was it a cat i saw"
	if isPalindrome(sentence) {
		fmt.Println("It is a palindrome!")
	} else {
		fmt.Println("It is not a palindrome")
	}
}
