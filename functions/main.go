// demostration of how sort slice function is used to sort a slice of structs
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name     string
	lastName string
	age      int
}

func main() {
	people := []Person{
		{"Subhayan", "Bhattacharya", 41},
		{"Shaayan", "Bhattacharya", 37},
		{"Poulomi", "Adhikary", 37},
	}
	for _, person := range people {
		fmt.Println(person.lastName)
	}

	sort.Slice(people, func(i int, j int) bool {
		fmt.Println(i, j)
		return people[i].age < people[j].age
	})

	fmt.Println(people)
}
