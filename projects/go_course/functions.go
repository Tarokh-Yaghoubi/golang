package main

import (
	"errors"
	"fmt"
)

func test_functions() {

	// functions in golang
	fmt.Println("an external function")

	var mystr string = "tarokh is learning go"
	printMe(mystr)


	var res, remainder, err = intDivision(10, 0)
	switch {
		case err != nil
			fmt.Printf("%v\n", err.Error())
		case remainder == 0
			fmt.Printf("The result of the integer division is ==> %v", res)
		dafault:
			fmt.Printf("The result of the int is => %v and the remainder is %v\n", res, remainder)	
	}
}


func printMe(printval string) {
	fmt.Println("String is => " + printval)
}

func intDivision(numerator int, denominator int) (int, int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err

}