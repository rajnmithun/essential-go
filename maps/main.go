package main

import (
	"fmt"
)

func main() {
	age := make(map[string]int)
	fmt.Println("empty:", age)

	age["rj"] = 31
	age["patre"] = 31
	fmt.Println(age)

	fmt.Printf("Rj's age is %v \n", age["rj"])

	delete(age, "rj")
	fmt.Println(age)

	fmt.Println("length of age map is", len(age))

	m := map[string]int{
		"rj":    31,
		"patre": 31,
	}
	fmt.Println(m)

	for key, value := range m {
		fmt.Printf("%v is %v years old.\n", key, value)
	}
}
