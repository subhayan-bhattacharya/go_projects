// understanding call by value
// for a map it can change attribute values in the function
package main

import (
	"fmt"
)

func modifyMap(person map[string]string) {
	person["name"] = "Shaayan"
}

func modifySlice(numbers []int) {
	for index, number := range numbers {
		numbers[index] = number * number
	}
	numbers = append(numbers, 420)
}

func main() {
	person := map[string]string{
		"name":     "Subhayan",
		"LastName": "Bhattacharya",
		"age":      "40",
	}
	fmt.Println(person)
	modifyMap(person)
	fmt.Println(person)
	numbers := []int{2, 4, 6, 8}
	modifySlice(numbers)
	fmt.Println(numbers)
}
