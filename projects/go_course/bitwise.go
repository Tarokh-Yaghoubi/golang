

package main 

import "fmt"

func main() {
	var flags uint8
	 // var fields uint8 = 0b00000100	 // another way to set the fields 

	 // OPERATOR PRECEDENCE
	 // PEMDAS => (parantheses, exponents, multiplication, division, addition, subtraction)

	fmt.Printf("%08b\n", flags)

	flags = flags | (1 << 2)
	fmt.Printf("%08b\n", flags)

	isOn := flags & (1 << 2) != 0 
	if isOn {
		fmt.Printf("flag is turned on => [%08b]\n", flags)
	}

	flags = flags | (1 << 3)
	flags = flags | (1 << 4)
	fmt.Printf("more flags are on now => [%08b]\n", flags)
	
	flags = flags & (0 << 2)
	fmt.Printf("flip the flag => [%08b]\n", flags)
}