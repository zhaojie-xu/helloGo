package main

import "io/ioutil"

func main() {
	data := []byte("Hello World\n")
	err := ioutil.WriteFile("/Users/danny/Desktop/data.cfg", data, 777)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
