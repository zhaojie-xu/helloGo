package main

import "fmt"

func main() {
	var a int = 100
	var b int = 123

	swapK(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
func swapK(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
