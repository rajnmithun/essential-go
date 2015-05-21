package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums))
}

func sum(nums []int) (sum int) {
	for _, value := range nums {
		sum += value
	}
	return
}
