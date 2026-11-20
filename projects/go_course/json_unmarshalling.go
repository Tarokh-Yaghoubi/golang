package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string   `json:"fullname"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"-"`
	Languages  []string `json:"spoken_languages"`
}

func main() {

	json_data := `{"fullname":"Tarokh Yaghoubi", "years_old":22, "spoken_languages":["German", "English"]}`
	var person Person
	err := json.Unmarshal([]byte(json_data), &person)

	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		panic(err)
	}

	fmt.Println("name: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Println("Langs: ", person.Languages)
	fmt.Println("Occupation: ", person.Occupation)

}
