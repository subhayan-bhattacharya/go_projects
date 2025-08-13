// testing anagrams in Go

package main

import (
	"fmt"
	"strings"
)

func checkAnagram(word1 string, word2 string) bool {
	word1Runes := []rune(strings.ToLower(word1))
	word2Runes := []rune(strings.ToLower(word2))
	charCount := map[rune]int{}
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1Runes); i++ {
		charCount[word1Runes[i]]++
		charCount[word2Runes[i]]--
	}

	for _, value := range charCount {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	if checkAnagram("silent", "listen") {
		fmt.Println("They are anagrams")
	} else {
		fmt.Println("They are not anagrams")
	}
}
