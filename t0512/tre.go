package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	a[1] = 10
	b := make([]int, 3)
	b[1] = 12
	c := new([]int)
	c = &b
	fmt.Println(c)
}
