package main

import "fmt"

func createMap(names ...string) (m map[string]bool) {
	m = map[string]bool{}

	for _, n := range names {
		m[n] = false
	}

	return m
}

func main() {
	fmt.Println(createMap("John", "Eric", "Linda"))
	fmt.Println(createMap("John", "Eric", "Linda", "Paul"))
}
