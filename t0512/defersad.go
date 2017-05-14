package main

func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y)
	}(x)

	x += 2000
	y += 200000
	println("x =", x, "y =", y)
}

func main() {
	test()
}
