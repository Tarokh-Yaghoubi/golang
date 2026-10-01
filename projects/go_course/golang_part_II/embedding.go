// Go is an object oriented language
// But it does not support inheritance as we have in C++ and Java
// It supports Compositions instead
// this file explains embedding in Compositions

package main

import "fmt"

type Engine struct {
	horsePower int
}

type Car struct {
	model string
	Engine
}

func (e *Engine) Start() {
	fmt.Printf("Engine starts with horsePower=%d\n", e.horsePower)
}

func (c *Car) Drive() {
	fmt.Printf("Driving a %s with horsePower %d\n", c.model, c.horsePower)
}

func main() {
	// here i can create an object from Engine, and Car
	// the Car object can also call Start()
	// the Engine object can only Start()

	var firstCar Car = Car{
		model:  "mercedes",
		Engine: Engine{horsePower: 515},
	}

	var v12Engine Engine = Engine{
		horsePower: 510,
	}

	firstCar.Start()
	firstCar.Drive()

	v12Engine.Start()
}
