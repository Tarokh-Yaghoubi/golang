
package main

import "fmt"

func main() {
	s := []int{1, 2}

	fmt.Println(s)

	someSlice(s)	// slices are implemented as reference types internally

	fmt.Println("slice after passing to the function => ", s, len(s))
}

func someSlice(s []int) {
	s[0] = 100 // this will change the original slice

	s = append(s, 1000)	// this will not change the original slice
	fmt.Println("len of slice in someSlice() => ", len(s))
	// because a copy of the slice is made inside this func, and the copy will not 
	// change the len property of the original slice outside of this function
}