package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

func main() {
	gobook := Book{
		Title:       "Go 语言编程",
		Authors:     []string{"XuShiwei", "HughLv", "XuDaoli"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
		Price:       9.99,
	}

	b, err := json.Marshal(gobook)

	if err == nil {
		fmt.Println(string(b))
		fmt.Println(b[1])
	} else {
		fmt.Println(err.Error())
	}

}
