package main

import (
	"errors"
	"fmt"
)

func main() {
	sayHiToSomeone("Tarokh")
	fullname := getFullName("John", "Smith")
	fmt.Println("Full name is => ", fullname)
	fullname2, len := getFullNameWithLen("John", "Smith")
	fmt.Printf("Full name is => %v and the length is => %v\n", fullname2, len)

	// we can use the underscore value to ignore the value we dont want to use
	fullname3, _ := getFullNameWithLen("John", "Smith")
	fmt.Printf("Full name is => %v and the length is ignored\n", fullname3)
	// this works completely fine
}

func test_functions() {

	// functions in golang
	fmt.Println("an external function")

	var mystr string = "tarokh is learning go"
	printMe(mystr)

	var res, remainder, err = intDivision(10, 0)
	switch {
	case err != nil:
		fmt.Printf("%v\n", err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is ==> %v", res)
	default:
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

func sayHiToSomeone(name string) {
	fmt.Printf("Hi %v, How are you?\n", name)
}

func getFullName(first, lastname string) string {
	var fullname string = fmt.Sprintf("%s %s", first, lastname)
	return fullname
}

func getFullNameWithLen(firstname, lastname string) (string, int) {
	var fullname string = fmt.Sprintf("%s %s", firstname, lastname)
	var len int = len(fullname)
	return fullname, len
}
