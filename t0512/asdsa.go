package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}

		a[i] = v + 100
	}

	fmt.Println(a)

	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
		}

		fmt.Println(i, v)
	}

}
