package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 234)
	numbers = append(numbers, 7, 9, 45)
	fmt.Println(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	fmt.Println(numbers1)

}
