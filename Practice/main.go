package main

import (
	"fmt"
)

func main() {
	name := "Remon"
	address := &name //? for getting memory address. use &

	fmt.Println("memory address of name: ", address)
	fmt.Println("value of memory address: ", *address) //? for getting value of memory address. use *

	fmt.Println(name)
	changValueUsingPointer(address)
	fmt.Println(name)
}

func changValueUsingPointer(x *string) { //? for changing value usint pointer use * before data type and data
	*x = "Remon Ahammad"
}
