package main

import "fmt"

func main() {
	fmt.Println("Hello Dev")
	fun1()
	fmt.Println("Added more")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fun2()
}

func fun1() {
	fmt.Println("Hello Fun 1")
}

func fun2() {
	fmt.Println("I m Fun 2")
}

//Control flow
//1# sequence
//2# loop
//3# conditional
