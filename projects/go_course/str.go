package main

import (
	"fmt"
	"strings"
)

// this is how we group multiple variables in golang

// this can be used for global variables, 
var (
	myInt int = 43
	myString string = "Hello Golang"
	myBool bool = false
	myFloat float64 = 3.14
)
// this can be defined inside a function as well

func testStrings() {
	var myRune = 'a'
	fmt.Printf("my rune = %c\n", myRune)
	var apartmentNumber, streetname = 54, "Main Street"
	apartmentNumber2, streetname2 := 54, "Main Street" // this is also possible 
	var strSlice = []string{"t", "a", "r", "o", "k", "h"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n my string is => %v\n", catStr)

	var str string 
	str = "hello golang"
	fmt.Printf("str is -> %s\n", str)

	var firstname, lastname string
	firstname = "Tarokh"
	lastname = "Yaghoubi"
	var fullname string 
	fullname = firstname + " " + lastname
	fmt.Printf("fullname is -> %s\n", fullname)

	// this is a better practice 
	// for concatenating strings in golang
	fmt.Printf("%s %s\n", firstname, lastname)	
	
	// or this:
	fullname = fmt.Sprintf("%s %s", firstname, lastname)
	fmt.Printf("fullname is -> %s\n", fullname)


	// if u wanna declare a variable, this is how to do:
	//var nickname string
	// then u can define it later:
	//nickname = "Craw"

	// but if u wanna define the variable, 
	// u can just omit the var keyword and use := operator:
	//otherName := "Craw"	// no var keyword
	

	// u can also use the old format:
	//var anothername string = "OtherName - Craw Again"
	//fmt.Printf("anothername is -> %s\n", anothername)

	// var age int
	// var age2 int = 22
	// var age3 = 22	// this is also possible
	// var age4, othername2 = 22, "Craw"	// this is also possible
	// age5 := 22	// this is also possible
	age6, othername3 := 22, "Craw"	// also this one :)
	fmt.Printf("age and other name values => %v %s\n", age6, othername3)
}

func main() {
	testStrings();
}
