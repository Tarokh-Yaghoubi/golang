// A method is a function associated with a receiver, usually a value or pointer
// of a particular type. Unlike a regular function, which is called by its name,
// a method is called on a receiver, such as account.Deposit().

// You can define methods for both custom or derived types.

// There are no overloaded methods in go, unline C++ or Java

package main

import "fmt"

type user struct {
	username string
	password string
	userId 	 int
	isActive bool
}

type Circle struct {
	radius float64
}

type Rectangle struct { 
	length float64
	width  float64
}

// This is the syntax of a method, slightly different with a normal function
// A method has a receiver
func (u * user) setUsername(newUserName string) {
	u.username = newUserName
}

// a method for the Circle struct
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// a method for the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.length * r.width
}


func main() {
	s := user {
		username: "John",
		password: "TheCakeISAFake##!@!",
		userId: 25,
		isActive: true,
	}
	fmt.Println("username before change => ", s.username)
	s.setUsername("Phillip")
	fmt.Println("username after change => ", s.username)

	rect := Rectangle {
		length: 20,
		width: 15,
	}

	var area float64 = rect.Area()
	fmt.Println("The area of the rectangle is => ", area)
}