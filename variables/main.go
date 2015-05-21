package main

import (
	"fmt"
)

var (
	author  = "@rajnmithun"
	Version = "0.0.1"
)

const (
	CCVisa       = "Visa"
	CCMasterCard = "MasterCard"
	CCAmex       = "Amex"
)

func main() {
	var greeting string = "hello world!"
	var a, b, c int = 1, 2, 3
	fmt.Println(greeting, a, b, c)

	var name = "Mithun Raj Nanjegowda"
	var d, e, f = 1, 2.0, "Three"
	fmt.Println(name, d, e, f)

	course := "Essential Go"
	x, y, z := 1, 2, 3
	fmt.Println(course, x, y, z)

}
