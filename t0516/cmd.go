package main

import (
	"os/exec"
	"fmt"
)

func run() {
	cmd := exec.Command("/bin/sh", "-c", "ls")
	a, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

	fmt.Println("MARK", a)
}

func main() {
	go run()
}
