package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)
		for _, l := range buffer[:n] {
			if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
				count++
			}
		}

		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, err
		}
	}
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := countAlphabets(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data => %d\n", n)

	r := strings.NewReader("Tarokh is a moderate programmer+++++++++++++++++")
	rv, errs := countAlphabets(r) // Special characters will be ignored, because of the condition we have written inside the func
	if errs != nil {
		panic(err)
	}

	fmt.Printf("Rv => %d\n", rv)
}
