package main

import "fmt"

var x = 34

func main() {
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%#X\n", x)
	fmt.Printf("%X\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\t%b\t%x\t%#x\n", x, x, x, x)

	str := fmt.Sprintf("%T\t%b\t%x\t%#x\n", x, x, x, x)
	fmt.Println(str)

	fmt.Printf("%v", str)

}
