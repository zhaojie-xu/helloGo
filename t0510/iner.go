package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	a = v

	fmt.Println(a.Abs())
}
