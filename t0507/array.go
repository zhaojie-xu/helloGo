package main

import "fmt"

func main() {
	var balance [10] float32
	var b2 = [5]float32{100.0, 2.0, 45.6, 43.5, 72.7}
	balance[1] = 12.121
	fmt.Println(balance)
	fmt.Println(b2)
}
