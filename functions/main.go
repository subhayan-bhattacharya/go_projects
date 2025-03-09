// function which gives even values from a slice of numbers

package main

import "fmt"

func produceEven(values []int) []int {
	results := []int{}
	for _, value := range values {
		if value%2 == 0 {
			results = append(results, value)
		}
	}
	return results
}

func main() {
	values := []int{20, 54, 67, 11, 3, 10}
	check := produceEven(values)
	fmt.Println(check)
}
