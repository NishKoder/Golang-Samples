package main

import (
	"fmt"
)

func dimensions(length, width, weight int) (area int, volume int) {
	area = length * width
	volume = length * width * weight
	return
}

func main() {
	area, volume := dimensions(4, 4, 3)
	fmt.Println("Area: ", area, " , Volume: ", volume)
}
