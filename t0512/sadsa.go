package main

import (
	"fmt"
	"time"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[2:4]
	s[0] += 100
	s[1] += 100
	fmt.Println(data)

	fmt.Println(len(s), cap(s))

	ss := s[1:4]
	fmt.Println(len(ss), cap(ss))

	t := time.Now()
	fmt.Println(t)
}
