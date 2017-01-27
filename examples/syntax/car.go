package main

import (
	"fmt"
	"time"
)

type Flight struct {
	ID       string
	Duration time.Duration
}

// String is a method (function of a type) of type Flight
func (f Flight) String() string {
	return fmt.Sprintf("Flight %v is about %v long", f.ID, f.Duration)
}

func main() {
	var lax2jfk = Flight{
		ID:       "D40",
		Duration: time.Duration(5*time.Hour + 29*time.Minute),
	}
	fmt.Println(lax2jfk.String())
}
