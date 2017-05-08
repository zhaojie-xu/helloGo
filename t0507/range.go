package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range numbers {
		fmt.Println(i, numbers[i])
	}

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "China": "Beijing"}
	fmt.Println(countryCapitalMap)

	bitmap := map[int]string{1: "One", 2: "Two", 5: "Five"}
	fmt.Println(bitmap[5])

}
