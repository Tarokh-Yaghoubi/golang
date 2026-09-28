
package main

import "fmt"
import "errors"

type student struct {
	firstname string 
	lastname string
	age int
}

func main() {
	var a int = 19
	increment(&a)
	fmt.Println(a)

	s := student {
		firstname: "Tarokh",
		lastname: "Smith",
		age: 22,
	}

	fmt.Println("current first name of struct => ", s.firstname)
	previousFirstname, err := updateFirstname(&s, "John")
	fmt.Println("updated first name of the struct => ", s.firstname)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("previous first name returned by the func => ", previousFirstname)
	}
}

func increment(x * int) {
	(*x)++
}

func updateFirstname(s * student, newFirstname string) (string, error) {
	// we do not need to dereference the variable here
	// go will handle it automatically in structs

	if newFirstname == "" {
		return "", errors.New("empty new firstname")
	}
	previous := s.firstname
	s.firstname = newFirstname
	return previous, nil
}