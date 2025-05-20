// using generics in go
// interfaces having type parameters
package main

type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Generic function to divide and mod two integers
// T is a type parameter that must satisfy the Integer interface
// This means that T can be any type that is an int, int8, int16, int32, or int64
func divAndMod[T Integer](a, b T) (T, T) {
	return a / b, a % b
}

func main() {
	// Example usage of divAndMod function
	// The function will work with any integer type
	a, b := 10, 3
	quotient, remainder := divAndMod(a, b)
	println("Quotient:", quotient)
	println("Remainder:", remainder)

	// You can also use other integer types
	var a8 int8 = 10
	var b8 int8 = 3
	quotient8, remainder8 := divAndMod(a8, b8)
	println("Quotient:", quotient8)
	println("Remainder:", remainder8)
}
