package main

import "fmt"

func arrays() {
	
	var intArr [3]int32;

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	fmt.Println("addresses =========> ")
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// we can directly initialize an array using this syntax
	var intArray [3]int32 = [3]int32{22, 23, 25}
	// or using the collon-equal shorthand: 
	arr := [3]int32{25, 26, 27}
	fmt.Printf("%v %v %v\n", intArray[0], intArray[1], intArray[2])
	fmt.Printf("%v %v %v\n", arr[0], arr[1], arr[2])

	// we can even ommit that first 3 
	arr2 := [...]int32{25, 26, 27}
	fmt.Println(arr2)
	// slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("intSlice before append len => %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("intSlice after append len => %v with capacity %v\n", len(intSlice), cap(intSlice))

	// another way to create a slice is to use the make() function. 
	// make() allows you to say how many members does your slice have,
	// and also pass the specified cap of the slice.
	var intSlice3 []int32 = make([]int32, 3, 8); 
	intSlice3 = append(intSlice3, 88)
	intSlice3 = append(intSlice3, 87)
	intSlice3 = append(intSlice3, 86)

	fmt.Printf("intSlice3 val =========> %v\n", intSlice3)	// it appends everything after the zeros, [0, 0, 0, 88, 87, 86]

	// MAP : key-value pairs
	var myMap map[string]string = make(map[string]string)
	myMap["firstname"] = "Tarokh"
	myMap["sirname"] = "Jacobi"
	myMap["gender"] = "male"
	var myMap2 = map[string]uint8{"Tarokh": 190, "Jerrard": 200}
	fmt.Printf("first map => %v\nsecond map => %v\n", myMap, myMap2)

	// maps return a second value which shows whether the key was successfully found inside the map or not
	var firstname, ok = myMap["firstname"]
	if !ok {
		fmt.Println("Firstname does not live in the map")
	} else {
		fmt.Printf("Firstname lives in the map => [%v]", firstname)
	}

}
