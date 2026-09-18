package main

import "fmt"

func main() {
	var txt string = "Tarokh is a normal programmer, nothing more or less"
	fmt.Println("len of txt is => ", len(txt))

	// this is the new way to write a switch statement in golang which does not exist in C programming (ITS NEW LAD!!!)

	switch ln := len(txt) > 10; ln {
	case true:
		fmt.Println("len of text is bigger than 10 => ", len(txt))
		break
	case false:
		fmt.Println("len of text is smaller than 10 => ", len(txt))
		break

	default:
		fmt.Println("Surprisingly none of above!")
	}
}
