package main

import (
	"fmt"
	"time"
)

func someFunction(num int) {
	fmt.Println(num)

}

func firstRun() {
	go someFunction(1)
	go someFunction(2)
	go someFunction(3)

	time.Sleep(3 * time.Second) // this will block the main thread for 3 seconds
	fmt.Println("hi")
}

// channels in golang

func secondRun() {
	myChannel := make(chan string)

	// here we will fork and create a new goroutine, and it will run the code inside the function
	go func() {
		time.Sleep(2 * time.Second)
		myChannel <- "data" // syntax to send data to a channel
	}()

	// the main function is blocked until it receives data from the channel
	msg := <-myChannel // syntax to receive data from a channel
	// now the main functions is receiving data from the channel.
	fmt.Println(msg)
	fmt.Println("finish")
}

func thirdRun() {
	myChannel := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		var result int = 23 + 43
		myChannel <- fmt.Sprintf("the result is => %d", result)
	}()

	msg := <-myChannel
	fmt.Println(msg)
	fmt.Println("finish third run")
}

func main() {
	// firstRun()
	// secondRun()
	thirdRun()
}
