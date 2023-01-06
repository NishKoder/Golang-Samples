package main

import "fmt"

func main() {

	//Use short declarations
	exeOne()

	//Global variable
	exeTwo()

	//Sprint used
	exeThree()
}

func exeOne() {
	// short declaration
	x := 42
	fmt.Print(x)
	fmt.Printf(" \t%T\n", x)

	y := "Nishant Gupta"
	fmt.Print(y)
	fmt.Printf("\t%T\n", y)

	z := true
	fmt.Print(z)
	fmt.Printf("\t%T\n", z)

	// Second
	fmt.Println(x, y, z)
}

var x int

var y string

var z bool

func exeTwo() {
	fmt.Println("")
	fmt.Println("exeTwo")
	fmt.Println("x value", x, "\ny value", y, "\nz value", z)
}

func exeThree() {
	fmt.Println("")
	fmt.Println("exeThree")
	x = 23
	y = "Hello"
	z = true
	s := fmt.Sprintf("x value is %v \t y value is %v \t z value is %v", x, y, z)
	fmt.Println(s)
}
