package main

import "fmt"

var x = 50

type koder int

var y koder

func main() {
	fmt.Println("")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// You can not assign x value to y
	// now y have diff TYPE 'main.koder'
	// Go-lang is STATIC programming language
	// here we are convert x into y TYPE not casting
	y = koder(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
