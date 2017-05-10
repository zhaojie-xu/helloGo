package main

import (
	"math"
	"fmt"
)

type Vtex struct {
	x, y float64
}

func (v Vtex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vtex) Scale(f float64) {
	v.x *= f
	v.y *= f
}

func main() {
	v := Vtex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
