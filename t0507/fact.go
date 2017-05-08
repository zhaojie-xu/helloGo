package main

import "fmt"

func main() {
	var i int = 3
	fmt.Println(factorial(i))
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
