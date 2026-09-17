// maps are data structures used to store key-value pairs
// behind the scene go uses hash maps
// the len function can be used to get the number of key-value pairs in a map

package main

import "fmt"


func main() {
	
	// the zero value for a map is NIL
	// a nil map can be created using a variable declaration
	// Example:
	var nameAge map[string]int	// keys=strings, values=int
	// this is a nil map, it is declared but not defined
	// len will return 0 if the map has nothing in it
	// writing to a nil map causes PANIC


	// EMPTY MAP
	nameMajor := map[string]int{}	// empty map, not NIL
	nameHeight := make(map[string]int)
	var nameLanguage map[string]int = map[string]int

	// this is a normal idiom in go, using ok
	name, ok := nameLanguage["username"]
	if !ok {
		fmt.Println("not found")
	}
	
	// map keys can only be of comparable types
	// values can be of any type
	// map can only be compared to nil
}
