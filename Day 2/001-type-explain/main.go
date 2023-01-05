package main

import "fmt"

var value = 43

// Declare the VARIABLE with the IDENTIFIRE 'str'
// is of TYPE string
// This is called STATIC PROGRAMMING LANGUAGE
// EACH VALUE have TYPE
// a VARIABLE is DELCARED to hold a VALUE certain TYPE

var str string = "Hel"

// Include the double qoutes at string

var dstr = `Hey Nishant, 
"Your used 
double quotes"`

func main() {
	fmt.Println("Type")
	fmt.Println(value)
	fmt.Printf("%T\n", value)
	fmt.Println(str)
	fmt.Printf("%T\n", str)
	fmt.Println(dstr)
	fmt.Printf("%T\n", dstr)
	// str = 42
	// fmt.Println(str)
	// fmt.Printf("%T\n", str)
}
