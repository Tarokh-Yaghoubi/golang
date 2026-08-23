package main 

import "fmt"

func testPointers() {
	var p *int = new (int) // p is a pointer to an int, and it points to a newly allocated int variable
	var i int
	fmt.Println("i:", i) // prints 0, the default value of int
	fmt.Println("p:", *p)	// dereferencing p, prints 0, the default value of int
	i = 32
	*p = i

	fmt.Println("p after change => ",  *p)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice) // same value as sliceCopy :), because they are all pointers under the hood
	fmt.Println(sliceCopy)
}
