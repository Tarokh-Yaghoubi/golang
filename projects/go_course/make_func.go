
package main

import (
	"fmt"
)


func test() {
	x := make([]string, 5)	// a slice of size 5 which holds strings
	x = append(x, "first")	// beginners mistake
	fmt.Printf("x is => %01v\n", x)


	// you can specify the capacity with len zero,
	// NOTE: cap must never be less than len, logically that is not 
	// possible

	y := make([]string, 0, 5)
	y = append(y, "tarokh", "tommy", "cplusplus", "golang")
	fmt.Printf("y is => %v\ncap is => %v\n", y, cap(y))

	// there is a func in golang v1.21.0 that clears all the elements inside 
	// a slice, but the len() will remain unchanged, that func is called clear()

	clear(y)
	fmt.Printf("y is => %v\ncap is => %v\nlen is => %v\n", y, cap(y), len(y))
	// len and cap are unchanged, it just emptied the slice with " ", which is 
	// the zero value for the string type

	var data []int	// a way to declare slices => this is nil right now
	var data1 = []int{}	// this is not nil anymore, this is just an empty slice
	// which is not NIL. 

	fmt.Println(nil == data)	// this is true 
	fmt.Println(nil == data1)	// this is false, data1 is not NIL, it is just empty

	otherData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	otherData[0] = 2342
	fmt.Println(otherData)

}

func main() {
	test()
}
