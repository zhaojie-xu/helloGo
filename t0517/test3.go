package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "abcdefgfcvcvcvc"
	s = strings.Replace(s, "c", "C", -1)
	fmt.Println(s)

	fmt.Printf("%s\n", s)
	bs := fmt.Sprintf("%s\n", s)
	fmt.Printf("%T\n", bs)
}
