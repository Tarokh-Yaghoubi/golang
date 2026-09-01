package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

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
	anotherChannel := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		var result int = 200 + 4343
		anotherChannel <- fmt.Sprintf("the anotherChannel result is => %d", result)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		var result int = 23 + 43
		myChannel <- fmt.Sprintf("the result is => %d", result)
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

	fmt.Println("finish third run")
}

func main() {
	// firstRun()
	// secondRun()
	thirdRun()
	// fourthRun()
	// funfFunction()
	// sixthRun()

}

func getDataFromChannel(channel chan string) *string {
	// this will get data from a channel, process the data and return its pointer
	msg := <- channel
	for i := 0; i < len(msg); i++ {
		fmt.Printf("my channel data => %c\n", msg[i])
	}

	return &msg
}

func getDataNormally(data string) string {
	var finalData string = fmt.Sprintf("my already packed data =======> %s", data)
	return finalData
}

func funfFunction() {
	var name string = "THis is another sentence"
	var data string
	go func() {
		fmt.Printf("Another string inside a goroutine\n")	
		data = getDataNormally(name)
	}()
	
	fmt.Printf("funfFunction data ======> %s\n", data)
}

func fourthRun() {
	myChannel := make(chan string)
	var name string = "This is my name"
	
	go func() {
		myChannel <- name
	}()

	var data * string = getDataFromChannel(myChannel)
	fmt.Printf("returned data from channel ==========> %v", *data)

}

func printDummyNumbers() {
	for i := 0; i< 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%dth line got printed\n", i)
		fmt.Printf("wg list ======> %v", wg)
	}
	wg.Done()
}

func sixthRun() {
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printDummyNumbers()
	}

	wg.Wait()
	
}