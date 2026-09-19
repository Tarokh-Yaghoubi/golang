package main

import "fmt"

func test() {
	var myMap map[string]string = map[string]string{} // empty map, not NIL
	myMap["username"] = "Tarokh"
	myMap["color"] = "Black"
	myMap["sirname"] = "Yaghoubi"

	var otherMap map[string]string = map[string]string{
		"username": "John",
		"color":    "Pink",
		"sirname":  "Smith",
	}

	fmt.Println("len of the map => ", len(myMap))
	fmt.Println("len of the otherMap => ", len(otherMap))

	var username string = myMap["username"]
	var username2 string = otherMap["username"]
	fmt.Printf("username=%v\tfavourite color=%v\n", username, myMap["color"])
	fmt.Printf("username=%v\tfavourite color=%v\n", username2, otherMap["color"])

	car, ok := otherMap["car"]
	if !ok {
		// ok will be false if the key does not exist
		fmt.Println("car does not exist, ret=> ", ok)
	} else {
		fmt.Printf("car is => %v", car)
	}

	// I can have slices as values, but not as keys
	var sliceMap map[string][]string = map[string][]string{
		"usernames": {"FirstTarokh", "SecondTarokh", "LastTarokh"},
		"colors":    {"Black", "Brown", "Silver"},
		"cars":      {"Mercedes", "Bentley", "Volkswagen"},
	}

	fmt.Println("len of the sliceMap => ", len(sliceMap))
	fmt.Printf("usernames=>%v\ncolors=>%v\ncars=>%v\n", sliceMap["usernames"], sliceMap["colors"], sliceMap["cars"])

	benz, ok := sliceMap["cars"]
	if !ok {
		fmt.Println("cars does not exist => ", ok)
	} else {
		// search for BENZ
		for i := 0; i <= len(benz)-1; i++ {
			if benz[i] == "Mercedes" {
				fmt.Println("Mercedes exists => ", benz[i])
				break
			} else {
				fmt.Println("BENZ is not here mate")
			}
		}
	}
}

func main() {
	test()
}
