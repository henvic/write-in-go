package main

import (
	"fmt"
	"io/ioutil"
)

func printRegistry() {
	var content, err = ioutil.ReadFile("registry.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func main() {
	printRegistry()
}
