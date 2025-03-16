// returning closures from functions
package main

import (
	"fmt"
)

func multiplyBybase(base int) func(int) int {
	return func(value int) int {
		return base * value
	}
}

func main() {
	multiplyFunc := multiplyBybase(20)
	for i := 0; i <= 100; i = i + 10 {
		fmt.Println("Value for i is: ", i)
		fmt.Println(multiplyFunc(i))
	}
}
