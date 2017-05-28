package main

import "fmt"

func main() {
	if 1 > 0 || 1 > 0 && 1 < 0 {
		fmt.Println("yes")
	}
}
