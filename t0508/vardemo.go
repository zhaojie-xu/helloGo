package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("%T %v\n", ToBe, ToBe)
	fmt.Printf("%T %v", MaxInt, MaxInt)

	var b bool
	var str string

	b = false

	str = "Hi"

	fmt.Printf(b)
	fmt.Printf(str)
}
