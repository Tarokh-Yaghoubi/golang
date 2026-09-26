package main

import "fmt"

func dummy() {
	fmt.Println("===========================")
}

// First rule:
// defer line runs last, not when the code reaches that line
// Last line will be printed last, first is the First line
func firstRuleExample() {
	defer fmt.Println("Last line")
	fmt.Println("First Line")
}

// Second rule:
// arguments are evaluated when the defer is reached
// of course, defer line will be executed last in the function scope,
// but this does not mean that it is going to print the updated x val in the end
// so 10 is printed, not 9
func secRuleExample() {
	x := 10
	defer fmt.Println("x is => ", x)
	x = 9
	// 10 is printed
}

// Third rule:
// multiple defers will run in reverse order, like this:
func thirdRuleExample() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	// this will be => C, B, A
}

// Fourth rule:
// a defered function can change a named return value
func fourthRuleExample() (n int) {
	defer func() {
		n++
	}()

	return 41 // this will return 42, not 41
}

// main

func main() {
	firstRuleExample()
	dummy()
	secRuleExample()
	dummy()
	thirdRuleExample()
	dummy()
	fmt.Println("fourthRuleExample -> ", fourthRuleExample())
}
