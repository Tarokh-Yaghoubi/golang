// Converting golang data structures to JSON data is called marshalling
// Converting JSON data to go data is called unmarshalling
// Package encoding/json provided functions to marshal and unmarshal

// JSON TAGS:
// JSON tags are used in go to map a field of a struct to its JSON counterpart
// Added inside backticks at end of the struct field declaration
// Inside the tags you provide the mapping and other options
// Tag 'omitempty' is used to omit the struct field from marshalling when zero-valued
// Tag '-' is used to indicate the the field must be completely ignored from marshalling

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"` // using omitempty, so whenever this field is empty, go will omit it from marshalling
	Age         int      `json:"age"`
	Password    string   `json:"-"`     // ignore password
	Permissions []string `json:"roles"` // rename "Permissions" to "roles"
}

type Person struct {
	Name       string   `json:"full_name"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"-"`
	Langauage  []string `json:"spoken_languages"`
}

func main() {
	user := User{
		ID:          1,
		Name:        "Tarokh",
		Age:         20,
		Password:    "TheCakeIsAFake2004!$$%",
		Permissions: []string{"owner"},
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error is => ", err)
		panic(err)
	}

	fmt.Println(string(b))
}
