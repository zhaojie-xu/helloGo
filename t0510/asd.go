package main

import (
	"fmt"
	"math"
)

type If interface {
	M()
}

type Ty struct {
	S string
}

func (t *Ty) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i If) {
	fmt.Printf("%v %T\n", i, i)
}

func main() {
	var i If
	i = &Ty{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
