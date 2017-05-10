package main

import (
	"math"
	"fmt"
)

type Vtx struct {
	X, Y float64
}

func abs(v Vtx) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vtx{3, 4}
	fmt.Println(abs(v))
}
