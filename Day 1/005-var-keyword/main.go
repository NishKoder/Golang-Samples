package main

import "fmt"

var global_var = 34

func main() {

	//short declaration operator
	//this is not work on outside scope of function like Local variable only
	//declare variable and assign value
	value := 20
	fmt.Println(value)

	//with var keyword
	// can use as global variable
	var value_var = 100
	fmt.Println(value_var)
	fmt.Println(global_var)

	// declare variable with identifier value_type with Type
	var value_type int
	fmt.Println(value_type) // 0 default for int type

}
