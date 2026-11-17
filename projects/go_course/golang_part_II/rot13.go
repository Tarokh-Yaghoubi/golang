package main

import "fmt"

func rot13(str string) string {
	ln := len(str)
	rot := make([]byte, ln)

	for i := 0; i < ln; i++ {
		if (str[i] >= 65 && str[i] <= 77) || (str[i] >= 97 && str[i] <= 109) {
			rot[i] = str[i] + 13
		} else if (str[i] >= 78 && str[i] <= 90) || (str[i] >= 110 && str[i] <= 122) {
			rot[i] = str[i] - 13
		} else if str[i] <= 64 || (str[i] >= 91 && str[i] <= 96 || str[i] >= 123) {
			rot[i] = str[i]
		}
	}

	return string(rot)
}

func main() {
	res := rot13("TEST")
	fmt.Println("res is => ", res)
}
