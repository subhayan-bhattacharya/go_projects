// using generics in go
// interfaces having type parameters
package main

import "fmt"

type Pair[T fmt.Stringer] struct {
	First  T
	Second T
}

// An interface which works with any type
// that implements the Stringer interface, has the String method
// that returns a string representation of the type
// and has a Diff method that takes a parameter of the same type
// and returns a float64 value.
type Different[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

// func which takes a parameter of type T
// and returns a Pair of type T
// the T is implementing the Different interface
// here we are using the Diff method of the Different interface
func FindCloser[T Different[T]](a, b Pair[T]) Pair[T] {
	d1 := a.First.Diff(a.Second)
	d2 := b.First.Diff(b.Second)
	if d1 < d2 {
		fmt.Println("a is closer")
		return a
	}
	fmt.Println("b is closer")
	return b
}

// Point struct implements the Different interface
type Point struct {
	X, Y int
}

func (p Point) String() string {
	fmt.Println("String method called")
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}
func (p Point) Diff(q Point) float64 {
	x := p.X - q.X
	y := p.Y - q.Y
	return float64(x*x + y*y)
}

func main() {
	p1 := Pair[Point]{Point{1, 2}, Point{3, 4}}
	p2 := Pair[Point]{Point{5, 6}, Point{7, 8}}
	check := FindCloser(p1, p2)
	fmt.Println(check)
	// output:
	// b is closer
	// String method called
	// String method called
	// {Point(5, 6) Point(7, 8)}

	/* When Go's fmt.Println(check) runs:

	It sees that check is a struct.

	It prints the struct in its default format: {field1 field2}

	Since both field1 and field2 are of a type (Point) that has a String() method, it uses that method. */
}
