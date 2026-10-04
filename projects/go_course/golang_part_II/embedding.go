
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