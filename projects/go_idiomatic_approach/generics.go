
package main

import "fmt"


func divAndRemainder(num, demon int) (int, int, error) {
	if demon == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / demon, num % demon, nil
}

type Node struct {
	val int
	Next *Node
}

func main() {

}