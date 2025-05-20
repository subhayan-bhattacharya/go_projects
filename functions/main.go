// using generics in go
// interfaces having type parameters
package main

import "fmt"

// The tilda (~) in the type constraint allows for type sets that include
// types that are defined by the user. In this case, it allows for any type
// that is defined as an alias of int, int8, int16, int32, or int64.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// The Convert function takes a value of type T1, which must satisfy the
// Integer interface, and returns a value of type T2, which also must
// satisfy the Integer interface. The function converts the value from
// type T1 to type T2 using a type conversion.
func Convert[T1, T2 Integer](value T1) T2 {
	return T2(value)
}

func main() {
	check := Convert[int, int8](10)
	fmt.Println(check) // 10
}
