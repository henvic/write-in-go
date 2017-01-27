package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name string
}

func main() {
	var me = Person{
		Name: "Henrique Vicente",
	}

	fmt.Fprintf(os.Stderr, "%v\n", me)
}
