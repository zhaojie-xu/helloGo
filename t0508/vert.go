package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 23})
	v := Vertex{23, 24}
	v.x = 4
	fmt.Println(v)

	p := &v
	p.x = 1e5
	fmt.Println(v)
}
