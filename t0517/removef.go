package main

import "os"

func main() {
	fileName := "/Users/danny/Desktop/Hello.txt"
	//os.Create(fileName)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	f.Write([]byte("Hello World\n"))

}
