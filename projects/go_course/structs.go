package main

import "fmt"

// while building real-world application, you probably need more data types than the basic
// types provided by golang
// GO provides "structs" to allow you to create your own custom data types
// A Struct helps group similar data types together

type Student struct {
	firstname string
	lastname  string
	age       int
	classes   []string
	isAlive   bool
}

// An empty struct is a struct that all of its values are having zero of their corresponding data type
// for example the zero value for a string is " ", but for an int is 0

// for the empty struct above you will see => [" ", " ", 0, [], false], ALL EMPTY

func test() {
	var student Student
	otherStudent := Student{} // another way to define an empty struct in golang
	fmt.Println("empty struct => ", student, " - ", otherStudent)

	noneEmptyStudent := Student{
		"Tarokh",
		"Yaghoubi",
		22,
		[]string{"Music Class", "German Class", "Boxing Class", "University Classes", "Learning Golang"},
		true,
	} // this is implicitly defined, so you are forced to set all the values in order

	// if I explicitly define a structure, I can skip some fields which are not needed, so golang will automatically set them to zero

	anotherNoneEmpty := Student{
		lastname:  "Smith",
		firstname: "John",
		isAlive:   false,
	} // all other fields will be zero, and because it is explicitly defined, the order we write the fields does not matter
	// it will put firstname at first itself

	fmt.Printf("noneEmptyStudent defined => %v\n", noneEmptyStudent)
	fmt.Printf("anotherNoneEmpty defined => %v\n", anotherNoneEmpty)

	// the DOT notation is used to interact with struct fields, reading and writing

	fmt.Println("firstname of the first noneEmptyStudent => ", noneEmptyStudent.firstname)
	noneEmptyStudent.firstname = "George St Piere"
	fmt.Println("firstname of the first noneEmptyStudent after CHANGE => ", noneEmptyStudent.firstname)
}

func main() {
	test()
}
