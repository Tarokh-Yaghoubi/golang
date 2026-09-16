
package main

import "fmt"

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

}

func main() {
	test()
}