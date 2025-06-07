// The marriage between generic data structures and generic functions
// generic Stack and functions to operate on that stack
// We need to use an external Go package here to make sure that we can compare
// two elements of a stack. The default comparable only does equal and not equal comparison

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Stack[T constraints.Ordered] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}
	element := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return element, true
}

func (s *Stack[T]) IsThere(element T) bool {
	for _, value := range s.values {
		if value == element {
			return true
		}
	}
	return false
}

func FindMaxInStack[T constraints.Ordered](s Stack[T]) (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}
	maxValue := s.values[0]
	for _, value := range s.values {
		if value > maxValue {
			maxValue = value
		}
	}
	fmt.Println("The max value about to be returned is :", maxValue)
	return maxValue, true
}

func main() {
	s := Stack[int]{}
	fmt.Println(s)
	s.Push(1)
	s.Push(56)
	s.Push(78)
	s.Push(81)
	fmt.Println(s)
	for {
		value, isNotEmpty := s.Pop()
		if isNotEmpty {
			fmt.Println(value)
		} else {
			fmt.Println("Is empty")
			break
		}

	}
	maxValue, isEmpty := FindMaxInStack(s)
	if isEmpty {
		fmt.Println("The max value in the stack is :", maxValue)
	} else {
		fmt.Println("The stack is empty!!")
	}

}
