package main

import (
	"fmt"
	"encoding/json"
)

type Gender string

const (
	MALE   = Gender("male")
	FEMALE = Gender("female")
)

type Person struct {
	name    string
	id      int
	gender  Gender
	address string
}

func main() {
	s := "abcdefg"
	fmt.Println(&s)
	fmt.Println(*(&s))

	p := &Person{
		name:    "Joe",
		id:      123,
		gender:  MALE,
		address: "Beijing",
	}

	fmt.Println(&p)
	fmt.Println(*p)
	fmt.Println(p)

	if d, err := json.Marshal(p); err == nil {
		fmt.Printf("%v\n", d)
	}

}
