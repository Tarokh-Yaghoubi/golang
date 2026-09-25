package main

import "fmt"

// In go you can assign a function to a variable and then you can use it later
// Functions can be passed to other functions

// we can use this signature instead of using func(int) int directly inside the code, so this makes the code more readable
// and also easier to maintain

type multiplier func(int) int

func main() {
	fmt.Println("Hello, World!")

	returnHeightAndAgeSum := func(age, height int) int {
		return age + height
	}

	result := foo("John", "Smith", returnHeightAndAgeSum)
	fmt.Println(result)

	crendentials := func(user, pass string) string {
		return fmt.Sprintf("username is => %s and password is => %s\n", user, pass)
	}

	otherFoo("John", "%TheCAKE$$isAFAKE$$0834!!", crendentials)

	res := apply(multi, 10)
	fmt.Printf("res => %v\n", res)

	// calling arithmeticOperations function:
	addFunc, _ := arithmeticOperations("add")
	if addFunc != nil {
		result := addFunc(1, 5)
		fmt.Printf("result of addFunc => %v\n", result)
	} else {
		fmt.Println("addFunc is nil")
	}

	// this is the way to catch a func returning from another func
	var multi func(int) int
	multi = multiplyBy(3)
	result2 := multi(10) // this results in 30
	fmt.Printf("result of multiplier => %v\n", result2)
}

func foo(firstname, lastname string, fn func(age, height int) int) string {
	fullname := fmt.Sprintf("%s %s", firstname, lastname)
	// Here i passed two random numbers to the function, but we can also pass what we are taking from the user in foo() params, to the function.

	sumOfAgeAndHeight := fn(25, 175)
	finalResult := fmt.Sprintf("fullname is => %s and the sum of age and height is => %d", fullname, sumOfAgeAndHeight)
	return finalResult
}

func otherFoo(username, password string, credentials func(user, pass string) string) {
	res := credentials(username, password)
	fmt.Printf("res => %v\n", res)
}

func multi(x int) int {
	return x * 3
}

func apply(fn func(int) int, value int) int {
	return fn(value)
}

// functions can return other functions

func arithmeticOperations(op string) (func(int, int) int, error) {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}, nil
	case "subtract":
		return func(a, b int) int {
			return a - b
		}, nil
	default:
		return nil, fmt.Errorf("unknown operation => %s", op)
	}
}

func multiplyBy(m int) multiplier {
	return func(i int) int {
		return i * m
	}
}
