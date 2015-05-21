package main

import (
	"fmt"
)

func main() {
	if true == true {
		fmt.Println("True is true")
	} else {
		fmt.Println("True is false")
	}

	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("This will go forever")
		break
	}

	j := true
	for j {
		fmt.Println("This should happen once")
		j = false
	}
}
