// Self implementation of map and filter func in go using generics

package main

import (
	"fmt"
	"strings"
)

func mapDemo[T any](values []T, mapFunc func(T) T) {
	for index, value := range values {
		fmt.Println("Checking the value :", value)
		values[index] = mapFunc(value)
	}
}

func filterDemo[T any](values []T, filterFunc func(T) bool) []T {
	results := []T{}
	for _, value := range values {
		fmt.Println("Checking the value :", value)
		if filterFunc(value) {
			results = append(results, value)
		}
	}
	return results
}

func upperCaseString(value string) string {
	return strings.ToUpper(value)
}

func stringsWithMoreThan5Chars(value string) bool {
	return len(value) > 5
}

func main() {
	names := []string{"Subhayan", "Shaayan", "Dimpu", "Papa"}
	mapDemo(names, upperCaseString)
	fmt.Println(names)
	filteredNames := filterDemo(names, stringsWithMoreThan5Chars)
	fmt.Println(filteredNames)
}
