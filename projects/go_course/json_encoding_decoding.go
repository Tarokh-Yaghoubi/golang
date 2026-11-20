package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Man struct {
	Username string `json:"user_name"`
	Password string `json:"user_password"`
	IsActive bool   `json:"isUserActive"`
}

func decodeJsonExample() {
	var man Man
	f, err := os.Open("output.json")
	if err != nil {
		fmt.Println("err is => ", err)
		panic(err)
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&man)
	if err != nil {
		fmt.Println("error encoding man => ", err)
		panic(err)
	}

	fmt.Println("man data => ", man)
}

func main() {
	m := Man{
		Username: "Tarokh Yaghoubi",
		Password: "TarokhIsNormalHaha0100",
		IsActive: false,
	}

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("error creating the file: ", err)
		panic(err)
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(m)
	if err != nil {
		fmt.Println("Error encoding MAN : ", err)
		panic(err)
	}

	fmt.Println("============== DECODE =================")
	decodeJsonExample()
}
