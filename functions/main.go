// compare interfaces
package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i, value := range d {
		d[i] = value * 2
	}
}

func DoubleComparer(d1 Doubler, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {
	var i1 DoubleInt = 10
	var i2 DoubleInt = 10
	var slice1 = DoubleIntSlice{1, 2, 3}
	var slice2 = DoubleIntSlice{1, 2, 3}
	DoubleComparer(&i1, &i2)
	fmt.Println(slice1, slice2)
	DoubleComparer(slice1, slice2)
}
