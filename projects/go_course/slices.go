
package main

import "fmt"

// slices are wrappers around arrays
// they manage the underlying arrays for you
// they help you overcode the limitations of arrays
// len is not attached to the type of the slice, u can pass the slices without paying attention to the len
// U cannot pass a slice to a func that takes exactly an array of type  [10]int 

// len of a slice is the len of the elements stored in the slice 
// cap is total size of the unbderlying array cap()

func test() {
	//var a [] int	// this is a slice, if u pass the len, it will be an array !
	//b := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}	// this is a slice
	
	// we can also use the make function
	c := make([]int, 0)	// a slice of len 0
	fmt.Printf("size of the slice => %d\n", len(c))
	// we can also pass the capacity of the slice as the third optional argument to make() func
	// d := make([]int, 5)
	e := make([]int, 5, 10)
	f := make([]int, 0, 10)

	fmt.Println(len(e), cap(e))
	fmt.Println(len(f), cap(f))

	// append is used to append elements to a slice
	
	// these are appended to the end of the slice 
	// because the slice 'e' is filled with zero-len in the beginning
	e = append(e, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("new e is => ", e)

	
	// SIMILARTIES WITH ARRAYS

	// like arrays u can fetch or manipulate values stored at given index
	// Both slices and arrays are zero-indexed 
	// like arrays, u cannot index -1, and u cannot index an address bigger than &arr[len(slice - 1)], it will overflow

	// NIL in go is different with NULL
	// nil means a lack of value and type
	// nil is thye zero value for slice (when there is nothing in the slice, then it is nil)
	// nil slice does not have underlying array
	// an EMPTY (NOT NILL) Slice has underlying array (with no values stored)
	var z [] int	// nil slice, z is a slice of type INT, no value in it
	h := []int{}	// empty slice, with underlying array, not NIL ! 
	fmt.Println("z NIL slice > ", len(z), cap(z))
	fmt.Println("h empty slice > ", len(h), cap(h))
	h = append(h, 100, 200, 300, 400, 500, 600, 700, 800, 900)
	fmt.Println("h after populating it with data => ", h, len(h), cap(h))

	h = append(h, e...)
	fmt.Println("h after appending e => ", h)

	
	// there is no problem with passing a slice of (len x or len y) to a function, the size is not assigned to the type like a normal array. 
	someFunction(h)		// this is ok
	someFunction([]int{1000, 2000, 3000, 4000, 5000})	// this is also ok 

}


func someFunction(slc [] int) {

}

func main() {
	test()
}