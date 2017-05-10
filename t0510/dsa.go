package main

import "fmt"

type Vs struct {
	X, Y float64
}

func (v *Vs) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vs, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vs{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vs{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)
}
