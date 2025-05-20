// using generics in go
// interfaces having type parameters
package main

// The tilda (~) in the type constraint allows for type sets that include
// types that are defined by the user. In this case, it allows for any type
// that is defined as an alias of int, int8, int16, int32, or int64.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Generic function to divide and mod two integers
// T is a type parameter that must satisfy the Integer interface
// This means that T can be any type that is an int, int8, int16, int32, or int64
func divAndMod[T Integer](a, b T) (T, T) {
	return a / b, a % b
}

func main() {
	// Using the generic function with different integer types
	// int
	type MyInt int
	quotient, remainder := divAndMod[MyInt](10, 3)
	println("Quotient:", quotient)
	println("Remainder:", remainder)
}
