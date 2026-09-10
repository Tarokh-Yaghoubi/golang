package main

import (
	"fmt"
	"strings"
)

func testStrings() {
	var myRune = 'a'
	fmt.Printf("my rune = %c\n", myRune)

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
}

func main() {
	testStrings();
}
