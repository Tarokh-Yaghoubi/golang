
package main 

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) {
	fmt.Println("camelcase => ")
	var result string;

	for i := 0; i < len(s); i++ {
		fmt.Printf("character => %c\n", s[i])
		if s[i] != '-' {
			result += string(s[i])
		} else {
			i++
			result += strings.ToUpper(string(s[i]))
		}
	}

	fmt.Println("string is now => ", result)
}

func main() {
	ToCamelCase("tarokh-IsAGood-Man")
} 