package main

import (
	"fmt"
	"regexp"
)

func main() {
	var exp1, err = regexp.Compile("[a-z]")
	var exp2 = regexp.MustCompile("[0-9]")

	if err != nil {
		println(err.Error())
	}

	fmt.Println(exp1.MatchString("oi"))   // true
	fmt.Println(exp2.MatchString("oi"))   // false
	fmt.Println(exp1.MatchString("0194")) // false
	fmt.Println(exp2.MatchString("0194")) // true
}
