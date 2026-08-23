package main

import "fmt"

func testPointers() {
	var p *int = new(int) // p is a pointer to an int, and it points to a newly allocated int variable
	var i int
	fmt.Println("i:", i)  // prints 0, the default value of int
	fmt.Println("p:", *p) // dereferencing p, prints 0, the default value of int
	i = 32
	*p = i

	fmt.Println("p after change => ", *p)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice) // same value as sliceCopy :), because they are all pointers under the hood
	fmt.Println(sliceCopy)

	fmt.Println("--------------------------")

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thin1 array is %p\n", &thing1)
	var result *[5]float64 = square(&thing1)
	fmt.Println("The result is => ", *result)

	fmt.Printf("\nthing1 after it is changed inside the square function => %v\n", thing1)
}

// square takes a pointer to a 5-element float array and updates each value to its square.
func square(thing2 *[5]float64) *[5]float64 {
	fmt.Printf("\nThe memory location of thin2 array is %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
