package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)
	s = append(s, 2)
	fmt.Println(s)
	s = append(s, 6, 1, 8)
	fmt.Println(s)
}
