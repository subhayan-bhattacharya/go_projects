package main

import "fmt"

func removeDuplicates(word string) string {
	runes := []rune(word)
	var removed []rune
	seen := map[rune]bool{}
	for _, r := range runes {
		_, exists := seen[r]
		if exists {
			continue
		} else {
			removed = append(removed, r)
		}
		seen[r] = true
	}
	return string(removed)
}

func main() {
	check := removeDuplicates("programming")
	fmt.Println(check)
}
