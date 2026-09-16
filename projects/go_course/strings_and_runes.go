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

	name := "tarokh"	// sequence of bytes
	// firstLetter := name[0]
	// fmt.Println(firstLetter)
	// lastLetter := name[len(name) - 1]
	// fmt.Println(lastLetter)

	// you can modify the string, but u cannot modify the bytes inside it directly
	firstLetter := string(name[0])
	fmt.Println(firstLetter)

	var b rune = 'x'
	fmt.Println(string(b))

	// convert a string to a byte/rune
	runeName := []rune(name)
	fmt.Println(runeName)


	solKey := "𝄞𝄞𝄞"
	unicodeString := []rune(solKey)	// first we will convert the string into a rune-array
	firstCharacter := unicodeString[0]	// we will fetch the first rune from the unicodeString
	fmt.Println(string(firstCharacter))	// we will print it out, so it will not be garbage
}

func main() {
	test()
}
