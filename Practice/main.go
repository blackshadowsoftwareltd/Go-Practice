package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Local())
	x := 5

	if x == 1 {
		fmt.Println("x is 1")
	} else if x == 2 {
		fmt.Println("x is 2")
	} else if x == 3 {
		fmt.Println("x is 3")
	} else if x == 4 {
		fmt.Println("x is 4")
	} else if x == 5 {
		fmt.Println("x is 5")
	} else {
		fmt.Println("x is not found")
	}
	fmt.Println(time.Now().UTC().Local())
}
