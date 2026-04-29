package main

import "fmt"

type Circle struct {
	radius float64
	area float64
}

func (c *Circle) CalculateArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	circle := Circle{
		radius: 10,
	}
	circle.CalculateArea()
	fmt.Printf("%+v", circle)
}