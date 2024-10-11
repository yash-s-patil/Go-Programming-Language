package main

import "fmt"

// alias for built-in data type
type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Yash"
	name.log()
}
