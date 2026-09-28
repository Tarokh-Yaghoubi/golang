
package main

import "fmt"


type Person struct { 
	name string
	age int
}

type Rectangle struct {
	length int
	width  int
}

func (r * Rectangle) SetLength(newLen int) {
	r.length = newLen
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d\n", p.name, p.age)
}

func showLater(show func()(string)) {
	fmt.Println("show => ", show())
}

func main() {
	first_person := &Person{name: "Tarokh", age: 22,}	// now this contains the address
	getter := first_person.GetDetails
	fmt.Println("getter data => ", getter())
	showLater(first_person.GetDetails)
	showLater(getter)


	// This is what we call a method expression
	// This will convert a method to a function, where the `receiver` becomes 
	// the first explicit parameter:

	f1 := Person.GetDetails
	p := Person{name: "Johny", age: 32,}
	fmt.Println("f1 data ===>", f1(p))

	// we can do the same thing with the RECTANGLE as well
	f2 := (*Rectangle).SetLength
	r := &Rectangle{length: 5, width: 3,} 
	f2(r, 1)
	fmt.Println("r => ", r.length)
}