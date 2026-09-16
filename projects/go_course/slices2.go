
package main

import "fmt"
import "slices"

func test() {
	a := []int {1, 2, 3, 4, 5, 6}
	b := a[2:5]	// indexes 2, 3 and 4
	fmt.Println("slice b => ", b)
	
	c := a[:5] // index 0, till index 5 (five is not counted)
	fmt.Println("slice c => ", c)

	d := a[:]	// this will copy all the indexes, from beginning to end
	fmt.Println("slice d => ", d)

	e := a[3:]	// index 3, till the end
	fmt.Println("slice e => ", e)

	// ARRAYS also work like this
	// If u copy slice a, to slice e, you are also copying the address, 
	// which means that the underlying array is shared after that
	// so if u change/append a value in 'e', it will also change 'a'

	
	fmt.Println(a)
	fmt.Println(b)
	b = a[2:]
	b = append(b, 100)
	b[0] = 500
	
	fmt.Println("final a => ", a)
	fmt.Println("final b => ", b)

	// we do not specify the size of a slice when we declare it: 
	var slice1 = []string{"Tarokh", "Jacob", "German", "French"}	// this is a slice
	// using [...] makes an array, using [] makes a slice.
	fmt.Println("slice1 is => ", slice1)

	var slice2 [] string
	fmt.Println(slice2 == nil)	// an empty slice will be NIL - true
	// you cannot compare slices with eachother, slice1 == slice2 is wrong
	// the only thing u can compare slices with is NIL, slice1 == nil (completely fine)

	/*
		There are two functions for comparing two slices in golang
		1. slices.Equal(x, y)	-> takes two functions, returns true/false if the
								   elements are the same and the size is the same
								   requires the elements of a slice to be comparable

		2. slices.EqualFunc(x) -> lets you pass in a function to determine equality
		and does not require the slice elements to be comparable 								   
	
	*/

	fmt.Println(slices.Equal(slice1, slice2))	// will be false

	// NOTE: 
	/*
		'reflect' package contains a function called DeepEqual() that can compare 
		almost anything, including slices

		now, with the presence of 'slices' package, it is less safe and also slower to
		use DeepEqual()
	*/
	fmt.Println(len(slice1))	// len also works on slices, 
	// len will return 0 if the slice is == nil

	// cap is also used for slices, it is less frequent than len() for other types
	// cap always returns the same value as len() for arrays
	// cap is used to check if a slice is large enough to hold new data
	// cap is used to check if a call to make is needed to create a new slice


	// Despite it is nice that slices grow automatically, it is far more efficient 
	// to size them one, we can do it with the make() func

}

func main() {
	test()
}