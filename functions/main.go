// using generics in go
// filter function that takes a slice of any type and a function that returns a boolean
// and returns a slice of the same type
package main

import "fmt"

func filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	evens := filter(ints, isEven)
	fmt.Println(evens) // [2 4 6]
}
