package main

import "fmt"

type Worker struct {
	username string
	password string
	age      int
	tasks    []string
}

func test() {
	// for loops in golang
	// the only loops statement we have in golang is the for loop
	// there is no while, do-while, foreach, separate loop in golang
	// all of these are accessible using the forloop itself

	// 1. for with a condition
	// 2. for with a for clause (most popular in golang)
	// 3. for with range clause

	// for with a condition example:
	a := 10
	for a > 0 {
		fmt.Println("a is => ", a)
		a--
	}
	fmt.Println("DONE!")

	// if it was an infinite loop we could manage it like this:
	b := 10

	// this one loops similar to a while (1) in C
	for {
		fmt.Println("b is => ", b)
		b--

		// if we do not write this statement below, the loop will continue forever
		if b == 0 {
			break
		}
	}

	// for loop with a clause example:
	// different parts of a for-with-clause loop => {initialization, condition, post statement}
	fmt.Println("for loop with clause in golang => ")
	for x := 0; x < 10; x++ {
		fmt.Println("x is => ", x)
		// this works exactly like a for loop in C
	}

	// for range loop example:
	var myMap map[string]string = map[string]string{
		"username": "Tarokh",
		"password": "Yaghoubi",
		"color":    "black",
	}

	for index, value := range myMap {
		fmt.Printf("index=%v\tvalue=%v\n", index, value)
		if index == "password" {
			myMap["password"] = "Jesus"
		}
	}

	fmt.Println("myMap password val after change => ", myMap["password"])
}

func main() {
	test()
}
