package main

import (
	"math"
	"fmt"
)

type Vxx struct {
	X, Y float64
}

func Abs(v Vxx) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vxx, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vxx{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
