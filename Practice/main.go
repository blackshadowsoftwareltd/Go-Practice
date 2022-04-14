package main

import (
	"fmt"
)

func main() {
	var data Objects = Objects{"Remon", "", 22}
	data.lastName = "Ahammad"
	fmt.Println("First name : ", data.firstName, "\nLast name : ", data.lastName, "\nAge : ", data.age)
}

type Objects struct {
	firstName, lastName string
	age                 int
}
