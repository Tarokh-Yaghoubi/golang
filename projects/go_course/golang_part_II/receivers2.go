
package main

import "fmt"

type Calculator struct {

}

func (c * Calculator) Add(a, b int) int {
	return a + b
}

func AddFunc(a, b int) int {
	return a + b
}

func ArithmeticOperation(fn func(int, int)(int), first, second int) int {
	return fn(first, second)
}

func main() {
	c := Calculator{}
	fmt.Println(ArithmeticOperation(c.Add, 10, 10))
	fmt.Println(ArithmeticOperation(AddFunc, 10, 10))
}