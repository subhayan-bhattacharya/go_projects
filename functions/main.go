// The marriage between generic data structures and generic functions
// this time we use maps
// Keys have to be comparable in Go , this is a Go rule

package main

import "fmt"

func addTo(base int, values ...int) []int {
	results := []int{}
	for _, value := range values {
		results = append(results, base+value)
	}
	return results
}

func main() {
	check := addTo(100, 23, 45, 67, 90)
	fmt.Println(check)
}
