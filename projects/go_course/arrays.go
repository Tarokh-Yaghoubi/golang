
package main

import "fmt"


func test() {
	// arrays in golang

	// arrays are data structures used to store data in contiguos blocks of memory
	// each piece of data of an array is called the element of the array
	// each element of the arrays needs to be of the same type (they must all be runes or ints or strings. ...)
	// we use the index of the element in the array to access the element
	// arrays are zero indexed (they start from zero)

	// Advantages of arrays:
	// - ARRAYS ARE VERY EFFICIENT TO FETCH ELEMENTS FROM MEMORY
	// - They are used to store collection of data, such a list of numbers, etc. 
	// - They are used in search and sort algorithms, buffers, etc.

	var str [10] string	// this is how we can declare an array 
	var a [10] int32
	var b = [5] int {1, 2, 3, 4, 5}
	var c = [5] int {5, 2:10, 50}
	var d = [...]int{10, 20, 30, 40, 50, 60, 70, 80}	// go will automatically set the array size based on the number of elements u write

	// len function is used to get the len/size of the array

	// U cannot change the size of the array after declaration
	// The size or length of the array are stored in the type of the array
	// THIS MEANS THAT: In GO, [10]int and [5]int, are as unrelated as [4]string and [5]bool ... 
	// You cannot assign c = a as I've defined above! 
	// FUNCTION PARAMS must match the exact length 
	/*
	
		func sum(arr [5]int) {
			total := 0
			for _, v := range arr {
				total += v
			}

			return total
		}

		sum(b)	// this is just FINE 
		sum(d)	// this is COMPILER ERROR, d is not [5]int!, it does not match the parameter at all.
	*/

	fmt.Println(str)
	fmt.Println(d)
	fmt.Println(a)	// it is initialised with zeros
	fmt.Println(b)
	fmt.Println(c)	// specified values are set, others are zero

	z := [...] string {"tarokh", "Jacob", "Phillip"}
	fmt.Println(len(z), cap(z))		// 3, 3
	var value * int = &d[4]	// returning the address of an index like this 
	fmt.Printf("value is => %d\n", *value)	// printing out what lives inside that address

	// two-dimensional array in golang:
	twodim := [2][2]int{{1, 2},{2, 1}}	// in two-dimensional arrays we are forced to specify the len
	fmt.Println(twodim)
}

func main() {
	test()
}