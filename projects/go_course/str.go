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
}
