package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printRegistry() error {
	var content, err = ioutil.ReadFile("registry.json")

	if err != nil {
		return err
	}

	fmt.Println(string(content))
	return nil
}

func main() {
	if err := printRegistry(); err != nil {
		fmt.Fprintf(os.Stderr, "Not able trying to print registry:\n%v\n", err)
	}
}
