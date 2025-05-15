// using generics in go
// implement a stack using generics
package main

import "fmt"

type Stack[T any] struct {
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
	fmt.Println(stringStack.Pop())
}
