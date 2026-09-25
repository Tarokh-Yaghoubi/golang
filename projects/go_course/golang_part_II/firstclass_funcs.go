package main

import "fmt"

// In go you can assign a function to a variable and then you can use it later
// Functions can be passed to other functions

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
