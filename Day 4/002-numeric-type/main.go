package main

import "fmt"

func main() {
	// With short declare
	func1()

	//With global decalre
	func2()
}

func func1() {
	a := 33
	b := 33.333
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

var x int

var y float64

func func2() {
	x = 33
	y = 33.333
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
