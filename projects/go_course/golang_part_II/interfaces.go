// interfaces are types, that define the functionality
// of a type, they define what type should do
// But NOT HOW IT SHOULD DO IT

// Abstract types such as interfaces, only specify functionality
// Concrete types such as structs with methods, specify both
// functionality and implementation

// Interfaces will declare the method members which they contain
// And to implement an Interface, a type must implement all the
// methods in the interface

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Name() string
	Shape
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	name   string
	radius float64
}

type Square struct {
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Name() string {
	return c.name
}

func (s Square) Area() float64 {
	return s.width * 4
}

func printArea(s Shape) {
	fmt.Printf("%.2f\n", s.Area())
}

func printObj(o Object) {
	fmt.Printf("object name => %s, object area => %.2f\n", o.Name(), o.Area())
}

func main() {
	var circle Circle = Circle{name: "Gerdaloo", radius: 4}
	var rectangle Rectangle = Rectangle{height: 3, width: 5}

	var square Square = Square{width: 4}
	printArea(square) // no problem because square is implementing Area()

	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		printArea(shape)
	}

	printObj(circle)
}
