
package main

import "fmt"

type Engine struct {
	model      string
	horsePower int
}

func (e * Engine) Start() {
	fmt.Println("starting the engine with horsePower => ", e.horsePower)
}

type Car struct {
	model string
	Engine
}

func (c * Car) Drive() {
	fmt.Printf("Driving a %s with horsePower => %d\n", c.Engine.model, c.horsePower)
}

func main() {
	// Go is an OBJECT-ORIENTED language, but it does not have the concept
	// of INHERITANCE

	// Instead, it uses COMPOSITION and PROMOTION to promote code reuse
	mercedes := Car{
		model: "cls63",
		Engine: Engine{model: "e200", horsePower: 518},
	} 

	fmt.Println(mercedes.model)
	fmt.Println(mercedes.horsePower)

	var engine Engine = Engine{horsePower: 520,}
	mercedes.Drive()
	engine.Start()
	mercedes.Start()	// Mercedes can also START(), because it has inherited ENGINE, and START
						// is a method for ENGINE. 
}
// Go is an object oriented language
// But it does not support inheritance as we have in C++ and Java
// It supports Compositions instead
// this file explains embedding in Compositions

// Compositions are used for reusing the code in golang

package main

import "fmt"

type GPS struct {
	model string
}

type Engine struct {
	model      string
	horsePower int
}

type Car struct {
	model string
	Engine
	GPS
}

func (e *Engine) Start() {
	fmt.Printf("Engine starts with horsePower=%d\n", e.horsePower)
}

func (c *Car) Drive() {
	fmt.Printf("Driving a %s with horsePower %d\n", c.Engine.model, c.horsePower)
	fmt.Printf("Driving a => %s\n", c.GPS.model)
	fmt.Printf("Driving a => %s\n", c.model)
}

func main() {
	// here i can create an object from Engine, and Car
	// the Car object can also call Start()
	// the Engine object can only Start()

	var firstCar Car = Car{
		model:  "mercedes",
		Engine: Engine{model: "e300", horsePower: 515},
		GPS:    GPS{model: "G-wagon"},
	}

	var v12Engine Engine = Engine{
		horsePower: 510,
	}

	firstCar.Start()
	firstCar.Drive()
	v12Engine.Start()
}
