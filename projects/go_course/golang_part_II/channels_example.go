//  way to enable goroutines to pass around information
// 	Channels have some sort of features like:
//	1. They can hold data
//	2. They are thread safe
//	3. You can listen for data until it arrives and enters a channel

// and we can block code execution until one of these happens

package main

import "fmt"

func main() {
	// var c = make(chan int) // this can only hold a single int value
	// c <- 1 // * CODE WILL BLOCK HERE FOREVER AND WAITS FOR DATA TO COME

	// after the line below, the channel will be empty and 'i' will hold the
	// value
	// var i = <-c

	// ! Channels are not meant to be used like this
	// ! We will use the proper way channel + goroutines below
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i) // code execution will be stopped here until process() returns
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
