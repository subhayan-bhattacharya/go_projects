// using generics in go
// implement a stack using generics
package main

import "fmt"

type Stack[T comparable] struct {
	// stack is a slice of type T
	stack []T
}

// Push adds an element to the stack
func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

// Pop removes an element from the stack
func (s *Stack[T]) Pop() T {
	if len(s.stack) == 0 {
		panic("stack is empty")
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}

// check if the stack has certain element
func (s *Stack[T]) Contains(value T) bool {
	for _, v := range s.stack {
		if v == value {
			return true
		}
	}
	return false
}
func main() {
	// create a stack of integers
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Pop()) // 3
	fmt.Println(intStack.Pop()) // 2
	fmt.Println(intStack.Pop()) // 1

	// create a stack of strings
	stringStack := Stack[string]{}
	stringStack.Push("Subhayan")
	stringStack.Push("Bhattacharya")
	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.Pop())
	stringStack.Push("Shaayan")
	// fmt.Println(stringStack.Pop())
	if stringStack.Contains("Shaayan") {
		fmt.Println("Subhayan is in the stack")
	} else {
		fmt.Println("Subhayan is not in the stack")
	}
}
