// count substring occurances

package main

import (
	"fmt"
	"strings"
)

func countSubstringsUsingBuiltin(input string, substring string) int {
	result := 0
	if len(input) == 0 {
		return result
	}
	if len(substring) == 0 {
		return result
	}
	start := 0
	for {
		foundIndex := strings.Index(input[start:], substring)
		if foundIndex == -1 {
			return result
		} else {
			result++
			// this is clever
			start = start + foundIndex + 1
			if start >= len(input)-1 {
				break
			}
		}
	}
	return result
}

func main() {
	check := countSubstringsUsingBuiltin("hello world", "l")
	fmt.Println(check)
}
