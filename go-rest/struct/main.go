package main

import (
	"fmt"
)

type Dimensions struct {
	length int
	width int
	weight int
}

func (d *Dimensions) Area () int {
	d.weight = 43
	return d.length * d.width
}

func (d Dimensions) Volume () int {
	return d.length * d.width * d.weight
}

func main(){
	d := Dimensions {3,6,7}
	fmt.Println(d.Area())
	fmt.Println(d)
	fmt.Println(d.Volume())
}