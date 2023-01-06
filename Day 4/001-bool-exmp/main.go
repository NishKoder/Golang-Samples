package main

import "fmt"

func main() {
	basicBoolExample()
	anotherBoolExample()
}

var b bool

func basicBoolExample() {
	fmt.Println(b) // O/P - false (By Defualt)
	b = true       // change value
	fmt.Println(b)
}

func anotherBoolExample() {
	a := true
	b := false
	fmt.Println(a == b) //compare bool value
}
