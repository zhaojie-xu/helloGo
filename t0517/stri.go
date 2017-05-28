package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("name\\d2")

	fmt.Println(r.MatchString("name92"))
	fmt.Println(r.MatchString(""))
	fmt.Println(r.MatchString("yaAxzxasxa"))
}
