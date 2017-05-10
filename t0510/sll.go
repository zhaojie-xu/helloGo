package main

import "fmt"

type VertexX struct {
	Lat, Long float64
}

var m map[string]VertexX

func main() {
	m = make(map[string]VertexX)
	m["Bell"] = VertexX{213.213, 4.3421}
	fmt.Println(m)
	fmt.Println(m["Bell"])

	var ma = map[string]VertexX{
		"Ne": {12.34, 234.12},
		"Ac": {34.31, 56.1},
	}
	fmt.Println(ma)

	ele, ok := m["Bell"]
	fmt.Println(ele, ok)
}
