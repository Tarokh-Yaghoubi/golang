package main

import "fmt"


// strings are stored as sequences of bytes, 
// each character can represent 1, 2 or more bytes of data
// The slice expression works with strings
// string, rune, byte types can be converted to each other

// strings are immutable, you cannot change characters inside a string using slice expressions

func test() {
	var str string = "SOL Key 𝄞𝄞𝄞"
	fmt.Println("str is => ", str)

	var s string = "Code & Learn"
	fmt.Println(string(s[0]))
	fmt.Println(s[0])	// this will also work, but it will just print out the ASCII Numeric of the character, not its string representation
	fmt.Println(string(s[:5]))
	fmt.Println(string(s[3:]))
	fmt.Println(string(s[:]))
	// s[1] = 'a'	// this is a bug
}

func main() {
	test()
}
