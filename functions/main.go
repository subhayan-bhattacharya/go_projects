// using generics in go
// Map function which can accept any slice and apply a function to each element
// and return a new slice with the results
// The function transforms each element of the slice into another type

package main

import (
	"fmt"
	"strconv"
)

// Signature of the Map function accepts a slice of type T and a
// function that takes an element of type T and returns an element of type U
func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := []U{}
	for _, v := range slice {
		result = append(result, fn(v))
	}
	return result
}

// Example function to convert int to string
// This function will be passed to the Map function
// It takes an int and returns a string
func convertIntToString(i int) string {
	return strconv.Itoa(i) // Convert int to string
}

func CountLetters(s string) int {
	// count the number of bytes in the string
	return len(s)
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	strings := Map(ints, convertIntToString)
	fmt.Println(strings) // Output: [1 2 3 4 5]
	new_strings := []string{"hello", "world", "golang"}
	lengths := Map(new_strings, CountLetters)
	fmt.Println(lengths)
}
