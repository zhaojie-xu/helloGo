package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	fmt.Println(reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println(p.CanSet())

	v = p.Elem()
	fmt.Println(v.CanSet())
	fmt.Println(v.Type())

}
