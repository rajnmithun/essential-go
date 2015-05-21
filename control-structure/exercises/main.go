package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 29; i++ {
		fmt.Print(i, " ")
		if (i+1)%10 == 0 {
			fmt.Println("")
		}
	}
}
