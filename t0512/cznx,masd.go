package main

import (
	"fmt"
	"runtime"
)

func main() {
	s := make([]int, 0, 5)
	fmt.Println(&s)
	s2 := append(s, 1, 2, 3, 4)
	fmt.Println(s2)
	fmt.Println(&s2)

	s2 = append(s, 5, 6, 7, 8, 9, 10)
	fmt.Println(s2)

	type User struct {
		id   int
		name string
	}

	fmt.Println(runtime.GOOS)
}
