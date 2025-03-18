// understanding call by value
// when a function is called a copies of the params are made
package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastName string
	age      int
}

func modifySomething(p Person) {
	p.lastName = "Adhikary"
}

func main() {
	p := Person{name: "Subhayan", lastName: "Bhattacharya", age: 40}
	fmt.Println(p)
	modifySomething(p)
	fmt.Println(p)
}
