// function which takes a slice of strings and gives back how many times they appear
// added code to also break a string into word slices

package main

import (
	"fmt"
)

func breakAStringIntoWords(value string) []string {
	results := []string{}
	temp := ""
	for _, character := range value {
		if character == ' ' {
			results = append(results, temp)
			temp = ""

		} else {
			temp += string(character)
		}
	}
	if len(temp) > 0 {
		results = append(results, temp)
	}
	return results
}

func findHowManytimes(values ...string) map[string]int {
	results := map[string]int{}
	for _, value := range values {
		_, ok := results[value]
		if ok {
			results[value] += 1
		} else {
			results[value] = 1
		}

	}
	return results
}

func main() {
	sentence := "Subhayan is a great guy. Subhayan has a brother callewd Shaayan . Subhayan also has a wife"
	broken := breakAStringIntoWords(sentence)
	final := findHowManytimes(broken...)
	fmt.Println(final)
}
