// combine type terms into interfaces example

package main

import (
	"fmt"
)

type PrintableInt interface {
	~int
	String() string
}

type newType int

func (n newType) String() string {
	return fmt.Sprintf("Value : {%d}", n)
}

func PrintTwice[T PrintableInt](value T) {
	check := value.String()
	fmt.Println(value.String())
	fmt.Println(check)
	fmt.Println(value + 100)
}

func main() {
	var x newType = 420
	PrintTwice(x)
}
