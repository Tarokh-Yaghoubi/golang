package main

import "fmt"

// This file demonstrates how structs are defined and used in Go.
// A struct is a custom data type that groups related fields together.
// In this example, we create a gas engine model with several properties.

// gasEngine represents a simple engine model.
// It contains a few fields that describe the engine state and ownership.
// The anonymous field `int` is embedded directly into the struct, which means
// its type name is used as the field name when accessed. This is a bit cryptic,
// so it is usually better to give the field a clearer name in real code.
type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	int
}

// owner stores a person's name.
// This is a nested struct used inside gasEngine to describe the owner.
type owner struct {
	name string
}

// This is an anonymous struct literal.
// It is created directly without a named type, and it is useful for quick,
// one-off data objects. Here it represents a person with a name and age.
type foo struct {
	name    string
	sirname string
	age     int
}

// fullName returns the person's full name by combining the first and last name fields; this relates to Go methods on structs (object behavior / encapsulated logic).
func (e foo) fullName() string {
	return e.name + " " + e.sirname
}

// testStruct shows multiple ways to create and print struct values.
// It demonstrates zero-value initialization, positional initialization,
// and named-field initialization in Go.
func testStruct() {
	// myEngine is declared without initialization, so all fields get their zero values.
	// For uint8, the zero value is 0.
	var myEngine gasEngine
	fmt.Println(myEngine.gallons, myEngine.mpg)

	// This creates a gasEngine using positional field order:
	// mpg, gallons, ownerInfo, int
	// This is valid, but it can be less readable if many fields exist.
	var otherEngine gasEngine = gasEngine{99, 99, owner{"Tarokh"}, 0}
	fmt.Println(otherEngine.gallons, otherEngine.mpg, otherEngine.ownerInfo.name, otherEngine.int)

	// This creates a gasEngine using named fields.
	// Named initialization is clearer and safer because it does not depend on field order.
	var anotherEngine gasEngine = gasEngine{gallons: 76, ownerInfo: owner{"John"}, int: 0}
	fmt.Println(anotherEngine.gallons, anotherEngine.mpg, anotherEngine.ownerInfo.name, anotherEngine.int)

	var myFoo foo = foo{name: "Tarokh", sirname: "Jacobi", age: 30}
	var fullname string = myFoo.fullName() // This will return "Tarokh Jacobi"
	fmt.Printf("Full name: %s\n", fullname)
}
