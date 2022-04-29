package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Hello World World"
	fmt.Println(strings.ToUpper(data))
	fmt.Println(strings.ToLower(data))
	fmt.Println(strings.Index(data,"W"))
	fmt.Println(strings.Split(data," "))

}
