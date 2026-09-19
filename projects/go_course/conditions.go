package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}
func main() {
	var isNumEven bool = isEven(5)
	fmt.Println(isNumEven)

	// a new if statement we have in golang (not in C)
	if even := isEven(10); even {
		fmt.Printf("even value is => %v", even)
	} else {
		fmt.Printf("even value is not true => %v", even)
	}
}
