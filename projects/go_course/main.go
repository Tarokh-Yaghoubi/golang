package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello, world!")

	var intNum int32 = 423
	fmt.Println(intNum)

	var floatNumber float64 = 432.543
	fmt.Println(floatNumber)


	var floatNum32 float32 = 11.2
	var intNum32 int32 = 4
	var result float32 = float32(intNum32) + floatNum32
	fmt.Println(result)

	var mystring string = "Tarokh"
	var otherString string = "Multipleline\nstring\n"
	var anotherString string = `
		This is a multiline string 
		this is the second line 
		and this is the last line
	`

	var concatString string = "First" + " " + "Second" + " " + "Last"

	fmt.Println(mystring)
	fmt.Println(otherString)
	fmt.Println(anotherString)
	fmt.Println(concatString)

	fmt.Println(len("µ"))	// in go, len returns the number of bytes not the number of characters, so this will return 2
	// not 1


	// "RuneCountInString" function in unicode/utf8 package will return the number of characters, and not the bytes
	fmt.Println(utf8.RuneCountInString("µµµµ"))	// 4
	

	var myRune rune = 't'
	fmt.Println(myRune)	// this will return the ascii code (dec) - 116

	var myCondition bool = false	// booleans in golang
	var yourCondition bool = true
	
	fmt.Println(myCondition)
	fmt.Println(yourCondition)

	var myInteger int	// automatically initializes to zero for INTEGERS
	fmt.Println(myInteger)

	// these are other ways to define a variable in golang
	variable := "tarokh"
	fmt.Println(variable)

	var variable2 = "jacob"
	fmt.Println(variable2)

	first, second := 25, 26
	fmt.Println(first)
	fmt.Println(second)
	
	// this is a const, u cannot change it later
	// you cannot keep it uninitialized
	const myconst string  = "this is a const value"
	fmt.Println(myconst)

	test_functions()
}

