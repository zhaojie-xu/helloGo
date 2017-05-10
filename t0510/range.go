package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("%d %d\n", i, v)
	}
}
