
package main 

import "fmt"

func main() {
	var char rune = '4'
	fmt.Printf("a rune => %c\n", char)
	var otherChar int32 = 100
	fmt.Printf("an int32 => %c\n", otherChar)
	var anotherChar int32 = 'a'
	fmt.Printf("an int32 => %c\n", anotherChar)
	var mybyte byte = 'A'
	fmt.Printf("a byte => %c\n", mybyte)
	fmt.Println("the same byte => ", mybyte)

	var str string = `a moderate programmer trying to stay alive`
	// strings are hold as a byte array internally, and are immutable
	
}