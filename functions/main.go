// replace a substring without built in replace

package main

import (
	"fmt"
	"strings"
)

func replaceWholeWord(sentence string, toReplace string, newWord string) string {
	words := strings.Fields(sentence)
	for index, word := range words {
		if word == toReplace {
			words[index] = newWord
		}
	}
	return strings.Join(words, " ")
}

func replaceSubstrings(sentence string, toReplace string, newWord string) (string, error) {
	if len(toReplace) == 0 {
		return "", fmt.Errorf("Invalid input passed in for substring which needs to be replaced %s", toReplace)
	}
	if len(sentence) == 0 {
		return "", fmt.Errorf("The sentence is invalid %s", sentence)
	}
	var result strings.Builder
	for i := 0; i < len(sentence); i++ {
		// here in the below sentence is the trick
		if i+len(toReplace) <= len(sentence) {
			if sentence[i:i+len(toReplace)] == toReplace {
				result.WriteString(newWord)
				// below line is also very important
				i = i + len(toReplace) - 1
			} else {
				result.WriteByte(sentence[i])
			}
		} else {
			result.WriteByte(sentence[i])
		}

	}
	return result.String(), nil
}

func main() {
	sentence := "Subhayan is a great guy!"
	result := replaceWholeWord(sentence, "great", "okay")
	fmt.Println(result)
	new := "programming"
	check, error := replaceSubstrings(new, "gram", "XX")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(check)
	}

}
