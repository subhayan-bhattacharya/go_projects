// Implementation of a queue

package main

import (
	"fmt"
	"math"
)

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1 Pair[T], pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair1.Val2.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	} else {
		return pair2
	}
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.x, p.y)
}

func (p Point) Diff(p2 Point) float64 {
	x := p.x - p2.x
	y := p.y - p2.y
	// formula which uses Pythagorus theorm to
	// calculate distance between two points
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func main() {
	// each of these points satisfy the interface Differ
	p1 := Point{
		x: 10,
		y: 20,
	}
	p2 := Point{
		x: 30,
		y: 40,
	}
	p3 := Point{
		x: 5,
		y: 15,
	}
	p4 := Point{
		x: 11,
		y: 21,
	}
	pair1 := Pair[Point]{
		Val1: p1,
		Val2: p2,
	}
	pair2 := Pair[Point]{
		Val1: p3,
		Val2: p4,
	}

	fmt.Println(FindCloser(pair1, pair2))
}
