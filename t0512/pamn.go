package main

func testx() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("panic error!")
}

func main() {
	testx()
}
