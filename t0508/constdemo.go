package main

import "fmt"

func main() {
	const world = "世界"
	fmt.Println(world)

	for i := 0; i < 10; i++ {
		fmt.Println(i + 10000)
	}

	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)

}
