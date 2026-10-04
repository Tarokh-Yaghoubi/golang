// interface{}
// this can store any value of type that has zero or more methods
// This matches all types in go

package main

import "fmt"

func printValue(value interface{}) {
	fmt.Println(value)
}

// any == interface{}
func checkTypeExample(value any) {
	// this is called a type-switch
	// (type) is only valid in a type-switch
	switch v := value.(type) {
	case int:
		fmt.Println("INT", v)
	case string:
		fmt.Println("STR", v)
	case bool:
		fmt.Println("BOOL", v)
	default:
		fmt.Println("SMTH ELSE")
	}
}

func typeAssertionExample() {
	var emptyInterface interface{}
	// var emptyInterface any -> this is also possible
	emptyInterface = "Tarokh is a moderate programmer"
	if str, ok := emptyInterface.(string); ok {
		fmt.Println("str is => ", str)
	} else {
		fmt.Println("assertion did not pass")
	}
}

func main() {
	mixedSlice := []interface{}{33, true, false, "Tarokh"}
	for _, val := range mixedSlice {
		printValue(val)
	}
}
