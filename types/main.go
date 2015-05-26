package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	p := point{20, 40}
	fmt.Println(p)

	r := rect{NewPoint(20, 40), 30, 40}
	fmt.Println(r)

	r2 := rect{
		pos:    NewPoint(20, 40),
		width:  30,
		height: 40,
	}
	fmt.Println(r2)
	fmt.Println("area of rectangle ", r2.area())
}
