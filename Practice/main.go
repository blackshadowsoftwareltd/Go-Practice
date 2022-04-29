package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Hello World World"
	fmt.Print(strings.ReplaceAll(data, "World", "Earth"))

}
