package main

import "fmt"

// structs and interfaces in go

// i think that 'int' property is so cryptic
// I might not use that
type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	int
}

type owner struct {
	name string
}

// we can also initialize structs like this:
var foo = struct {
	name    string
	sirname string
	age     int
}{"Tarokh", "Yaghoubi", 22}

func testStruct() {
	var myEngine gasEngine
	fmt.Println(myEngine.gallons, myEngine.mpg)

	var otherEngine gasEngine = gasEngine{99, 99, owner{"Tarokh"}, 0}
	fmt.Println(otherEngine.gallons, otherEngine.mpg, otherEngine.ownerInfo.name, otherEngine.int)

	var anotherEngine gasEngine = gasEngine{gallons: 76, ownerInfo: owner{"John"}, int: 0}
	fmt.Println(anotherEngine.gallons, anotherEngine.mpg, anotherEngine.ownerInfo.name, anotherEngine.int)
}
