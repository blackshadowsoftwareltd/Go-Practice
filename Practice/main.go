package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Local())
	x := 5

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("default")
	}
	fmt.Println(time.Now().UTC().Local())
}