package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"jeremy": 24,
		"jordie": 22,
		"josh":   27,
	}
	fmt.Println(keys(m))
}

func keys(m map[string]int) (names []string) {
	for key := range m {
		names = append(names, key)
	}
	return
}
