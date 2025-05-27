package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

// The method is called Double
// And this is a method on the DoubleInt type
func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

// for this we use a value receiver method
func (d DoubleIntSlice) Double() {
	for index, value := range d {
		d[index] = value * 2
	}
}

// interfaces are also comparable
// both their values and types need to be equal
func CompareDouble(d1, d2 Doubler) bool {
	return d1 == d2
}

func main() {
	d1 := DoubleInt(10)
	d2 := DoubleInt(10)
	if CompareDouble(&d1, &d2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not comparable")
	}
}
