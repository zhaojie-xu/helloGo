package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	swaps(a, b)

	fmt.Println(a)
	fmt.Println(b)
}

func swaps(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp

	return temp
}
