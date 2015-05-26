package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type circle struct {
	pos    point
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * 2 * c.radius
}

func main() {
	c := circle{NewPoint(2, 5), 3}
	fmt.Println("circle area ", c.area())
}
