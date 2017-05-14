package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	s := person{name: "Sean", age: 24}
	fmt.Println(s.age)
	sp := &s
	fmt.Println(sp.name, sp.age)
	sp.age = 25
	fmt.Println(sp.age)
	fmt.Println(s)
}
