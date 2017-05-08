package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var a, b, c = 3, 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\t", a)
	fmt.Printf("%T\t", b)
	fmt.Printf("%T\t", c)

	const LENGTH int = 10
	const WIDTH int = 5
	fmt.Printf("value of ares:%d\n", LENGTH*WIDTH)

	fmt.Println(&a)

	var ptr *int
	ptr = &b
	fmt.Println(*ptr)

	as, bs := swap("Mash", "Kuna")
	fmt.Println(as, bs)

}

func swap(x, y string) (string, string) {
	return y, x

}
