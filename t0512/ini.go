package main

import "fmt"

func main() {

	var a = struct {
		x int
		s string
	}{12, "Hi"}
	a.s = "Go"
	fmt.Println(a)

	arr := []int{1, 2, 3}
	fmt.Println(arr)

	s := "abs"
	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(s[i])
	}
}
