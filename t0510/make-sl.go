package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := make([]int, 3, 5)
	c := b[:2]
	d := c[2:5]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
