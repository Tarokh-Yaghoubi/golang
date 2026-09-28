
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

func (u * user) setUsername(newUserName string) {
	if u == nil {
		return 0
	}
	
	u.username = newUserName
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}


func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	
}