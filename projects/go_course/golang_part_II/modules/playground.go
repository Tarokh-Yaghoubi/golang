package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Tarokh is a moderate programmer\n")
	color.Red("Tarokh is a moderate programmer\n")

	str := color.CyanString("Go is a good choice")
	fmt.Printf("str => %s\n", str)
}
