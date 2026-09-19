<<<<<<< HEAD

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

	
=======
package main

import (
	"fmt"
	"slices"
)

func test() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:5] // indexes 2, 3 and 4
	fmt.Println("slice b => ", b)

	c := a[:5] // index 0, till index 5 (five is not counted)
	fmt.Println("slice c => ", c)

	d := a[:] // this will copy all the indexes, from beginning to end
	fmt.Println("slice d => ", d)

	e := a[3:] // index 3, till the end
	fmt.Println("slice e => ", e)

	// ARRAYS also work like this
	// If u copy slice a, to slice e, you are also copying the address,
	// which means that the underlying array is shared after that
	// so if u change/append a value in 'e', it will also change 'a'

>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
	fmt.Println(a)
	fmt.Println(b)
	b = a[2:]
	b = append(b, 100)
	b[0] = 500
<<<<<<< HEAD
	
	fmt.Println("final a => ", a)
	fmt.Println("final b => ", b)

	// we do not specify the size of a slice when we declare it: 
	var slice1 = []string{"Tarokh", "Jacob", "German", "French"}	// this is a slice
	// using [...] makes an array, using [] makes a slice.
	fmt.Println("slice1 is => ", slice1)

	var slice2 [] string
	fmt.Println(slice2 == nil)	// an empty slice will be NIL - true
=======

	fmt.Println("final a => ", a)
	fmt.Println("final b => ", b)

	// we do not specify the size of a slice when we declare it:
	var slice1 = []string{"Tarokh", "Jacob", "German", "French"} // this is a slice
	// using [...] makes an array, using [] makes a slice.
	fmt.Println("slice1 is => ", slice1)

	var slice2 []string
	fmt.Println(slice2 == nil) // an empty slice will be NIL - true
>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
	// you cannot compare slices with eachother, slice1 == slice2 is wrong
	// the only thing u can compare slices with is NIL, slice1 == nil (completely fine)

	/*
		There are two functions for comparing two slices in golang
		1. slices.Equal(x, y)	-> takes two functions, returns true/false if the
								   elements are the same and the size is the same
								   requires the elements of a slice to be comparable

		2. slices.EqualFunc(x) -> lets you pass in a function to determine equality
<<<<<<< HEAD
		and does not require the slice elements to be comparable 								   
	
	*/

	fmt.Println(slices.Equal(slice1, slice2))	// will be false

	// NOTE: 
	/*
		'reflect' package contains a function called DeepEqual() that can compare 
=======
		and does not require the slice elements to be comparable

	*/

	fmt.Println(slices.Equal(slice1, slice2)) // will be false

	// NOTE:
	/*
		'reflect' package contains a function called DeepEqual() that can compare
>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
		almost anything, including slices

		now, with the presence of 'slices' package, it is less safe and also slower to
		use DeepEqual()
	*/
<<<<<<< HEAD
	fmt.Println(len(slice1))	// len also works on slices, 
=======
	fmt.Println(len(slice1)) // len also works on slices,
>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
	// len will return 0 if the slice is == nil

	// cap is also used for slices, it is less frequent than len() for other types
	// cap always returns the same value as len() for arrays
	// cap is used to check if a slice is large enough to hold new data
	// cap is used to check if a call to make is needed to create a new slice

<<<<<<< HEAD

	// Despite it is nice that slices grow automatically, it is far more efficient 
	// to size them one, we can do it with the make() func

=======
	// Despite it is nice that slices grow automatically, it is far more efficient
	// to size them one, we can do it with the make() func

	firstSlice := []int{1, 2, 3, 4}
	var secondSlice = make([]int, 4)
	var ret int = copy(secondSlice, firstSlice)
	fmt.Printf("firsSlice=%v\nsecondSlice=%v\nretValue=%d\n", firstSlice, secondSlice, ret)

	// You dont have to copy the whole slice, you can only copy a part of that
	fmt.Println("==============================================")

	anotherFirstSlice := []int{10, 20, 30, 40}
	var anotherSecSlice = make([]int, 2) // this only holds two values, len is important in copy, not capacity!
	var thirdSlice = make([]int, 2)
	var retval int = copy(anotherSecSlice, anotherFirstSlice) // [10, 20], retval = 2
	fmt.Printf("anotherFirsSlice=%v\nanotherSecSlice=%v\nretval=%d\n", anotherFirstSlice, anotherSecSlice, retval)

	fmt.Println("==============================================")
	retval = copy(thirdSlice, anotherFirstSlice[2:])             // this will copy from the middle of the slice
	fmt.Printf("thirdSlice=%v\nretval=%d\n", thirdSlice, retval) // you will see [30, 40]

>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
}

func main() {
	test()
<<<<<<< HEAD
}
=======
}
>>>>>>> 39566c82be7df5c6639f63e510e88461bd0b0fcf
