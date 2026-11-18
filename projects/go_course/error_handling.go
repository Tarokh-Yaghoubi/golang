package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
)


type Person struct {
	firstName string
	lastName  string
	age 	  int
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("number is not even")
	}

	return i * 2, nil
}

func zipProcessing(fileName string) (int, error) {
	if fileName == "" {
		return 1, fmt.Errorf("filename is not valid")
	}

	r, err := zip.OpenReader(fileName)
	if err != nil {
		if errors.Is(err, zip.ErrFormat) {
			return 1, fmt.Errorf("not a proper format")
		}
		return 1, fmt.Errorf("open zip file error : %w", err)
	}
	defer r.Close()

	return 0, nil
}

func validatePerson(p Person) error {
	var errs []error
	if len(p.firstName) == 0 {
		errs = append(errs, errors.New("field firstName cannot be empty"))
	}

	if len(p.lastName) == 0 {
		errs = append(errs, errors.New("field lastName cannot be empty"))
	}
	if p.age < 0 {
		errs = append(errs, errors.New("field age cannot be negative"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}

	f.Close()
	return nil
}


type MyErr struct {
	Codes []int
}


func (me *MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

/// NOTE: I fairly know how it works :)
func (me MyErr) Is(target error) bool {
	if me2, ok := target.(*MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

func main() {

	fileCheckerErr := fileChecker("not_here.dat")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file does not exist")
		}
	}

	first_person := Person{firstName: "Tarokh", lastName: "", age: 22}
	person_ret := validatePerson(first_person)
	fmt.Println("validate person return val => ", person_ret)

	remainder, mod, ret := calcRemainderAndMod(10, 2)
	if ret == nil {
		fmt.Println(remainder, " - ", mod)
	} else {
		fmt.Println(ret)
		os.Exit(1)
	}

	ans, err := doubleEven(4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("doubled even => ", ans)

	ans, err = zipProcessing("file.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
