// find first non repeating character in a string

package main

import "fmt"

func findFirstNonRepeatingCharacter(s string) string {
	runes := []rune(s)
	counter := make(map[rune]int, len(runes))

	for _, r := range runes {
		_, ok := counter[r]
		if ok {
			counter[r]++
		} else {
			counter[r] = 1
		}
	}
	for _, r := range runes {
		if counter[r] == 1 {
			return string(r)
		}
	}
	return ""
}

func main() {
	check := findFirstNonRepeatingCharacter("swiss")
	fmt.Println(check)
}
