package main

import (
	"context"
	"fmt"
	"os/exec"
	"reflect"
)

func main() {
	var cmd = exec.CommandContext(context.Background(), "uname", "?")
	var err = cmd.Run()

	println(err.Error())
	println(reflect.TypeOf(err))

	var exitError = err.(*exec.ExitError)
	fmt.Printf("Exited = %v\n", exitError.Exited())
	fmt.Printf("PID = %d\n", exitError.Pid())
}
