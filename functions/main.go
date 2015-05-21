package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(double(5))
	first, _ := parseName("Rj Nanjegowda")
	fmt.Println(first)

	greet := func() {
		fmt.Println("hello", first)
	}

	greet()
}

func double(n int) int {
	return n + n
}

func parseName(name string) (first, last string) {
	parse := strings.Split(name, " ")
	return parse[0], parse[1]
}
