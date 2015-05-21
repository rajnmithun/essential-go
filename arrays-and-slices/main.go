package main

import (
	"fmt"
)

func main() {
	var nums [5]int
	fmt.Println("empty:", nums)

	nums[4] = 100
	fmt.Println("set", nums)
	fmt.Println("get", nums[4])

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("empty", ints)

	ints = append(ints, 6)
	fmt.Println(ints)

	fmt.Println(ints[:2])
	fmt.Println(ints[2:4])
	fmt.Println(ints[4:])

	for i, value := range ints {
		fmt.Println(i, value)
	}
}
