package main

import (
	"fmt"
)

func main() {
	var data Objects = Objects{"Remon", 22}
	fmt.Println("Name : ", data.name, "\nAge : ", data.age)
}

type Objects struct {
	name string
	age  int
}
