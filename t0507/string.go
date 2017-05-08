package main

import "fmt"

func main() {
	var greeting = "Hello World"
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%X ", greeting[i])
	}
	fmt.Println()
	fmt.Printf("%c ", greeting[8])
}
