
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now => ", now)	

	twoHours := 2 * time.Hour
	futureTime := now.Add(twoHours)
	fmt.Println("two hours from now => ", futureTime)

	start := time.Now()
	a := 1
	for i := 0; i < 1000000; i++ {
		a++
	}
	elapsed := time.Since(start)
	fmt.Println("elapsed time => ", elapsed)
}
