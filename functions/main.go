// In Go, whether a type implements an interface depends on the receiver types of its methods:

// If a method has a value receiver (e.g. func (e RealEngine) Start()), both RealEngine and *RealEngine can use it.

// If a method has a pointer receiver (e.g. func (e *RealEngine) Stop()), only *RealEngine can use it.

package main

type Engine interface {
	// Start the engine
	Start() string
	// Stop the engine
	Stop() string
}

type RealEngine struct {
	// Engine name
	Name string
	// Engine power
	Power int
}

func (e *RealEngine) Start() string {
	return "Engine started"
}
func (e *RealEngine) Stop() string {
	return "Engine stopped"
}

type Car struct {
	// Car name
	Name string
	// Car engine
	Engine Engine
}

func (c *Car) Start() string {
	return c.Engine.Start()
}
func (c *Car) Stop() string {
	return c.Engine.Stop()
}

func NewCar(name string, engine Engine) Car {
	return Car{
		Name:   name,
		Engine: engine,
	}
}

func main() {
	// Create a new engine
	engine := RealEngine{
		Name:  "V8",
		Power: 500,
	}

	// Create a new car
	// only *RealEngine implements the Engine interface, not RealEngine (the value type).
	car := NewCar("Mustang", &engine)

	// Start the car
	println(car.Start())

	// Stop the car
	println(car.Stop())
}
