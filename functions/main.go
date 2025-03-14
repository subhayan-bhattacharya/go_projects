// function which accepts a struct as argument and then calculates the area and perimeter

package main

import (
	"fmt"
)

type Rectange struct {
	width  int
	height int
}

func findPermiter(rectange Rectange) int {
	return 2*rectange.height + 2*rectange.width
}

func findArea(rectangle Rectange) int {
	return rectangle.height * rectangle.width
}

func main() {
	r1 := Rectange{width: 10, height: 20}
	fmt.Println(findPermiter(r1))
	fmt.Println(findArea(r1))

}
