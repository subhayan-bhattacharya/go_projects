package main

import (
	"fmt"
	"snippets"
)

func main() {
	input := []string{"tea", "eat", "ate", "listen", "silent", "⌘⌥", "⌥⌘", "⌘⌘⌥", "café", "facé", "hello"}
	check := snippets.GroupAnagrams(input)
	for _, group := range check {
		fmt.Println(group)
	}
}
