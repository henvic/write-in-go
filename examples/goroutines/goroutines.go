package main

import (
	"fmt"
	"time"
)

func f(from string) {
	fmt.Println(from, ":", time.Now())
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("i = %v\n", i)
		}()
	}

	go f("another")
	time.Sleep(time.Second)
}
