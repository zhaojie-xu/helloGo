package main

import "fmt"

func main() {
	s := "Hello World"
	bs := []byte(s)
	bs[1] = 'B'
	fmt.Println(s)
	fmt.Println(string(bs))

	for _, r := range s {
		fmt.Printf("%c ", r)
	}
}
