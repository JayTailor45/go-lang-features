package main

import "fmt"

type CustomString string

func (text CustomString) log() {
	fmt.Println(text)
}

func main() {
	var name CustomString = "Hello, World!"
	name.log()
}
