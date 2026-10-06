
// packages are a collection of source code files that are logically grouped
// first line of code is the package delcaration in a go file
// only one package can live inside a directory with exception of test packages

// Idiomatically a package directory must have the same name as the package file

package math

// exported
const PI = 3.14159

// exported
func Add(first, second int) int {
	return first + second
}

// unexported
// unexported functions are written in lowercase
func subtract(first, second int) int {
	return first - second
}
