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

	// we can also match multiple cases
	var str string = txt // this is now true, but with a shorter text it'll be false
	switch ln := len(str) > 10; ln {
	// here i set one rule for both conditions, true and false:
	case true, false:
		fmt.Printf("It doesnt matter what the len is => %v, %v", len(str), ln)
	default:
		fmt.Println("Surprisingly none of above!")
	}
}
