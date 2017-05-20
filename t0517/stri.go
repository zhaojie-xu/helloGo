package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("[\\\\w]*")

	fmt.Println(r.MatchString(""))
	fmt.Println(r.MatchString("yax"))
	fmt.Println(r.MatchString("yaAxzxasxa"))
}
