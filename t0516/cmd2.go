package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("echo", "HELLO")
	v, err := cmd.Output()
	if err == nil {
		fmt.Println(string(v))
	}

}
