package main

import "fmt"

var (
	first = 989
)

func main() {

	x := 43

	if true {
		x := 65
		fmt.Println("x inside the inner block => ", x)
		// this will not be 43, inside block is separated from the outside block, or the file block
	}

	// this is the outer scope value, not related to the if statement inner scope
	fmt.Println("x inside the outer block => ", x)

	fmt.Println("'first' outer value is ==========> ", first) // this will print the outer first
	{
		first := 545
		fmt.Println("'first' inner value is =======> ", first) // this one will not be 989, this will print the inner scope first variable
	}
}
