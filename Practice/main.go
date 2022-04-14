package main

import (
	"fmt"
)

func main() {
	sum := 1
	for {
		sum = sum + 1
		if sum == 10 {
			break
		}
	}
	fmt.Println(sum)

}
