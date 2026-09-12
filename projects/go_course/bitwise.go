

package main 

import "fmt"

func test() {
	 // OPERATOR PRECEDENCE
	 // PEMDAS => (parantheses, exponents, multiplication, division, addition, subtraction)
	a := 10
	b := 32
	c := 43.54
	sum := a + b
	fmt.Println("Sum => ", sum)

	firstname := "Tarokh"
	lastname := "yaghoubi"
	fmt.Printf("%s %s\n", firstname, lastname)

	difference := a - b
	fmt.Println("difference is => ", difference)


	complex1 := complex(10, 15)
	complex2 := complex(3, 4)
	complexSum := complex1 + complex2
	fmt.Println("full complex sum val => ", complexSum)
	fmt.Println("complex sum is => ", imag(complexSum))
	fmt.Println("complex sum real => ", real(complexSum))
	newSum := a + int(c)
	fmt.Printf("new sum is => %d\n", newSum)
}

func bitwise() {
	var flags uint8
	 // var fields uint8 = 0b00000100	 // another way to set the fields 


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
	
	flags = flags & ~(1 << 2)	// this will flip the bit, if will do the shift, then it will NOT the value, then it will do the AND
	fmt.Printf("flip the flag => [%08b]\n", flags)

	// cryptography algorithms, system programming, network programming, device drviers, etc.

	
}

func main() {

	bitwise()
	test()
}
