package main

import (
	"fmt"
)

type Object struct {
	name string
	age  int
}

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	/// range return index and value from slice
	for index, value := range pow {
		fmt.Printf(" %d index = value %d\n", index, value)
	}

	/// range without index
	for _, value := range pow {
		fmt.Printf("value %d\n", value)
	}

	/// range without value
	for index, _ := range pow {
		fmt.Printf("index %d\n", index)
	}
}
