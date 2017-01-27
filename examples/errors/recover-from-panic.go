package main

import (
	"fmt"
	"os"
	"regexp"
)

func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%v\n", r)
		}
	}()

	regexp.MustCompile("[0-9")
	println("Bye, World! (never printed)")
}

func main() {
	println("Hello, World!")
	example()
	println("Bye, World!")
}
