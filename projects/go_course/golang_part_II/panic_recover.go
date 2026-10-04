
// a panic occurs when GO program has reached a state of no return
// example: trying to access an element which is out of bounds, causes runtime panic

// there is a function called recover, that can be used to recover from panics
// recover() needs to be called in a deferred function

// The return value of calling recover() is the argument provided to panic. 

// Purpose of recover:
/*
	- to gracefully shutdown your code
	- it is not designed to act like exception handling in other programming languages
	- if you recover and continue execution,  the program will most probably panic again
	  and then the panic may go unnotice

*/

package main

import "fmt"

func panicExample() {
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("recovered from panic ", r)
		}
	}()

	panic("Something went wrong")
}

// func func1() {
// 	defer func(){
// 		fmt.Println("function1 deferred function called")
// 	}()
// 	func2()
// }


// func func2() {
// 	defer func(){
// 		fmt.Println("function2 deferred function called")
// 	}()

// 	panic("function2 panic")
// }

func main() {
	// func1()

	fmt.Println("Start Func Execution")
	panicExample()
	fmt.Println("Stop Func Execution")
}