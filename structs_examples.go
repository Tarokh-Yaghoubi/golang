package main

import "fmt"

func main() {

	type Student struct {
		firstname string
		lastname  string
		age       int
		classes   []string
	}

	var firststudent Student
	fmt.Printf("firststudent => %+v\n", firststudent)

	firststudent = Student{"Tarokh", "Yaghoubi", 22, []string{"Programming", "German Language"}}
	fmt.Printf("firststudent => %+v\n", firststudent)

	firststudent.classes = append(firststudent.classes, "Music Class")
	fmt.Printf("firststudent => %+v\n", firststudent)

	// Anonymous struct
	guardian := struct {
		username string
		password string
		isActive bool
	}{
		username: "tarokh",
		password: "golangislovable0010!",
		isActive: true,
	}

	fmt.Printf("guardian anon struct => %+v\n", guardian)
}
