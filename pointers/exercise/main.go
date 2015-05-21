package main

import (
	"fmt"
)

func main() {
	a, b := 10, 5
	fmt.Println("before swap", "a =", a, ",b =", b)
	swap(&a, &b)
	fmt.Println("after swap", "a =", a, ",b =", b)
}

func swap(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}
