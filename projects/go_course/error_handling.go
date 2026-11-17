package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
)

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

func main() {
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
