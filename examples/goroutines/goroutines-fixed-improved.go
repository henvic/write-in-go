package main

import "fmt"
import "sync"
import "time"

func f(from string) {
	fmt.Println(from, ":", time.Now())
}

func main() {
	var queue sync.WaitGroup
	var max = 5
	queue.Add(max)
	for i := 0; i < max; i++ {
		go func(num int) {
			fmt.Printf("num = %v\n", num)
			queue.Done()
		}(i)
	}
	go f("another")
	queue.Wait()
}
