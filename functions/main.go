// type assertion in Go
package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      int
}

func main() {
	var x interface{} = Person{
		name:     "Subhayan",
		lastName: "Bhattacharya",
		age:      40,
	}
	x2 := x
	fmt.Println(x2.(Person))
}
