
package main

import (
	"fmt"
	"time"
)

func main() {
	charChannel := make(chan string, 3)	// buffered channel
	keyboard := make(chan string, 3)	// another buffered channel
	// in buffered channels, the sending goroutin will not block on the channel, it will write the data to the channel and it will continue what it was doing
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
			fmt.Println("char channel => ")
		case keyboard <- s:
			fmt.Println("keyboard channel => ")
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}


func main2() {
	go func() {
		for {
			select {
				default:
					fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}