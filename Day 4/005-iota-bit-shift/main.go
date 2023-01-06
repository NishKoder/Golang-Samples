package main

import "fmt"

// For multiple variable
// iota is auto incremental
const (
	a = iota
	b
	c
)

func main() {
	// basic iota
	func1()

	// basic bit - shifting
	func2()

	//iota with bit shift
	func3()
}

func func1() {
	fmt.Println(a, b, c)
}

func func2() {
	s := 6
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)

	kb := 1024
	mb := 1024 * 1024
	gb := 1024 * 1024 * 1024
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func func3() {
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
