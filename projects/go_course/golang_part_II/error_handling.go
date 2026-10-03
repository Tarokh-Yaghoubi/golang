// Go has a different way of error handling
// There is not try-catch, we can do if-check for error-handling
// It is a convention that if a function is returning an error, it must be the last value that it returns
// The error string must not be capitalized and must not have any punctuations in the end
// The error type in Go is an interface{}
// It only has one method called error() which returns a string
// You can implement the error method to implement your own way of handling errors, first we need to create a new error using the New() method, then using fmt.Errorf
// to throw the error message

package main

import (
	"errors"
	"fmt"
)

// this func will divide first and second params and return the error as well
func divide(first, second float64) (float64, error) {
	if second == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return (first / second), nil
}

func main() {
	ans, ok := divide(10, 0)
	if ok == nil {
		fmt.Println(ans)
	} else {
		fmt.Println(ok)
	}
}
