package main

import "fmt"

type MethodError struct {
	Method  string
	Message string
	Code    int
}

func (a MethodError) Error() string {
	return fmt.Sprintf("Method %s error (%d): %s", a.Method, a.Code, a.Message)
}

func getMethodError() error {
	return MethodError{"UNK", "Not implemented", -1}
}

func main() {
	if err := getMethodError(); err != nil {
		println(err.Error())
	}
}
