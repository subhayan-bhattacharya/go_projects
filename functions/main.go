// getting the longest word in a sentence

package main

import (
	"fmt"
	"strings"
)

func getLongestWord(sentence string) string {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return ""
	}
	longestWord := words[0]
	longest := len(words[0])
	for _, word := range words {
		if len(word) > longest {
			longest = len(word)
			longestWord = word
		}
	}
	return longestWord
}

func main() {
	sentence := "The quick brown fox"
	longest := getLongestWord(sentence)
	fmt.Printf("The longest word is %s", longest)
}
