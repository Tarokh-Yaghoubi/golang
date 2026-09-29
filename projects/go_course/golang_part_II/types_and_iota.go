
package main

import "fmt"

type Direction int

const (
	North Direction = iota	 // 3
	East	// 1
	South	// 2
	West	// 3
)

func main() {
	fmt.Println(North, East, South, West)
}