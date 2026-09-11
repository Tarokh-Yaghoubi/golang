package main

import "fmt"

// bool values can be constants
// numeric values can be constants
// functions like imag, cap, len, real can be constants
// strings and runes can be constants
// always-true conditions can be constants

// we have types and untyped constants

// constants can be declared in package-level and function-level scope
// constants can be declared in packs like variables
// constants cannot change unlike variables
// constants can remain unused! and go compiler will not throw an error for unused constants

const (
	first                  = 12  // untyped constant
	second           int32 = 434 // typed constant
	application_name       = "Go Course"
	isRunning              = true
	character              = 'a'
	isTrue                 = 1 == 1 // always-true conditions can be constants
)

// despite it is untyped, u cannot assign it to a non-numeric type like string or bool, it will throw an error

func test() {
	var temp int = first
	fmt.Println(temp)

	var floatNum float64 = float64(first)
	fmt.Println(floatNum)

	var secondTemp int32 = second
	fmt.Println(secondTemp)

	const z = complex(1, 43) // this is possible because u are using a constant function like complex
	const y = imag(z)        // this is possible because u are using a constant function like imag
	fmt.Printf("%v %v \n", z, y)

}

type Sum struct {
	first  int
	second int
}

func (s Sum) Add() int {
	return s.first + s.second
}

func main() {
	test()

	var sum Sum
	sum.first = 10
	sum.second = 20

	fmt.Println(sum.Add())
}
